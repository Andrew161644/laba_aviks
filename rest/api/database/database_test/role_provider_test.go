package database

import (
	"github.com/Andrew161644/avicks_laba/api/database/providers"
	"log"
	"testing"
)

const host = "localhost"

func TestGetAllRoles(t *testing.T) {
	db, err := providers.Connect(host, 5432, "postgres", "postgres", "postgres")
	if err != nil {
		log.Fatal(err)
	}
	roles, err := db.GetAllRoles()
	if err != nil {
		log.Fatal(err)
	}
	if len(roles) == 0 {
		log.Fatal("Roles are empty")
	}
	log.Println(roles)
}
