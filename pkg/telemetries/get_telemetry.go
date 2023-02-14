package telemetries

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/krakensda/go-mqtt-api/pkg/common/models"
)

func (handler Handler) GetTelemetry(context *gin.Context) {
	id := context.Param("id")
	var telemetry models.Telemetry
	result := handler.DB.First(&telemetry, id)

	if result.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "ID not found",
			"data": nil,
		})
		return
	}

	var device models.Device
	handler.DB.First(&device, telemetry.DeviceID)
	telemetry.Device = device

	context.JSON(http.StatusOK, gin.H{
		"message": "Successfully get a telemetry",
		"data": &telemetry,
	})
}