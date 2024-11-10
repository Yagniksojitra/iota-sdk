package shared

import (
	"context"
	"github.com/a-h/templ"
	"github.com/benbjohnson/hashfs"
	"github.com/gorilla/mux"
	"github.com/iota-agency/iota-erp/internal/application"
	"github.com/iota-agency/iota-erp/internal/domain/entities/permission"
)

type ControllerConstructor func(app *application.Application) Controller

type Controller interface {
	Register(r *mux.Router)
}

type NavigationItem struct {
	Name        string
	Href        string
	Children    []NavigationItem
	Icon        templ.Component
	Permissions []permission.Permission
}

type Module interface {
	Name() string
	Seed(ctx context.Context) error
	NavigationItems() []NavigationItem
	Controllers() []ControllerConstructor
	Assets() *hashfs.FS
	LocaleFiles() []string
}
