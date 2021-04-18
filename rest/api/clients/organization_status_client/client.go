package organization_status_client

import (
	"encoding/json"
	"log"
	"net/http"
)

var CalcUri = "http://neiro_service:5000/coefficient"

func CallGetOrgStatusInfo(uri string, rm OrgStatusRequestModel) (*OrgStatusResponseModel, error) {
	data := rm.AsUrlValues()
	log.Println(data)
	resp, err := http.PostForm(uri, data)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var res OrgStatusResponseModel

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &res, nil
}
