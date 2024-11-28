// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

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

// Injectors from wire.go:

func NewWire(configConfig *config.Config, logger *log.Logger) (*app.App, func(), error) {
	jwtJWT := jwt.NewJwt(configConfig)
	handler := admin.NewHandler(logger)
	adminHandler := admin.NewAdminHandler(handler)
	serviceService := service.NewServer(logger, jwtJWT)
	db := infra.NewDB(configConfig, logger)
	repositoryRepository := repository.NewRepository(logger, db)
	adminRepository := repository.NewAdminRepository(repositoryRepository)
	adminService := service.NewAdminService(serviceService, adminRepository)
	authHandler := admin.NewAuthHandler(handler, adminService)
	mobileHandler := mobile.NewHandler(logger)
	userRepository := repository.NewUserRepository(repositoryRepository)
	userService := service.NewUserService(serviceService, userRepository)
	userHandler := mobile.NewUserHandler(mobileHandler, userService)
	httpServer := server.NewHTTPServer(logger, configConfig, jwtJWT, adminHandler, authHandler, userHandler)
	appApp := newApp(httpServer)
	return appApp, func() {
	}, nil
}

// wire.go:

var infraSet = wire.NewSet(infra.NewDB)

var repositorySet = wire.NewSet(repository.NewRepository, repository.NewUserRepository, repository.NewAdminRepository)

var serviceSet = wire.NewSet(service.NewServer, service.NewUserService, service.NewAdminService)

var adminHandlerSet = wire.NewSet(admin.NewHandler, admin.NewAuthHandler, admin.NewAdminHandler, admin.NewUserHandler)

var mobileHanderSet = wire.NewSet(mobile.NewHandler, mobile.NewUserHandler)

var serverSet = wire.NewSet(server.NewHTTPServer)

// build App
func newApp(
	httpServer *http.Server,

) *app.App {
	return app.NewApp(app.WithServer(httpServer), app.WithName("demo-server"))
}
