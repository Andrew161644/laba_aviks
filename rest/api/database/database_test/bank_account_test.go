package database_test

import (
	"github.com/Andrew161644/avicks_laba/api/database/models"
	"log"
	"testing"
)

func TestCanGetAccountsBuUserId(t *testing.T) {
	var bankAccs, error = db.GetAllBankAccountsByUserId(models.UserModel{
		ID:       6,
		Name:     "",
		Password: "",
		RoleId:   1,
	})
	if error != nil {
		log.Fatal("Err: ", error)
	}
	log.Println(bankAccs)
}

func TestCanAddBankAccount(t *testing.T) {
	var id, error = db.AddBankAccount(models.BankAccount{
		Value:      8888,
		CurrencyId: 1,
		UserId:     1,
	})
	if error != nil {
		log.Fatalln(error)
	}
	log.Println(id)
}

func TestCanUpdateBankAccountById(t *testing.T) {
	var acc, error = db.UpdateBankAccountById(models.BankAccount{
		ID:         1,
		Value:      12,
		CurrencyId: 2,
		UserId:     2,
	})
	if error != nil {
		log.Fatalln(error)
	}
	log.Println(acc)
}

func TestCanDeleteBankAccount(t *testing.T) {
	var error = db.DeleteBankAccountById(models.BankAccount{
		ID: 1,
	})
	if error != nil {
		log.Fatalln(error)
	}
}
