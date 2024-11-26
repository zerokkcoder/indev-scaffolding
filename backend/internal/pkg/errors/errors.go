package errors

import (
	"fmt"
	"net/http"
)

// ErrorCode 错误码类型
type ErrorCode int

const (
	// ErrInternal 内部错误
	ErrInternal ErrorCode = iota + 1000
	// ErrInvalidParam 无效参数
	ErrInvalidParam
	// ErrUserNotFound 用户不存在
	ErrUserNotFound
	// ErrUserExists 用户已存在
	ErrUserExists
	// ErrInvalidPassword 密码错误
	ErrInvalidPassword
	// ErrUnauthorized 未授权
	ErrUnauthorized
	// ErrForbidden 禁止访问
	ErrForbidden
)

type AppError struct {
	Code    ErrorCode
	Message string
	Err     error
}

// Error 实现 error 接口
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// Unwrap 实现 errors.Unwrap 接口
func (e *AppError) Unwrap() error {
	return e.Err
}

// StatusCode 获取HTTP状态码
func (e *AppError) StatusCode() int {
	switch e.Code {
	case ErrInvalidParam:
		return http.StatusBadRequest
	case ErrUserNotFound:
		return http.StatusNotFound
	case ErrUserExists:
		return http.StatusConflict
	case ErrInvalidPassword:
		return http.StatusUnauthorized
	case ErrUnauthorized:
		return http.StatusUnauthorized
	case ErrForbidden:
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}

// NewAppError 创建应用错误
func NewAppError(code ErrorCode, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}
