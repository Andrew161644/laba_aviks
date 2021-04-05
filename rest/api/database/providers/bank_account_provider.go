package providers

import "github.com/Andrew161644/avicks_laba/api/database/models"

func (db Database) AddBankAccount(account models.BankAccount) (int, error) {
	var id int
	query := `INSERT INTO bank_account (value, currencyid, userid) values ($1, $2, $3) RETURNING id`
	err := db.Conn.QueryRow(query, account.Value, account.UserId, account.CurrencyId).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (db Database) GetAllBankAccountsByUserId(model models.UserModel) (models.BankAccountList, error) {
	var bankAccs = models.BankAccountList{}
	query := `SELECT * FROM bank_account WHERE userid = $1;`
	rows, err := db.Conn.Query(query, model)
	if err != nil {
		return bankAccs, err
	}
	for rows.Next() {
		var bankAcc = models.BankAccount{}
		err := rows.Scan(&bankAcc.ID, &bankAcc.Value, &bankAcc.UserId, &bankAcc.CurrencyId)
		if err != nil {
			return bankAccs, err
		}
		bankAccs.BankAccounts = append(bankAccs.BankAccounts, bankAcc)
	}
	return bankAccs, err
}
