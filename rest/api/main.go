package main

import (
	"flag"
	"github.com/Andrew161644/avicks_laba/api/config"
	"github.com/Andrew161644/avicks_laba/api/database/providers"
	"github.com/Andrew161644/avicks_laba/api/handlers"
	. "github.com/Andrew161644/avicks_laba/api/session"
	"log"
	"net/http"
	"time"
)

// для запуска открываем в терминале и вводим
//go run exchanger.go -host=localhost
// запускаем только в контейнере
func main() {
	var conf, error = config.GetConfig()
	if error != nil {
		log.Fatal(error)
	}
	var db, err = providers.Connect(conf.DbHost, conf.DbPort, conf.DbUsername, conf.DbPassword, conf.DbName)
	if err != nil {
		log.Fatal(err)
	}

	var (
		listen = flag.String("listen", conf.Port, "HTTP listen address")
	)
	flag.Parse()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("resources/static"))))
	var session = CreateNewUserSession()

	var app = handlers.Injection{
		DataBase:    db,
		UserSession: &session,
		Conf:        conf,
	}
	AddRoutes(app) // добавляет пути

	log.Println("Listen: ", listen)

	err = http.ListenAndServe(*listen, nil)
	time.Sleep(10 * time.Second)

	if err != nil {
		log.Fatal(err)
	}

}
