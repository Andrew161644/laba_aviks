package handlers

import (
	"fmt"
	"github.com/Andrew161644/avicks_laba/api/broker/connection"
	. "github.com/Andrew161644/avicks_laba/api/broker/publisher"
	"github.com/Andrew161644/avicks_laba/api/broker/task"
	cl "github.com/Andrew161644/avicks_laba/api/clients/organization_status_client"
	"github.com/Andrew161644/avicks_laba/api/database/models"
	"log"
	"net/http"
)

func (app *Injection) NeiroServiceTest(w http.ResponseWriter, r *http.Request) {
	var res, err = cl.CallGetOrgStatusInfo(app.Conf.CalcUri, cl.OrgStatusRequestModel{
		Name:            "",
		BorrCap:         1,
		OwnCap:          1,
		BalanceCurr:     1,
		AllCash:         1,
		LongTimeDuties:  1,
		ShortTermDuties: 1,
		ShortFinInv:     1,
		ShortRec:        1,
		SumMoney:        1,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
	fmt.Fprintf(w, res.Report)
}

func (app *Injection) SendTaskExample(w http.ResponseWriter, r *http.Request) {

	var connectionIn, connectionOut = connection.CreateInOutConnectionForUser(app.Conf.RabbitHost, app.Conf.QueueName, models.UserModel{ID: 10})

	var channel = PublishSubscribe(
		connectionIn,
		connectionOut,
		task.BuildCurrencyExchangeTaskForUser(
			connectionOut.Queue.Name,
			15680,
			"USD",
			"RUB"),
		func(tasks chan task.CurrencyExchangeTask, exchangeTask task.CurrencyExchangeTask) {
			tasks <- exchangeTask
		})

	fmt.Fprintf(w, fmt.Sprintf("%f", (<-channel).Result))
	connectionIn.DeleteQueue()
	connectionOut.DeleteQueue()
}
