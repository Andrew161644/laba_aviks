package database_test

import (
	"github.com/Andrew161644/avicks_laba/api/database/models"
	"github.com/Andrew161644/avicks_laba/api/database/providers"
	"testing"
)

// необходимо сначала поднять базу локально

const host = "localhost"

func TestAddUser(t *testing.T) {
	var db, err = providers.Connect(host, 5432, "postgres", "postgres", "postgres")
	_, err = db.AddUser(models.User{RoleId: 1, Name: "Andrew"})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestCanGetUser(t *testing.T) {
	var db, err = providers.Connect(host, 5432, "postgres", "postgres", "postgres")

	id, err := db.AddUser(models.User{RoleId: 1, Name: "Andrew"})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = db.GetUserById(id)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestCanUpdateUser(t *testing.T) {
	var db, err = providers.Connect(host, 5432, "postgres", "postgres", "postgres")

	id, err := db.AddUser(models.User{RoleId: 1, Name: "Andrew"})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = db.UpdateUser(id, models.User{Name: "West", RoleId: 1})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestCanDeleteUser(t *testing.T) {
	var db, err = providers.Connect(host, 5432, "postgres", "postgres", "postgres")

	id, err := db.AddUser(models.User{RoleId: 1, Name: "Andrew"})
	if err != nil {
		t.Fatal(err.Error())
	}

	err = db.DeleteUsers(id)
	if err != nil {
		t.Fatal(err.Error())
	}
}
