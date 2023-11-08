package handler

import "github.com/gin-gonic/gin"

type HttpHandler interface {
	InitRoutes() *gin.Engine
}

type Handler struct {
	HttpHandler
}
