package server

import (
	"github.com/gin-gonic/gin"
	"github.com/zerokkcoder/indevsca/internal/handler/admin"
	"github.com/zerokkcoder/indevsca/internal/handler/mobile"
	"github.com/zerokkcoder/indevsca/pkg/config"
	"github.com/zerokkcoder/indevsca/pkg/jwt"
	"github.com/zerokkcoder/indevsca/pkg/log"
	"github.com/zerokkcoder/indevsca/pkg/reponse"
	"github.com/zerokkcoder/indevsca/pkg/server/http"
)

func NewHTTPServer(
	logger *log.Logger,
	config *config.Config,
	jwt *jwt.JWT,
	adminHandler *admin.AdminHandler,
	adminAuthHandler *admin.AuthHandler,
	mobileHandler *mobile.UserHandler,
) *http.Server {
	gin.SetMode(gin.DebugMode)
	engine := gin.Default()

	s := http.NewServer(
		engine,
		logger,
		http.WithHost(config.App.Host),
		http.WithPort(config.App.Port),
	)

	s.GET("/", func(ctx *gin.Context) {
		logger.InfoContext(ctx, "hello")
		reponse.HandleSuccess(ctx, map[string]interface{}{
			":)": "Thank you for using indevsca!",
		})
	})

	v1 := s.Group("/v1")
	{
		adminGroup := v1.Group("/admin")
		{
			noAuthGroup := adminGroup.Group("/")
			{
				noAuthGroup.POST("/login", adminAuthHandler.Login)
			}
		}

		// mobileGroup := v1.Group("/mobile")
		// {
		// 	mobileGroup.POST("/login", mobileHandler.Login)
		// }
	}

	return s
}
