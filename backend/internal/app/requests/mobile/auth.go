package requests

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required" example:"johndoe"`     // 用户名
	Password string `json:"password" binding:"required" example:"password123"` // 密码
}

// RegisterRequest 注册请求
type RegisterRequest struct {
	Username string `json:"username" binding:"required" example:"johndoe"`             // 用户名
	Password string `json:"password" binding:"required" example:"password123"`         // 密码
	Email    string `json:"email" binding:"required,email" example:"john@example.com"` // 邮箱
}

// Response 通用响应
type Response struct {
	Code    int         `json:"code"`           // 状态码
	Message string      `json:"message"`        // 消息
	Data    interface{} `json:"data,omitempty"` // 数据
}

// TokenResponse 登录/注册成功响应
type TokenResponse struct {
	Token string   `json:"token"` // JWT令牌
	User  UserInfo `json:"user"`  // 用户信息
}

// UserInfo 用户信息
type UserInfo struct {
	ID       uint   `json:"id"`       // 用户ID
	Username string `json:"username"` // 用户名
	Email    string `json:"email"`    // 邮箱
}
