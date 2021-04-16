package organization_status_client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var uri = "http://localhost:5000/coefficient"

func CallGetOrgStatusInfo(rm OrgStatusRequestModel) error {
	data := rm.AsUrlValues()
	log.Println(data)
	resp, err := http.PostForm(uri, data)

	if err != nil {
		log.Fatal(err)
		return err
	}

	//as struct example
	var res OrgStatusResponseModel

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println(res)
	return nil
}
