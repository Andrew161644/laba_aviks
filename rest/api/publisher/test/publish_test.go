package test

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"testing"
	"time"
)

type RequestCurrencyExchangeModel struct {
	Value               float64 `json:"value"`
	CurrentCurrencyName string  `json:"currentCurrencyName"`
	NewCurrencyName     string  `json:"newCurrencyName"`
}

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Test_Publish(t *testing.T) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	handleError(err, "Can't connect to AMQP")
	defer conn.Close()

	amqpChannel, err := conn.Channel()
	handleError(err, "Can't create a amqpChannel")

	defer amqpChannel.Close()

	queue, err := amqpChannel.QueueDeclare("add", true, false, false, false, nil)
	handleError(err, "Could not declare `add` queue")

	rand.Seed(time.Now().UnixNano())

	addTask := RequestCurrencyExchangeModel{
		Value:               180,
		CurrentCurrencyName: "RUB",
		NewCurrencyName:     "EUR",
	}
	body, err := json.Marshal(addTask)
	if err != nil {
		handleError(err, "Error encoding JSON")
	}

	err = amqpChannel.Publish("", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})

	if err != nil {
		log.Fatalf("Error publishing message: %s", err)
	}

}
