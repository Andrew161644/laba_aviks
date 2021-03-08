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
	_, err = db.GetUserById(1)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestCanUpdateUser(t *testing.T) {
	var db, err = providers.Connect(host, 5432, "postgres", "postgres", "postgres")
	_, err = db.UpdateUser(1, models.User{Name: "West", RoleId: 1})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestCanDeleteUser(t *testing.T) {
	var db, err = providers.Connect(host, 5432, "postgres", "postgres", "postgres")
	err = db.DeleteUsers(1)
	if err != nil {
		t.Fatal(err.Error())
	}
}
