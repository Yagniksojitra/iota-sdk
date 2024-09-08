package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/benbjohnson/hashfs"
	"github.com/gorilla/mux"
	"github.com/iota-agency/iota-erp/internal/app"
	"github.com/iota-agency/iota-erp/internal/configuration"
	"github.com/iota-agency/iota-erp/internal/interfaces/graph"
	"github.com/iota-agency/iota-erp/internal/presentation/assets"
	"github.com/iota-agency/iota-erp/internal/presentation/controllers"
	localMiddleware "github.com/iota-agency/iota-erp/pkg/middleware"
	"github.com/iota-agency/iota-erp/sdk/middleware"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/rs/cors"
	"golang.org/x/text/language"
)

var (
	AllowMethods = []string{
		http.MethodConnect,
		http.MethodOptions,
		http.MethodGet,
		http.MethodPost,
		http.MethodHead,
		http.MethodPatch,
		http.MethodPut,
		http.MethodDelete,
	}
)

type Server struct {
	conf *configuration.Configuration
}

func (s *Server) init() error {
	if err := s.conf.Load(); err != nil {
		return err
	}
	return nil
}

func (s *Server) useBundle() *i18n.Bundle {
	bundle := i18n.NewBundle(language.Russian)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.MustLoadMessageFile("pkg/locales/en.json")
	bundle.MustLoadMessageFile("pkg/locales/ru.json")
	return bundle
}

func (s *Server) useRouter(middlewares ...mux.MiddlewareFunc) *mux.Router {
	r := mux.NewRouter()
	r.Use(middlewares...)
	return r
}

func (s *Server) Start() error {
	if err := s.init(); err != nil {
		return err
	}

	log.Println("Connecting to database:", s.conf.DbOpts)
	db, err := ConnectDB(s.conf.DbOpts)
	if err != nil {
		return err
	}
	bundle := s.useBundle()
	application := app.New(db)
	allowOrigins := []string{"http://localhost:3000", "ws://localhost:3000"}

	controllerInstances := []controllers.Controller{
		controllers.NewHomeController(application),
		controllers.NewLoginController(application),
		controllers.NewUsersController(application),
	}

	r := s.useRouter(
		cors.New(cors.Options{
			AllowedOrigins:   allowOrigins,
			AllowedMethods:   AllowMethods,
			AllowCredentials: true,
		}).Handler,
		middleware.RequestParams(middleware.DefaultParamsConstructor),
		middleware.WithLogger(log.Default()),
		middleware.LogRequests(),
		middleware.Transactions(db),
		localMiddleware.WithLocalizer(bundle),
		localMiddleware.Authorization(application.AuthService),
	)

	for _, controller := range controllerInstances {
		controller.Register(r)
	}

	r.Handle("/query", graph.NewDefaultServer(application))
	r.Handle("/playground", playground.Handler("GraphQL playground", "/query"))
	r.HandleFunc("/oauth/google/callback", application.AuthService.OauthGoogleCallback)
	r.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("internal/presentation/static"))))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", hashfs.FileServer(assets.FS)))
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("internal/presentation/public")))

	log.Printf("connect to http://localhost:%s/playground for GraphQL playground", s.conf.ServerPort)
	return http.ListenAndServe(s.conf.SocketAddress, r)
}

func New() *Server {
	return &Server{
		conf: configuration.Use(),
	}
}
