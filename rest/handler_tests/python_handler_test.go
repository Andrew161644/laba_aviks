package handler_tests

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestNeiro(t *testing.T) {
	resp, err := http.Get("http://localhost:5000")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	log.Println(bodyString)
}
