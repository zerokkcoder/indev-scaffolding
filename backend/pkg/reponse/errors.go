package reponse

var (
	// common errors
	ErrSuccess             = newError(0, "ok")
	ErrInternalServerError = newError(500, "Internal Server Error")
	ErrNotFound            = newError(404, "Not Found")
	ErrTooManyRequests     = newError(429, "Too Many Requests")
	ErrBadRequest          = newError(400, "Bad Request")
	ErrServiceUnavailable  = newError(503, "Service Unavailable")

	// 认证相关错误
	ErrUnauthorized       = newError(401, "Unauthorized")
	ErrInvalidToken       = newError(1001, "Invalid token")
	ErrTokenExpired       = newError(1002, "Token expired")
	ErrInvalidCredentials = newError(1003, "Invalid credentials")
	ErrPermissionDenied   = newError(1004, "Permission denied")

	// 用户相关错误
	ErrUserNotFound     = newError(2001, "User not found")
	ErrUserDisabled     = newError(2002, "User disabled")
	ErrUserExists       = newError(2003, "User exists")
	ErrPasswordMismatch = newError(2004, "Password mismatch")
	ErrInvalidPassword  = newError(2005, "Invalid password")

	// 数据相关错误
	ErrDataNotFound   = newError(3001, "Data not found")
	ErrDataExists     = newError(3002, "Data exists")
	ErrDataInvalid    = newError(3003, "Data invalid")
	ErrDataConstraint = newError(3004, "Data constraint")
)
