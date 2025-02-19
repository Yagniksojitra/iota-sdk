package controllers

import (
	"github.com/iota-uz/iota-sdk/modules/core/presentation/templates/pages/bichat"
	"github.com/iota-uz/iota-sdk/pkg/application"
	"github.com/iota-uz/iota-sdk/pkg/middleware"
	"github.com/iota-uz/iota-sdk/pkg/shared"
	"github.com/iota-uz/iota-sdk/pkg/types"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gorilla/mux"
	"github.com/iota-uz/iota-sdk/pkg/composables"
)

type BiChatController struct {
	app      application.Application
	basePath string
}

func NewBiChatController(app application.Application) application.Controller {
	return &BiChatController{
		app:      app,
		basePath: "/bi-chat",
	}
}

func (c *BiChatController) Key() string {
	return c.basePath
}

func (c *BiChatController) Register(r *mux.Router) {
	commonMiddleware := []mux.MiddlewareFunc{
		middleware.Authorize(),
		middleware.RedirectNotAuthenticated(),
		middleware.ProvideUser(),
		middleware.Tabs(),
		middleware.WithLocalizer(c.app.Bundle()),
		middleware.NavItems(),
	}
	getRouter := r.PathPrefix(c.basePath).Subrouter()
	getRouter.Use(commonMiddleware...)
	getRouter.HandleFunc("", c.Index).Methods(http.MethodGet)

	setRouter := r.PathPrefix(c.basePath).Subrouter()
	setRouter.Use(commonMiddleware...)
	setRouter.Use(middleware.WithTransaction())
	setRouter.HandleFunc("/{id:[0-9]+}", c.Delete).Methods(http.MethodDelete)
}

func (c *BiChatController) Index(w http.ResponseWriter, r *http.Request) {
	pageCtx, err := composables.UsePageCtx(
		r,
		types.NewPageData("BiChat.Meta.Index.Title", ""),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	props := &bichat.ChatPageProps{
		PageContext: pageCtx,
		Suggestions: []string{"Hello", "World", "IOTA", "UZ"},
	}
	templ.Handler(bichat.Index(props), templ.WithStreaming()).ServeHTTP(w, r)
}

func (c *BiChatController) Delete(w http.ResponseWriter, r *http.Request) {
	//id, err := shared.ParseID(r)
	//if err != nil {
	//	http.Error(w, "Error parsing id", http.StatusInternalServerError)
	//	return
	//}

	shared.Redirect(w, r, c.basePath)
}
