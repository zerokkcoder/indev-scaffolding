package admin

type AdminHandler struct {
	*Handler
}

func NewAdminHandler(handler *Handler) *AdminHandler {
	return &AdminHandler{
		Handler: handler,
	}
}
