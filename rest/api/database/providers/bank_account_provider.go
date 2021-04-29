package providers

import (
	"database/sql"
	"github.com/Andrew161644/avicks_laba/api/database/models"
)

func (db Database) AddBankAccount(account models.BankAccount) (string, error) {
	var id string
	query := `INSERT INTO bank_account (value, currencyid, userid) values ($1, $2, $3) RETURNING id`
	err := db.Conn.QueryRow(query, account.Value, account.CurrencyId, account.UserId).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (db Database) GetAllBankAccountsByUserId(account models.BankAccount) (models.BankAccountList, error) {
	var bankAccs = models.BankAccountList{}
	query := `SELECT * FROM bank_account WHERE userid = $1;`
	rows, err := db.Conn.Query(query, account.UserId)
	if err != nil {
		return bankAccs, err
	}
	for rows.Next() {
		var bankAcc = models.BankAccount{}
		err := rows.Scan(&bankAcc.ID, &bankAcc.CurrencyId, &bankAcc.UserId, &bankAcc.Value)
		if err != nil {
			return bankAccs, err
		}
		bankAccs.BankAccounts = append(bankAccs.BankAccounts, bankAcc)
	}
	return bankAccs, err
}

func (db Database) DeleteBankAccountById(account models.BankAccount) error {
	query := `DELETE FROM bank_account WHERE id = $1;`
	_, err := db.Conn.Exec(query, account.ID)
	switch err {
	case sql.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}

func (db Database) UpdateBankAccountById(account models.BankAccount) (models.BankAccount, error) {
	query := `UPDATE bank_account SET value=$1, currencyid=$2, userid=$3 WHERE id=$4 RETURNING id, value, currencyid, userid;`
	err := db.Conn.QueryRow(query, account.Value, account.CurrencyId, account.UserId, account.ID).Scan(&account.ID, &account.Value, &account.UserId, &account.CurrencyId)
	if err != nil {
		return models.BankAccount{}, err
	}
	return account, nil
}
