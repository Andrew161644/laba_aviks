package organization_status_client

import (
	"strconv"
)

type OrgStatusRequestModel struct {
	Name            string `json:"name"`
	BorrCap         int    `json:"borrCap"`
	OwnCap          int    `json:"ownCap"`
	BalanceCurr     int    `json:"balanceCurr"`
	AllCash         int    `json:"allCash"`
	LongTimeDuties  int    `json:"longTermDuties"`
	ShortTermDuties int    `json:"shortTermDuties"`
	ShortFinInv     int    `json:"shortFinInv"`
	ShortRec        int    `json:"shortRec"`
	SumMoney        int    `json:"sumMoney"`
}

func (obj *OrgStatusRequestModel) AsUrlValues() map[string][]string {
	return map[string][]string{
		"name":            {obj.Name},
		"borrCap":         {strconv.Itoa(obj.BorrCap)},
		"ownCap":          {strconv.Itoa(obj.OwnCap)},
		"balanceCurr":     {strconv.Itoa(obj.BalanceCurr)},
		"allCash":         {strconv.Itoa(obj.AllCash)},
		"longTermDuties":  {strconv.Itoa(obj.LongTimeDuties)},
		"shortTermDuties": {strconv.Itoa(obj.ShortTermDuties)},
		"shortFinInv":     {strconv.Itoa(obj.ShortFinInv)},
		"shortRec":        {strconv.Itoa(obj.ShortRec)},
		"sumMoney":        {strconv.Itoa(obj.SumMoney)},
	}
}
