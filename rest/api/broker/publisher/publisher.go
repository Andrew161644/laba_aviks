package publisher

import (
	"encoding/json"
	. "github.com/Andrew161644/avicks_laba/api/broker/connection"
	. "github.com/Andrew161644/avicks_laba/api/broker/subcriber"
	. "github.com/Andrew161644/avicks_laba/api/broker/task"
	"github.com/streadway/amqp"
	"log"
)

func PublishSubscribe(connectionIn Connection, connectionOut Connection, exchangeTask CurrencyExchangeTask, handleFunk CurrencyExchangeTaskHandler) chan CurrencyExchangeTask {

	body, err := json.Marshal(exchangeTask)
	err = connectionIn.Channel.Publish("", connectionIn.Queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})

	if err != nil {
		log.Fatalf("Error publishing message: %s", err)
	}
	log.Println("Connection out name is: ", connectionOut.Queue.Name)
	messageChannel, err := connectionOut.Channel.Consume(
		connectionOut.Queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println(err)
	}
	res := make(chan CurrencyExchangeTask)
	ListenOnce(
		res,
		messageChannel,
		handleFunk)
	return res
}
