package publisher_tests

import (
	"encoding/json"
	. "github.com/Andrew161644/currency_exchange/api/subcriber"
	. "github.com/Andrew161644/currency_exchange/api/task"
	"github.com/streadway/amqp"
	"log"
	"testing"
)

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

var queueName = "exchange"
var resultQueueName = "exchange_10"
var conn, amqpChannel, queue = Connect(localHost, queueName)
var localHost = "amqp://guest:guest@localhost:5672/"

func TestCanPablish(t *testing.T) {
	defer conn.Close()
	defer amqpChannel.Close()

	exchangeTask := CurrencyExchangeTask{
		ResultQueueName:     resultQueueName,
		Value:               34534,
		CurrentCurrencyName: "RUB",
		NewCurrencyName:     "EUR",
	}

	body, err := json.Marshal(exchangeTask)
	if err != nil {
		handleError(err, "Error encoding JSON")
	}

	err = amqpChannel.Publish("", queueName, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})

	if err != nil {
		log.Fatalf("Error publishing message: %s", err)
	}
	conn, amqpChannel, queue = Connect(localHost, exchangeTask.ResultQueueName)

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
	ListenOnce(
		res,
		messageChannel,
		Handle)
	amqpChannel.QueueDelete(exchangeTask.ResultQueueName, true, true, true)
}

func Handle(result chan CurrencyExchangeTask, task CurrencyExchangeTask) {
	defer close(result)

	log.Println("Result is ", task.Result)

	result <- task
}
