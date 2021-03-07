package main

import (
	"github.com/Andrew161644/avicks_laba/api/database"
	"log"
)

func main() {
	var db, err = database.Connect("postgres", "postgres", "postgres")
	if err != nil {
		log.Fatal(err)
	}
	db.Conn.Exec("INSERT INTO Users (name) values ($1)", "admin")
}
