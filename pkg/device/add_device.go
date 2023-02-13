package device

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krakensda/go-mqtt-api/pkg/common/models"
)

func (handler Handler) AddDevice(context *gin.Context) {
	body := models.DeviceRequestBody{}

	// getting request's body
	err := context.BindJSON(&body)

	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var device models.Device
	device.Name = body.Name
	device.Token = body.Token
	result := handler.DB.Create(&device)

	if result.Error != nil {
		context.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Successfully add new device",
		"data":    &device,
	})
}
