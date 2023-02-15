package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/krakensda/go-mqtt-api/pkg/common/models"
	"github.com/spf13/viper"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Thermometer struct {
	Id			uint		`json:"id"`
	Token		string		`json:"token"`
	Timestamp	time.Time	`json:"timestamp"`
	Data 		int			`json:"data"`
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, message mqtt.Message) {
	var payload Thermometer
	json.Unmarshal(message.Payload(), &payload)

	input := models.TelemetryRequestBody {
		Data: payload.Data,
		Token: payload.Token,
	}

	inputJson, err := json.Marshal(input)

	if err != nil {
		log.Fatal(err)
	}

	response, err := http.Post("http://localhost:3000/telemetries/", "application/json", bytes.NewBuffer(inputJson))

	if err != nil {
		log.Fatal(err)
	}

	var responseMap map[string]interface{}
	json.NewDecoder(response.Body).Decode(&responseMap)
	fmt.Println(responseMap)
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectionLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connection Lost: %s\n", err.Error())
}

func main() {
	broker := "tcp://broker.emqx.io:1883"
	clientId := "go_mqtt_example"
	options := mqtt.NewClientOptions()
	options.AddBroker(broker)
	options.SetClientID(clientId)
	options.SetDefaultPublishHandler(messagePubHandler)
	options.OnConnect = connectHandler
	options.OnConnectionLost = connectionLostHandler

	client := mqtt.NewClient(options)
	token := client.Connect()

	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	topic := "topic/temperature"
	token = client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic %s\n", topic)
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()
	port := viper.Get("PORT").(string)

	for {
		response, err := http.Get("http://localhost" + port + "/telemetries/last")

		if err != nil {
			fmt.Print(err.Error())
			break
		}

		responseData, err := ioutil.ReadAll(response.Body)

		if err != nil {
			log.Fatal(err)
		}

		var telemetries models.Telemetry
		json.Unmarshal(responseData, &telemetries)

		thermometer := Thermometer{
			Id: telemetries.ID + 1,
			Data: rand.Intn(100),
			Token: "qwertyuiop",
			Timestamp: time.Now(),
		}
		
		thermometerString, err := json.Marshal(&thermometer)

		if err != nil {
			log.Fatal(err)
		}

		token = client.Publish(topic, 0, false, thermometerString)
		token.Wait()

		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
		}
	}
}
