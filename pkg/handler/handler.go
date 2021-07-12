package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/p12s/2gis-catalog-api/pkg/service"
)

type Handler struct {
	services *service.Service
}

// NewHandler - конструируем хендлер передачей в него сервиса
func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}


// InitRoutes - инициализаруем пути и их хендлеры
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		city := api.Group("/city")
		{
			city.POST("/", h.createCity)
		}
	}
	return router
}