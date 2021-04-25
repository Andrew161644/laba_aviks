package main

import (
	"encoding/json"
	"github.com/Andrew161644/currency_exchange/api/exchanger"
	. "github.com/Andrew161644/currency_exchange/api/model"
	"github.com/streadway/amqp"
	"log"
	"os"
)

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}

}

func main() {
	conn, err := amqp.Dial(Config.AMQPConnectionURL)

	handleError(err, "Can't connect to AMQP")
	defer conn.Close()

	amqpChannel, err := conn.Channel()
	handleError(err, "Can't create a amqpChannel")

	defer amqpChannel.Close()

	queue, err := amqpChannel.QueueDeclare("add", true, false, false, false, nil)
	handleError(err, "Could not declare `add` queue")

	err = amqpChannel.Qos(1, 0, false)
	handleError(err, "Could not configure QoS")

	messageChannel, err := amqpChannel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "Could not register consumer")

	stopChan := make(chan bool)

	go func() {
		log.Printf("Consumer ready, PID: %d", os.Getpid())
		for d := range messageChannel {
			log.Printf("Received a message: %s", d.Body)

			addTask := &RequestCurrencyExchangeModel{}

			err := json.Unmarshal(d.Body, addTask)

			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
			}

			log.Println(exchanger.GetRate(
				addTask.Value,
				addTask.CurrentCurrencyName,
				addTask.NewCurrencyName,
				exchanger.GetEnvelope()))

			if err := d.Ack(false); err != nil {
				log.Printf("Error acknowledging message : %s", err)
			} else {
				log.Printf("Acknowledged message")
			}

		}
	}()

	// Stop for program termination
	<-stopChan
}
