package controllers

import (
	"fmt"
	"github.com/go-faster/errors"
	"github.com/iota-agency/iota-erp/internal/domain/entities/employee"
	"github.com/iota-agency/iota-erp/pkg/middleware"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gorilla/mux"
	"github.com/iota-agency/iota-erp/internal/app/services"
	"github.com/iota-agency/iota-erp/internal/presentation/mappers"
	"github.com/iota-agency/iota-erp/internal/presentation/templates/pages/employees"
	"github.com/iota-agency/iota-erp/internal/presentation/viewmodels"
	"github.com/iota-agency/iota-erp/pkg/composables"
)

type EmployeeController struct {
	app      *services.Application
	basePath string
}

func NewEmployeeController(app *services.Application) Controller {
	return &EmployeeController{
		app:      app,
		basePath: "/operations/employees",
	}
}

func (c *EmployeeController) Register(r *mux.Router) {
	router := r.PathPrefix(c.basePath).Subrouter()
	router.Use(middleware.RequireAuthorization())
	router.HandleFunc("", c.List).Methods(http.MethodGet)
	router.HandleFunc("", c.Create).Methods(http.MethodPost)
	router.HandleFunc("/{id:[0-9]+}", c.GetEdit).Methods(http.MethodGet)
	router.HandleFunc("/{id:[0-9]+}", c.PostEdit).Methods(http.MethodPost)
	router.HandleFunc("/{id:[0-9]+}", c.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/new", c.GetNew).Methods(http.MethodGet)
}

func (c *EmployeeController) List(w http.ResponseWriter, r *http.Request) {
	pageCtx, err := composables.UsePageCtx(
		r,
		composables.NewPageData("Employees.Meta.List.Title", ""),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	params := composables.UsePaginated(r)
	employeeEntities, err := c.app.EmployeeService.GetPaginated(r.Context(), params.Limit, params.Offset, []string{})
	if err != nil {
		http.Error(w, errors.Wrap(err, "Error retrieving employees").Error(), http.StatusInternalServerError)
		return
	}
	viewEmployees := make([]*viewmodels.Employee, len(employeeEntities))
	for i, entity := range employeeEntities {
		viewEmployees[i] = mappers.EmployeeToViewModel(entity)
	}
	isHxRequest := len(r.Header.Get("Hx-Request")) > 0
	props := &employees.IndexPageProps{
		PageContext: pageCtx,
		Employees:   viewEmployees,
		NewURL:      fmt.Sprintf("%s/new", c.basePath),
	}
	if isHxRequest {
		templ.Handler(employees.EmployeesTable(props), templ.WithStreaming()).ServeHTTP(w, r)
	} else {
		templ.Handler(employees.Index(props), templ.WithStreaming()).ServeHTTP(w, r)
	}
}

func (c *EmployeeController) GetEdit(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		http.Error(w, "Error parsing id", http.StatusInternalServerError)
		return
	}

	pageCtx, err := composables.UsePageCtx(
		r,
		composables.NewPageData("Employees.Meta.Edit.Title", ""),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	entity, err := c.app.EmployeeService.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Error retrieving account", http.StatusInternalServerError)
		return
	}
	props := &employees.EditPageProps{
		PageContext: pageCtx,
		Employee:    mappers.EmployeeToViewModel(entity),
		Errors:      map[string]string{},
		SaveURL:     fmt.Sprintf("%s/%d", c.basePath, id),
		DeleteURL:   fmt.Sprintf("%s/%d", c.basePath, id),
	}
	templ.Handler(employees.Edit(props), templ.WithStreaming()).ServeHTTP(w, r)
}

func (c *EmployeeController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		http.Error(w, "Error parsing id", http.StatusInternalServerError)
		return
	}

	if _, err := c.app.EmployeeService.Delete(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	redirect(w, r, c.basePath)
}

func (c *EmployeeController) PostEdit(w http.ResponseWriter, r *http.Request) {
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
		if _, err := c.app.EmployeeService.Delete(r.Context(), id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case FormActionSave:
		dto := employee.UpdateDTO{} //nolint:exhaustruct
		var pageCtx *composables.PageContext
		pageCtx, err = composables.UsePageCtx(r, composables.NewPageData("Employees.Meta.Edit.Title", ""))
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
			if err := c.app.EmployeeService.Update(r.Context(), id, &dto); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			entity, err := c.app.EmployeeService.GetByID(r.Context(), id)
			if err != nil {
				http.Error(w, "Error retrieving account", http.StatusInternalServerError)
				return
			}
			props := &employees.EditPageProps{
				PageContext: pageCtx,
				Employee:    mappers.EmployeeToViewModel(entity),
				Errors:      errorsMap,
				SaveURL:     fmt.Sprintf("%s/%d", c.basePath, id),
				DeleteURL:   fmt.Sprintf("%s/%d", c.basePath, id),
			}
			templ.Handler(employees.EditForm(props), templ.WithStreaming()).ServeHTTP(w, r)
			return
		}
	}
	redirect(w, r, c.basePath)
}

func (c *EmployeeController) GetNew(w http.ResponseWriter, r *http.Request) {
	pageCtx, err := composables.UsePageCtx(r, composables.NewPageData("Employees.Meta.New.Title", ""))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	props := &employees.CreatePageProps{
		PageContext: pageCtx,
		Errors:      map[string]string{},
		Employee:    mappers.EmployeeToViewModel(&employee.Employee{}), //nolint:exhaustruct
		PostPath:    c.basePath,
	}
	templ.Handler(employees.New(props), templ.WithStreaming()).ServeHTTP(w, r)
}

func (c *EmployeeController) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dto := employee.CreateDTO{} //nolint:exhaustruct
	if err := decoder.Decode(&dto, r.Form); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pageCtx, err := composables.UsePageCtx(r, composables.NewPageData("Employees.Meta.New.Title", ""))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if errorsMap, ok := dto.Ok(pageCtx.UniTranslator); !ok {
		entity := dto.ToEntity()
		props := &employees.CreatePageProps{
			PageContext: pageCtx,
			Errors:      errorsMap,
			Employee:    mappers.EmployeeToViewModel(entity),
			PostPath:    c.basePath,
		}
		templ.Handler(employees.CreateForm(props), templ.WithStreaming()).ServeHTTP(w, r)
		return
	}

	if err := c.app.EmployeeService.Create(r.Context(), &dto); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	redirect(w, r, c.basePath)
}
