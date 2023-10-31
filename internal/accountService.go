package internal

import (
	"database/sql"
	"fmt"
	"regexp"
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
	var validName = regexp.MustCompile(`^[A-Z][a-z]+ [A-Z][a-z]+$`)
	var validEmail = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	var validCarplate = regexp.MustCompile(`^[A-Z]{3}[0-9]{1}[A-Z]{1}[0-9]{2}$`)

	accountId := uuid.NewString()
	verificationCode := uuid.NewString()
	date := time.Now().UTC()
	s.db.QueryRow(`SELECT count(*) FROM account WHERE email = $1`, input["email"]).Scan(&existingAccount)
	if existingAccount > 0 {
		return "", fmt.Errorf("Account already exists")
	}
	if !validName.MatchString(input["name"]) {
		return "", fmt.Errorf("Invalid name")
	}
	if !validEmail.MatchString(input["email"]) {
		return "", fmt.Errorf("Invalid e-mail")
	}
	if len(input["carplate"]) > 0 && !validCarplate.MatchString(input["carplate"]) {
		return "", fmt.Errorf("Invalid carplate")
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
