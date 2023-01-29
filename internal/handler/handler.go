package handler

import (
	"bcc-hackathon-go/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/infoFromNumber", h.InfoFromNumber)
	router.POST("/hrCall", h.HrCall)
	router.POST("/techCall", h.TechCall)

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	return router
}
