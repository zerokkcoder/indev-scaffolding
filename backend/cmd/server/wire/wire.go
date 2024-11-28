//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/zerokkcoder/indevsca/internal/server"
	"github.com/zerokkcoder/indevsca/pkg/app"
	"github.com/zerokkcoder/indevsca/pkg/config"
	"github.com/zerokkcoder/indevsca/pkg/log"
	"github.com/zerokkcoder/indevsca/pkg/server/http"
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
		serverSet,
		newApp,
	))
}
