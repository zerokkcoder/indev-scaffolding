//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/zerokkcoder/indevsca/internal/handler/admin"
	"github.com/zerokkcoder/indevsca/internal/handler/mobile"
	"github.com/zerokkcoder/indevsca/internal/infra"
	"github.com/zerokkcoder/indevsca/internal/repository"
	"github.com/zerokkcoder/indevsca/internal/server"
	"github.com/zerokkcoder/indevsca/internal/service"
	"github.com/zerokkcoder/indevsca/pkg/app"
	"github.com/zerokkcoder/indevsca/pkg/config"
	"github.com/zerokkcoder/indevsca/pkg/jwt"
	"github.com/zerokkcoder/indevsca/pkg/log"
	"github.com/zerokkcoder/indevsca/pkg/server/http"
)

var infraSet = wire.NewSet(
	infra.NewDB,
)

var repositorySet = wire.NewSet(
	repository.NewRepository,
	repository.NewUserRepository,
	repository.NewAdminRepository,
)

var serviceSet = wire.NewSet(
	service.NewServer,
	service.NewUserService,
	service.NewAdminService,
)

var adminHandlerSet = wire.NewSet(
	admin.NewHandler,
	admin.NewAuthHandler,
	admin.NewAdminHandler,
	admin.NewUserHandler,
)

var mobileHanderSet = wire.NewSet(
	mobile.NewHandler,
	mobile.NewUserHandler,
)

var serverSet = wire.NewSet(
	server.NewHTTPServer,
)

// build App
func newApp(
	httpServer *http.Server,
	// job *server.Job,
	// task *server.Task,
) *app.App {
	return app.NewApp(
		app.WithServer(httpServer),
		app.WithName("demo-server"),
	)
}

func NewWire(*config.Config, *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		infraSet,
		repositorySet,
		serviceSet,
		adminHandlerSet,
		mobileHanderSet,
		jwt.NewJwt,
		serverSet,
		newApp,
	))
}
