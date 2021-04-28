package broker

import (
	"github.com/Andrew161644/avicks_laba/api/broker/connection"
	. "github.com/Andrew161644/avicks_laba/api/broker/publisher"
	"github.com/Andrew161644/avicks_laba/api/broker/task"
	"github.com/Andrew161644/avicks_laba/api/database/models"
	"log"
	"testing"
)

func Test_CanSubcribe(t *testing.T) {
	var connectionIn, connectionOut = connection.CreateInOutConnectionForUser("amqp://guest:guest@localhost:5672/", "exchange", models.UserModel{ID: 20})

	log.Println(connectionOut.Queue.Name)
	PublishSubscribe(
		connectionIn,
		connectionOut,
		task.BuildCurrencyExchangeTaskForUser(
			connectionOut.Queue.Name,
			145,
			"USD",
			"RUB"),
		func(tasks chan task.CurrencyExchangeTask, exchangeTask task.CurrencyExchangeTask) {
			log.Println(exchangeTask)

			tasks <- exchangeTask
		})
	connectionOut.DeleteQueue()
	connectionIn.DeleteQueue()
}
