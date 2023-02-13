package device

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

	routes := engine.Group("/devices")
	routes.POST("/", handler.AddDevice)
	routes.GET("/", handler.GetDevices)
	routes.GET("/:id", handler.GetDevice)
}
