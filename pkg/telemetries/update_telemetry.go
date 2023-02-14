package telemetries

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/krakensda/go-mqtt-api/pkg/common/models"
)

func (handler Handler) UpdateTelemetry(context *gin.Context) {
	id := context.Param("id")
	body := models.TelemetryRequestBody{}
	err := context.BindJSON(&body)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "JSON not found",
			"data": nil,
		})
		return
	}

	var telemetry models.Telemetry
	result := handler.DB.First(&telemetry, id)

	if result.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Telemetry ID not found",
			"data": nil,
		})
		return
	}

	var device models.Device
	result = handler.DB.First(&device, body.DeviceID)

	if result.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Device ID not found",
			"data": nil,
		})
		return
	}

	telemetry.Data = body.Data
	telemetry.DeviceID = body.DeviceID
	handler.DB.Save(&telemetry)
	telemetry.Device = device

	context.JSON(http.StatusOK, gin.H{
		"message": "Successfully updated a telemetry",
		"data": &telemetry,
	})
}