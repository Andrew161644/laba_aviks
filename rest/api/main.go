package main

import (
	"flag"
	"github.com/Andrew161644/avicks_laba/api/database/providers"
	. "github.com/Andrew161644/avicks_laba/api/handlers"
	"log"
	"net/http"
)

// для запуска открываем в терминале и вводим
//go run main.go -host=localhost
func main() {
	var host = flag.String("host", "db", "HTTP listen address")
	flag.Parse()
	log.Println("Use host: " + *host)
	var db, err = providers.Connect(*host, 5432, "postgres", "postgres", "postgres")
	if err != nil {
		log.Fatal(err)
	}
	var app = Injection{
		DataBase: db,
	}
	var (
		listen = flag.String("listen", ":8080", "HTTP listen address")
	)
	flag.Parse()

	HandleExample(app)

	err = http.ListenAndServe(*listen, nil)
	if err != nil {
		log.Fatal(err)
	}
}
func HandleExample(app Injection) {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/page", app.HelloPageHandler)
}
