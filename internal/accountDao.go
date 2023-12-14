package internal

import (
	"database/sql"
	"fmt"
)

type AccountDaoInterface interface {
	save(input map[string]string) (string, error)
	getByEmail(email string) (map[string]interface{}, error)
	getById(id string) (map[string]interface{}, error)
}

type AccountDao struct {
	db *sql.DB
}

func NewAccountDao() (*AccountDao, error) {
	connStr := "user=postgres dbname=cccat13 password=pg123 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &AccountDao{
		db: db,
	}, nil
}

func (ad *AccountDao) save(account map[string]string) (bool, error) {
	_, err := ad.db.Query(
		`INSERT INTO account (account_id, name, email, cpf, car_plate, is_passenger, is_driver, date, is_verified, verification_code) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		account["id"], account["name"], account["email"], account["cpf"], account["carplate"], true, false, account["date"], false, account["verificationCode"],
	)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (ad *AccountDao) getByEmail(email string) (map[string]interface{}, error) {
	rows, err := ad.db.Query(`SELECT account_id,name,email,cpf FROM account WHERE email = $1`, email)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var account_id, name, email, cpf string
		err := rows.Scan(
			&account_id,
			&name,
			&email,
			&cpf,
		)
		if err == nil {
			output := map[string]interface{}{
				"account_id": account_id,
				"name":       name,
				"email":      email,
				"cpf":        cpf,
			}
			return output, err
		}
	}

	return nil, fmt.Errorf("Account %s not found", email)
}

func (ad *AccountDao) getById(id string) (map[string]interface{}, error) {
	rows, err := ad.db.Query(`SELECT account_id,name,email,cpf FROM account WHERE account_id = $1`, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var account_id, name, email, cpf string
		err := rows.Scan(
			&account_id,
			&name,
			&email,
			&cpf,
		)
		if err == nil {
			output := map[string]interface{}{
				"account_id": account_id,
				"name":       name,
				"email":      email,
				"cpf":        cpf,
			}
			return output, err
		}
	}

	return nil, fmt.Errorf("Account %s not found", id)
}
