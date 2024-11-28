package server

import (
	"github.com/gin-gonic/gin"
	"github.com/zerokkcoder/indevsca/pkg/config"
	"github.com/zerokkcoder/indevsca/pkg/log"
	"github.com/zerokkcoder/indevsca/pkg/server/http"
)

func NewHTTPServer(
	logger *log.Logger,
	config *config.Config,
) *http.Server {
	gin.SetMode(gin.DebugMode)
	s := http.NewServer(
		gin.Default(),
		logger,
		http.WithHost(config.App.Host),
		http.WithPort(config.App.Port),
	)
	return s
}
