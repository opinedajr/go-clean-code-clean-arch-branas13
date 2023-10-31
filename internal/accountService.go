package internal

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"

	"github.com/google/uuid"
)

type AccountService struct {
	db *sql.DB
}

func NewAccountService() (*AccountService, error) {
	connStr := "user=postgres dbname=cccat13 password=pg123 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &AccountService{
		db: db,
	}, nil
}

func (s *AccountService) Signup(input map[string]string) (string, error) {
	var existingAccount int
	accountId := uuid.NewString()
	verificationCode := uuid.NewString()
	date := time.Now().UTC()
	s.db.QueryRow(`SELECT count(*) FROM account WHERE email = $1`, input["email"]).Scan(&existingAccount)
	if existingAccount > 0 {
		return "", fmt.Errorf("Account already exists")
	}
	_, err := s.db.Query(
		`INSERT INTO account (account_id, name, email, cpf, car_plate, is_passenger, is_driver, date, is_verified, verification_code) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		accountId, input["name"], input["email"], input["cpf"], input["carplate"], true, false, date, false, verificationCode,
	)
	if err != nil {
		return "", err
	}

	return accountId, nil
}
