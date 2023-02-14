package main

import (
	"github.com/gin-gonic/gin"
	"github.com/krakensda/go-mqtt-api/pkg/device"
	"github.com/krakensda/go-mqtt-api/pkg/telemetries"
	"github.com/krakensda/go-mqtt-api/pkg/common/db"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()
	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)
	app := gin.Default()
	handler := db.Init(dbUrl)
	device.RegisterRoutes(app, handler)
	telemetries.RegisterRoutes(app, handler)
	app.Run(port)
}