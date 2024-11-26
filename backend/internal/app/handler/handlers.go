package handler

// Handlers 所有HTTP处理器的集合
type Handlers struct {
	Auth *AuthHandler
}

// NewHandlers 创建处理器集合
func NewHandlers(
	auth *AuthHandler,
) *Handlers {
	return &Handlers{
		Auth: auth,
	}
}
