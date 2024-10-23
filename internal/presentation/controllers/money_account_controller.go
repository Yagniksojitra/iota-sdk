package controllers

import (
	"fmt"
	"github.com/go-faster/errors"
	"github.com/iota-agency/iota-erp/pkg/middleware"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gorilla/mux"
	"github.com/iota-agency/iota-erp/internal/app/services"
	moneyAccount "github.com/iota-agency/iota-erp/internal/domain/aggregates/money_account"
	"github.com/iota-agency/iota-erp/internal/presentation/mappers"
	"github.com/iota-agency/iota-erp/internal/presentation/templates/pages/moneyaccounts"
	"github.com/iota-agency/iota-erp/internal/presentation/viewmodels"
	"github.com/iota-agency/iota-erp/pkg/composables"
)

type MoneyAccountController struct {
	app      *services.Application
	basePath string
}

func NewMoneyAccountController(app *services.Application) Controller {
	return &MoneyAccountController{
		app:      app,
		basePath: "/finance/moneyaccounts",
	}
}

func (c *MoneyAccountController) Register(r *mux.Router) {
	router := r.PathPrefix(c.basePath).Subrouter()
	router.Use(middleware.RequireAuthorization())
	router.HandleFunc("", c.List).Methods(http.MethodGet)
	router.HandleFunc("", c.Create).Methods(http.MethodPost)
	router.HandleFunc("/{id:[0-9]+}", c.GetEdit).Methods(http.MethodGet)
	router.HandleFunc("/{id:[0-9]+}", c.PostEdit).Methods(http.MethodPost)
	router.HandleFunc("/{id:[0-9]+}", c.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/new", c.GetNew).Methods(http.MethodGet)
}

func (c *MoneyAccountController) viewModelCurrencies(r *http.Request) ([]*viewmodels.Currency, error) {
	currencies, err := c.app.CurrencyService.GetAll(r.Context())
	if err != nil {
		return nil, err
	}
	viewCurrencies := make([]*viewmodels.Currency, 0, len(currencies))
	for _, currency := range currencies {
		viewCurrencies = append(viewCurrencies, mappers.CurrencyToViewModel(currency))
	}
	return viewCurrencies, nil
}

func (c *MoneyAccountController) List(w http.ResponseWriter, r *http.Request) {
	pageCtx, err := composables.UsePageCtx(
		r,
		composables.NewPageData("Accounts.Meta.List.Title", ""),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	params := composables.UsePaginated(r)
	accountEntities, err := c.app.MoneyAccountService.GetPaginated(r.Context(), params.Limit, params.Offset, []string{})
	if err != nil {
		http.Error(w, errors.Wrap(err, "Error retrieving moneyaccounts").Error(), http.StatusInternalServerError)
		return
	}
	viewAccounts := make([]*viewmodels.MoneyAccount, len(accountEntities))
	for i, entity := range accountEntities {
		viewAccounts[i] = mappers.MoneyAccountToViewModel(entity, fmt.Sprintf("%s/%d", c.basePath, entity.ID))
	}
	isHxRequest := len(r.Header.Get("Hx-Request")) > 0
	props := &moneyaccounts.IndexPageProps{
		PageContext: pageCtx,
		Accounts:    viewAccounts,
	}
	if isHxRequest {
		templ.Handler(moneyaccounts.AccountsTable(props), templ.WithStreaming()).ServeHTTP(w, r)
	} else {
		templ.Handler(moneyaccounts.Index(props), templ.WithStreaming()).ServeHTTP(w, r)
	}
}

func (c *MoneyAccountController) GetEdit(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		http.Error(w, "Error parsing id", http.StatusInternalServerError)
		return
	}

	pageCtx, err := composables.UsePageCtx(
		r,
		composables.NewPageData("Accounts.Meta.Edit.Title", ""),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	entity, err := c.app.MoneyAccountService.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Error retrieving account", http.StatusInternalServerError)
		return
	}
	currencies, err := c.viewModelCurrencies(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	props := &moneyaccounts.EditPageProps{
		PageContext: pageCtx,
		Account:     mappers.MoneyAccountToViewUpdateModel(entity),
		Currencies:  currencies,
		Errors:      map[string]string{},
		PostPath:    fmt.Sprintf("%s/%d", c.basePath, id),
		DeletePath:  fmt.Sprintf("%s/%d", c.basePath, id),
	}
	templ.Handler(moneyaccounts.Edit(props), templ.WithStreaming()).ServeHTTP(w, r)
}

func (c *MoneyAccountController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		http.Error(w, "Error parsing id", http.StatusInternalServerError)
		return
	}

	if _, err := c.app.MoneyAccountService.Delete(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	redirect(w, r, c.basePath)
}

func (c *MoneyAccountController) PostEdit(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	action := FormAction(r.FormValue("_action"))
	if !action.IsValid() {
		http.Error(w, "Invalid action", http.StatusBadRequest)
		return
	}
	r.Form.Del("_action")

	switch action {
	case FormActionDelete:
		if _, err := c.app.MoneyAccountService.Delete(r.Context(), id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case FormActionSave:
		dto := moneyAccount.UpdateDTO{} //nolint:exhaustruct
		var pageCtx *composables.PageContext
		pageCtx, err = composables.UsePageCtx(r, composables.NewPageData("Accounts.Meta.Edit.Title", ""))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := decoder.Decode(&dto, r.Form); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		errorsMap, ok := dto.Ok(pageCtx.UniTranslator)
		if ok {
			if err := c.app.MoneyAccountService.Update(r.Context(), id, &dto); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			entity, err := c.app.MoneyAccountService.GetByID(r.Context(), id)
			if err != nil {
				http.Error(w, "Error retrieving account", http.StatusInternalServerError)
				return
			}
			currencies, err := c.viewModelCurrencies(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			props := &moneyaccounts.EditPageProps{
				PageContext: pageCtx,
				Account:     mappers.MoneyAccountToViewUpdateModel(entity),
				Currencies:  currencies,
				Errors:      errorsMap,
				PostPath:    fmt.Sprintf("%s/%d", c.basePath, id),
				DeletePath:  fmt.Sprintf("%s/%d", c.basePath, id),
			}
			templ.Handler(moneyaccounts.EditForm(props), templ.WithStreaming()).ServeHTTP(w, r)
			return
		}
	}
	redirect(w, r, c.basePath)
}

func (c *MoneyAccountController) GetNew(w http.ResponseWriter, r *http.Request) {
	pageCtx, err := composables.UsePageCtx(r, composables.NewPageData("Accounts.Meta.New.Title", ""))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	currencies, err := c.viewModelCurrencies(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	props := &moneyaccounts.CreatePageProps{
		PageContext: pageCtx,
		Currencies:  currencies,
		Errors:      map[string]string{},
		Account:     mappers.MoneyAccountToViewModel(&moneyAccount.Account{}, ""), //nolint:exhaustruct
		PostPath:    c.basePath,
	}
	templ.Handler(moneyaccounts.New(props), templ.WithStreaming()).ServeHTTP(w, r)
}

func (c *MoneyAccountController) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dto := moneyAccount.CreateDTO{} //nolint:exhaustruct
	if err := decoder.Decode(&dto, r.Form); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pageCtx, err := composables.UsePageCtx(r, composables.NewPageData("Accounts.Meta.New.Title", ""))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if errorsMap, ok := dto.Ok(pageCtx.UniTranslator); !ok {
		currencies, err := c.viewModelCurrencies(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		entity, err := dto.ToEntity()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		props := &moneyaccounts.CreatePageProps{
			PageContext: pageCtx,
			Currencies:  currencies,
			Errors:      errorsMap,
			Account:     mappers.MoneyAccountToViewModel(entity, ""),
			PostPath:    c.basePath,
		}
		templ.Handler(moneyaccounts.CreateForm(props), templ.WithStreaming()).ServeHTTP(w, r)
		return
	}

	if err := c.app.MoneyAccountService.Create(r.Context(), &dto); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	redirect(w, r, c.basePath)
}
