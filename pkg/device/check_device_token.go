package device

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/krakensda/go-mqtt-api/pkg/common/models"
)

func (handler Handler) CheckDeviceToken(context *gin.Context) {
	body := models.TokenRequestBody{}
	err := context.BindJSON(&body)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "token is required",
			"data": nil,
		})
		return
	}

	var device models.Device
	result := handler.DB.Where(&models.Device{Token: body.Token}).First(&device)
	
	if result.Error != nil {
		
	}
}