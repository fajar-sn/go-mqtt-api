package telemetries

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/krakensda/go-mqtt-api/pkg/common/models"
)

func (handler Handler) AddTelemetry(context *gin.Context) {
	body := models.TelemetryRequestBody{}

	// getting request's body
	err := context.BindJSON(&body)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err,
			"data": nil,
		})
		return
	}

	var device models.Device
	result := handler.DB.Find(&device, body.DeviceID)

	if result.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": result.Error,
			"data": nil,
		})
		return
	}

	var telemetry models.Telemetry
	telemetry.Data = body.Data
	telemetry.DeviceID = body.DeviceID
	telemetry.Device = device
	result = handler.DB.Create(&telemetry)

	if result.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": result.Error,
			"data": nil,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Successfully add new telemetry",
		"data": gin.H{
			"id": telemetry.ID,
			"timestamp": telemetry.Timestamp,
			"data": telemetry.Data,
			"device_id": telemetry.DeviceID,
		},
	})
}
