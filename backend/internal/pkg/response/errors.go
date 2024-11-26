package response

import "errors"

var (
	// 通用错误
	ErrInternalServer   = errors.New("内部服务器错误")
	ErrInvalidParams    = errors.New("无效的参数")
	ErrNotFound         = errors.New("资源不存在")
	ErrTooManyRequests  = errors.New("请求过于频繁")
	ErrServiceUnavailable = errors.New("服务不可用")

	// 认证相关错误
	ErrUnauthorized      = errors.New("未经授权的访问")
	ErrInvalidToken      = errors.New("无效的令牌")
	ErrTokenExpired      = errors.New("令牌已过期")
	ErrInvalidCredentials = errors.New("无效的凭证")
	ErrPermissionDenied   = errors.New("权限不足")

	// 用户相关错误
	ErrUserNotFound      = errors.New("用户不存在")
	ErrUserDisabled      = errors.New("用户已被禁用")
	ErrUserExists        = errors.New("用户已存在")
	ErrPasswordMismatch  = errors.New("密码不匹配")
	ErrInvalidPassword   = errors.New("无效的密码")

	// 数据相关错误
	ErrDataNotFound      = errors.New("数据不存在")
	ErrDataExists        = errors.New("数据已存在")
	ErrDataInvalid       = errors.New("无效的数据")
	ErrDataConstraint    = errors.New("数据约束错误")
)
