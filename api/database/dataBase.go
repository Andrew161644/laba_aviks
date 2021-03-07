package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

const (
	HOST = "db"
	PORT = 5432
)

type Database struct {
	Conn *sql.DB
}

func Initialize(username, password, database string) (Database, error) {

	db := Database{}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, username, password, database)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return db, err
	}

	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}

	log.Println("Database connection established")
	return db, nil
}

func Connect(username, password, database string) (*Database, error) {
	var err error
	for i := 0; i < 10; i++ {
		var dbs, err = Initialize(username, password, database)
		if err == nil {
			log.Println("Connected")
			return &dbs, nil
		} else {
			log.Println("Waits")
			time.Sleep(1000 * time.Millisecond)
		}
	}
	return nil, err
}
