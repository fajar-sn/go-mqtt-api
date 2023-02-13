package device

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/krakensda/go-mqtt-api/pkg/common/models"
)

func (handler Handler) DeleteDevice(context *gin.Context) {
	id := context.Param("id")
	var device models.Device
	result := handler.DB.First(&device, id)

	if result.Error != nil {
		context.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	handler.DB.Delete(&device)
	context.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted a device",
	})
}