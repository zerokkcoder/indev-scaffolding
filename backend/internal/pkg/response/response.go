package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CustomError 错误响应
type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *CustomError) Error() string {
	return e.Message
}

// NewError 创建新的错误
func NewError(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

// Error 错误响应
func Error(c *gin.Context, err error) {
	if e, ok := err.(*CustomError); ok {
		c.JSON(http.StatusOK, Response{
			Code:    e.Code,
			Message: e.Message,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    50000,
		Message: err.Error(),
	})
}
