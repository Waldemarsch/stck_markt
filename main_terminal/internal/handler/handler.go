package handler

import (
	http_custom "main_terminal/internal/handler/http"
	"main_terminal/internal/service"
)

type Handler struct {
	*http_custom.GinHttp
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		GinHttp: http_custom.NewGinHttp(s),
	}
}
