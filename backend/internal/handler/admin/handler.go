package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/zerokkcoder/indevsca/pkg/log"
)

type Handler struct {
	logger *log.Logger
}

func NewHandler(
	logger *log.Logger,
) *Handler {
	return &Handler{
		logger: logger,
	}
}
func GetUserIdFromCtx(ctx *gin.Context) string {
	_, exists := ctx.Get("claims")
	if !exists {
		return ""
	}
	// TODO jwt
	// return v.(*jwt.MyCustomClaims).UserId
	return ""
}
