package admin

import "github.com/zerokkcoder/indevsca/internal/model"

// LoginRequest 管理员登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponseData 管理员登录响应
type LoginResponseData struct {
	Admin       *model.Admin `json:"admin"`
	AccessToken string       `json:"accessToken"`
}
