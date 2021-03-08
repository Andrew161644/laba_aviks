package handlers

import (
	"fmt"
	"net/http"
)

func HelloPageHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello, %s", "People")
}
