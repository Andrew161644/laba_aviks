package handlers

import (
	"fmt"
	cl "github.com/Andrew161644/avicks_laba/api/clients/organization_status_client"
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

}
