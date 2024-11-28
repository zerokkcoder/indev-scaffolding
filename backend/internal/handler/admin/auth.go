package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zerokkcoder/indevsca/internal/request/admin"
	"github.com/zerokkcoder/indevsca/internal/service"
	"github.com/zerokkcoder/indevsca/pkg/reponse"
)

type AuthHandler struct {
	*Handler
	adminService service.AdminService
}

func NewAuthHandler(handler *Handler, adminService service.AdminService) *AuthHandler {
	return &AuthHandler{
		Handler:      handler,
		adminService: adminService,
	}
}

// Login 登录
func (h *AuthHandler) Login(ctx *gin.Context) {
	var req admin.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		reponse.HandleError(ctx, http.StatusBadRequest, reponse.ErrBadRequest, nil)
		return
	}

	adminData, token, err := h.adminService.Login(ctx, &req)
	if err != nil {
		reponse.HandleError(ctx, http.StatusUnauthorized, reponse.ErrUnauthorized, nil)
		return
	}
	reponse.HandleSuccess(ctx, admin.LoginResponseData{
		Admin:       adminData,
		AccessToken: token,
	})
}
