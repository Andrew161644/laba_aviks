package clients_tests

import (
	"log"
	"testing"
)
import (
	. "github.com/Andrew161644/avicks_laba/api/clients/organization_status_client"
)

func TestMakeRequestOrgStatus(t *testing.T) {
	var err = CallGetOrgStatusInfo(OrgStatusRequestModel{
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
}
