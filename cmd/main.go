package main

import (
	"github.com/gin-gonic/gin"
	"github.com/krakensda/go-mqtt-api/pkg/common/db"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()
	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)
	app := gin.Default()
	db.Init(dbUrl)

	app.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"port": port,
			"dbUrl": dbUrl,
		})
	})

	app.Run(port)
}