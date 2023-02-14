package telemetries

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/krakensda/go-mqtt-api/pkg/common/models"
)

func (handler Handler) GetTelemetries(context *gin.Context) {
	var telemetries []models.Telemetry
	result := handler.DB.Find(&telemetries)

	if result.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "No data available",
			"data": nil,
		})
		return
	}

	for i := 0; i < len(telemetries); i++ {
		var device models.Device
		handler.DB.Find(&device, telemetries[i].DeviceID)
		telemetries[i].Device = device
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Successfully get all telemetries",
		"data": gin.H{
			"telemetries": &telemetries,
		},
	})
}