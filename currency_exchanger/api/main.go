package main

import (
	"encoding/json"
	"github.com/Andrew161644/currency_exchange/api/exchanger"
	. "github.com/Andrew161644/currency_exchange/api/subcriber"
	. "github.com/Andrew161644/currency_exchange/api/task"
	"github.com/streadway/amqp"
	"log"
)

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

var dockerHost = "amqp://guest:guest@rabbitmq:5672/"
var localHost = "amqp://guest:guest@localhost:5672/"

func main() {
	var conn, amqpChannel, queue = Connect(dockerHost, "exchange")
	defer conn.Close()
	defer amqpChannel.Close()
	//amqpChannel.QueueDelete()
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

	res := make(chan CurrencyExchangeTask)
	ListenMessageQueue(
		res,
		messageChannel,
		Handle)

	num, opened := <-res

	if opened {
		log.Println(num)
	}
}

func Handle(result chan CurrencyExchangeTask, task CurrencyExchangeTask) {
	defer close(result)

	var conn, amqpChannel, _ = Connect(dockerHost, task.ResultQueueName)
	defer conn.Close()
	defer amqpChannel.Close()

	var res = exchanger.GetRate(task)
	task.Result = res
	log.Println(task)
	body, err := json.Marshal(task)
	if err != nil {
		handleError(err, "Error encoding JSON")
	}

	err = amqpChannel.Publish("", task.ResultQueueName, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})

	if err != nil {
		log.Fatalf("Error publishing message: %s", err)
	}

	result <- task
}
