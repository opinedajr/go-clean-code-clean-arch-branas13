package internal

import (
	"fmt"
	"regexp"

	_ "github.com/lib/pq"

	"github.com/google/uuid"
)

type Service interface {
	Signup(input map[string]string) (string, error)
	GetAccount(id string) (map[string]interface{}, error)
}

type AccountService struct {
	dao AccountDao
}

func NewAccountService(dao AccountDao) (*AccountService, error) {

	return &AccountService{
		dao: dao,
	}, nil
}

func (s *AccountService) GetAccount(id string) (map[string]interface{}, error) {
	output, err := s.dao.getById(id)
	if err != nil {
		return nil, err
	}

	return output, err
}

func (s *AccountService) Signup(input map[string]string) (string, error) {

	var validName = regexp.MustCompile(`^[A-Z][a-z]+ [A-Z][a-z]+$`)
	var validEmail = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	var validCarplate = regexp.MustCompile(`^[A-Z]{3}[0-9]{1}[A-Z]{1}[0-9]{2}$`)

	accountId := uuid.NewString()
	verificationCode := uuid.NewString()

	existingAccount, _ := s.dao.getByEmail(input["email"])

	if existingAccount != nil {
		return "", fmt.Errorf("Account already exists")
	}
	if validCpf := ValidateCpf(input["cpf"]); !validCpf {
		return "", fmt.Errorf("Invalid CPF")
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

	input["id"] = accountId
	input["date"] = "2023-12-07 23:28:18.203"
	input["verificationCode"] = verificationCode

	saved, err := s.dao.save(input)

	if !saved {
		return "", err
	}

	return accountId, nil
}
