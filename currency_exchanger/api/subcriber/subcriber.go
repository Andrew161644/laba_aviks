package subcriber

import (
	"encoding/json"
	. "github.com/Andrew161644/currency_exchange/api/task"
	"github.com/streadway/amqp"
	"log"
)

type DecodeExchangeTask func(d amqp.Delivery) CurrencyExchangeTask
type CurrencyExchangeTaskHandler func(chan CurrencyExchangeTask, CurrencyExchangeTask)

func Decoder(d amqp.Delivery) CurrencyExchangeTask {
	addTask := &CurrencyExchangeTask{}

	err := json.Unmarshal(d.Body, addTask)
	if err := d.Ack(false); err != nil {
		log.Printf("Error acknowledging message : %s", err)
	} else {
		log.Printf("Acknowledged message")
	}

	if err != nil {
		log.Printf("Error decoding JSON: %s", err)
	}
	return *addTask
}

func HandleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func ListenMessageQueue(res chan CurrencyExchangeTask, messageChannel <-chan amqp.Delivery, handleFunk CurrencyExchangeTaskHandler) {
	for d := range messageChannel {
		go handleFunk(
			res,
			Decoder(d))
	}
}

func ListenOnce(res chan CurrencyExchangeTask, messageChannel <-chan amqp.Delivery, handleFunk CurrencyExchangeTaskHandler) {
	for d := range messageChannel {
		go handleFunk(
			res,
			Decoder(d))
		break
	}
}

// amqp://guest:guest@localhost:5672/
func Connect(connString string, queueName string) (*amqp.Connection, *amqp.Channel, amqp.Queue) {
	conn, err := amqp.Dial(connString)
	HandleError(err, "Can't connect to AMQP")
	amqpChannel, err := conn.Channel()
	HandleError(err, "Can't create a amqpChannel")
	queue, err := amqpChannel.QueueDeclare(queueName, true, false, false, false, nil)
	HandleError(err, "Could not declare `add` queue")
	err = amqpChannel.Qos(1, 0, false)
	HandleError(err, "Could not configure QoS")

	return conn, amqpChannel, queue
}
