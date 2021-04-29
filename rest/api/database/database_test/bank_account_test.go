package database_test

import (
	"github.com/Andrew161644/avicks_laba/api/database/models"
	"github.com/nu7hatch/gouuid"
	"log"
	"testing"
)

func TestCanGetAccountsBuUserId(t *testing.T) {
	var bankAccs, error = db.GetAllBankAccountsByUserId(models.BankAccount{
		UserId: 7,
	})
	if error != nil {
		log.Fatal("Err: ", error)
	}
	log.Println(bankAccs)
}

func TestCanAddBankAccount(t *testing.T) {
	var uuidVal, _ = uuid.NewV4()
	var id, error = db.AddBankAccount(models.BankAccount{
		ID:         uuidVal.String(),
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
		ID:         "5fa81258-a019-11eb-bcbc-0242ac130002",
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
		ID: "5fa81258-a019-11eb-bcbc-0242ac130002",
	})
	if error != nil {
		log.Fatalln(error)
	}
}
