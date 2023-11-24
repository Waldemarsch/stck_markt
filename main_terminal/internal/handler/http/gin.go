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
	api := router.Group("/api")
	{
		stocks := api.Group("/stocks")
		{
			stocksShort := stocks.Group("/short")
			{
				stocksShort.GET("/get", h.stocksShortGet)
			}
		}
	}

	return router
}
