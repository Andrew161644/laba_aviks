package database_test

import (
	"github.com/Andrew161644/avicks_laba/api/database/models"
	"github.com/Andrew161644/avicks_laba/api/database/providers"
	"log"
	"testing"
)

const host = "localhost"

var db, err = providers.Connect(host, 5432, "postgres", "postgres", "postgres")

func TestGetAllRoles(t *testing.T) {

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

func TestAddRole(t *testing.T) {
	_, err = db.AddRole(models.Role{
		RoleName: "TestRole",
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestCanDeleteRole(t *testing.T) {
	role, err := db.GetRoleByName(models.Role{RoleName: "TestRole"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(role)
	err = db.DeleteRole(role.ID)
	if err != nil {
		log.Fatal(err)
	}
}
