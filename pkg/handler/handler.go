package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/p12s/2gis-catalog-api/pkg/service"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/p12s/2gis-catalog-api/docs"
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		city := api.Group("/city")
		{
			city.POST("/", h.createCity)
		}
		company := api.Group("/company")
		{
			company.POST("/", h.createCompany)
		}
	}
	return router
}
