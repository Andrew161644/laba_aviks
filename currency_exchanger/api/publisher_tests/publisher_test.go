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

func TestCanPablish(t *testing.T) {
	var conn, amqpChannel, queue = Connect("amqp://guest:guest@localhost:5672/", "exchange")
	defer conn.Close()
	defer amqpChannel.Close()

	exchangeTask := CurrencyExchangeTask{
		Value:               180,
		CurrentCurrencyName: "RUB",
		NewCurrencyName:     "EUR",
	}
	body, err := json.Marshal(exchangeTask)
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

	log.Printf("%f", exchangeTask.Value)
}
