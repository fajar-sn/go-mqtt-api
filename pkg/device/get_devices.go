package device

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/krakensda/go-mqtt-api/pkg/common/models"
)

func (handler Handler) GetDevices(context *gin.Context) {
	var devices []models.Device
	result := handler.DB.Find(&devices)

	if result.Error != nil {
		context.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Successfully get all devices",
		"data": &devices,
	})
}