package organization_status_client

import (
	"encoding/json"
	"log"
	"net/http"
)

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

func CallGetOrgStatusInfoFromData(uri string, data map[string][]string) (*OrgStatusResponseModel, error) {
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
