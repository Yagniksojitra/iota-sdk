package core

import (
	"embed"
	"github.com/iota-uz/iota-sdk/pkg/presentation/templates/icons"
	"github.com/iota-uz/iota-sdk/pkg/spotlight"

	"github.com/iota-uz/iota-sdk/modules/core/infrastructure/persistence"
	"github.com/iota-uz/iota-sdk/modules/core/interfaces/graph"
	"github.com/iota-uz/iota-sdk/modules/core/presentation/controllers"
	"github.com/iota-uz/iota-sdk/modules/core/seed"
	"github.com/iota-uz/iota-sdk/modules/core/services"
	"github.com/iota-uz/iota-sdk/pkg/application"
	"github.com/iota-uz/iota-sdk/pkg/presentation/assets"
)

//go:generate go run github.com/99designs/gqlgen generate

//go:embed locales/*.json
var localeFiles embed.FS

//go:embed migrations/*.sql
var migrationFiles embed.FS

func NewModule() application.Module {
	return &Module{}
}

type Module struct {
}

func (m *Module) Register(app application.Application) error {
	app.RegisterMigrationDirs(&migrationFiles)
	app.RegisterLocaleFiles(&localeFiles)
	app.RegisterSeedFuncs(
		seed.CreatePermissions,
		seed.CreateCurrencies,
		seed.CreateUser,
	)
	fsStorage, err := persistence.NewFSStorage()
	if err != nil {
		return err
	}
	app.RegisterServices(
		services.NewUserService(persistence.NewUserRepository(), app.EventPublisher()),
		services.NewSessionService(persistence.NewSessionRepository(), app.EventPublisher()),
	)
	app.RegisterServices(
		services.NewAuthService(app),
		services.NewCurrencyService(persistence.NewCurrencyRepository(), app.EventPublisher()),
		services.NewRoleService(persistence.NewRoleRepository(), app.EventPublisher()),
		services.NewPositionService(persistence.NewPositionRepository(), app.EventPublisher()),
		services.NewEmployeeService(persistence.NewEmployeeRepository(), app.EventPublisher()),
		services.NewUploadService(persistence.NewUploadRepository(), fsStorage, app.EventPublisher()),
		services.NewTabService(persistence.NewTabRepository()),
	)
	app.RegisterControllers(
		controllers.NewDashboardController(app),
		controllers.NewLoginController(app),
		controllers.NewSpotlightController(app),
		controllers.NewAccountController(app),
		controllers.NewEmployeeController(app),
		controllers.NewLogoutController(app),
		controllers.NewUploadController(app),
		controllers.NewUsersController(app),
	)
	app.RegisterHashFsAssets(assets.HashFS)
	app.RegisterGraphSchema(application.GraphSchema{
		Value: graph.NewExecutableSchema(graph.Config{
			Resolvers: graph.NewResolver(app),
		}),
		BasePath: "/",
	})
	sl := app.Spotlight()
	for _, l := range NavItems {
		sl.Register(spotlight.NewItem(l.Icon, l.Name, l.Href))
	}
	app.Spotlight().Register(
		spotlight.NewItem(
			icons.PlusCircle(icons.Props{Size: "24"}),
			"Users.List.New",
			"/users/new",
		),
	)
	return nil
}

func (m *Module) Name() string {
	return "core"
}
