package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var neiroHost = "http://neiro_service:5000/"

func (app *Injection) NeiroServiceTest(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get(neiroHost + "/")
	if err != nil && resp != nil {
		fmt.Println(w, "Error")
		return
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	fmt.Fprintf(w, bodyString)
}
