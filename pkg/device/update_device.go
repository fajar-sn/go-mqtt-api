package device

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/krakensda/go-mqtt-api/pkg/common/models"
)

func (handler Handler) UpdateDevice(context *gin.Context) {
	id := context.Param("id")
	body := models.DeviceRequestBody{}

	// gettings request's body
	err := context.BindJSON(&body)

	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var device models.Device
	result := handler.DB.First(&device, id)

	if result.Error != nil {
		context.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	device.Name = body.Name
	device.Token = body.Token
	handler.DB.Save(&device)

	context.JSON(http.StatusOK, gin.H{
		"message": "Successfully updated a device",
		"data": &device,
	})
}