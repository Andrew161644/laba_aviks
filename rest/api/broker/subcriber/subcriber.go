package subcriber

import (
	"encoding/json"
	. "github.com/Andrew161644/avicks_laba/api/broker/task"
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
