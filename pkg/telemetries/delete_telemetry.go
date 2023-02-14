package telemetries

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/krakensda/go-mqtt-api/pkg/common/models"
)

func (handler Handler) DeleteTelemetry(context *gin.Context) {
	id := context.Param("id")
	var telemetry models.Telemetry
	result := handler.DB.First(&telemetry, id)

	if result.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Telemetry ID not found",
			"data": nil,
		})
		return
	}

	handler.DB.Delete(&telemetry)

	context.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted a telemetry",
	})
}