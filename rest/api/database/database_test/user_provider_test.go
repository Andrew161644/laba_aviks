package database_test

import (
	"github.com/Andrew161644/avicks_laba/api/database/models"
	"log"
	"testing"
)

// необходимо сначала поднять базу локально
func TestAddUser(t *testing.T) {
	_, err = db.AddUser(models.UserModel{RoleId: 1, Name: "Olgerd", Password: "olgerd"})
	if err != nil {
		log.Fatal(err.Error())
	}
}

func TestGetUserByNameAndPassword(t *testing.T) {
	user, err := db.GetUserByNameAndPassword("Andrew", "admin")
	if err != nil {
		log.Fatal(err)
	}
	var expected = models.UserModel{
		Name:     "Andrew",
		Password: "admin",
	}
	if user.Name != expected.Name || user.Password != user.Password {
		log.Fatal("Error")
	}
	log.Println(user)
}

func TestCanGetUser(t *testing.T) {
	_, err = db.GetUserById(1)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func TestGetUserByName(t *testing.T) {
	user, err := db.GetUserByName("ivan")
	if err != nil {
		log.Fatal(err.Error())
	}
	if user.Name != "ivan" {
		log.Fatal("Expected error")
	}
	log.Println(user)
}

func TestCanUpdateUser(t *testing.T) {
	_, err = db.UpdateUser(1, models.UserModel{Name: "West", RoleId: 1})
	if err != nil {
		log.Fatal(err.Error())
	}
}

func TestGetUserByRole(t *testing.T) {
	users, err := db.GetUserByRole(1)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(users)
}
