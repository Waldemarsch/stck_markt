package http_custom

import (
	"github.com/gin-gonic/gin"
	"main_terminal/internal/service"
)

type GinHttp struct {
	service *service.Service
}

func NewGinHttp(s *service.Service) *GinHttp {
	return &GinHttp{
		service: s,
	}
}

func (h *GinHttp) InitRoutes() *gin.Engine {
	router := gin.New()

	return router
}
