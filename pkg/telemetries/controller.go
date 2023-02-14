package telemetries

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func RegisterRoutes(engine *gin.Engine, db *gorm.DB) {
	handler := &Handler{
		DB: db,
	}

	routes := engine.Group("/telemetries")
	routes.POST("/", handler.AddTelemetry)
	routes.GET("/", handler.GetTelemetries)
	routes.GET("/last", handler.GetLastTelemetry)
	routes.GET("/:id", handler.GetTelemetry)
	routes.PUT("/:id", handler.UpdateTelemetry)
	routes.DELETE("/:id", handler.DeleteTelemetry)
}