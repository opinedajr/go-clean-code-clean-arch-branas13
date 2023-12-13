package internal

import (
	"log"
	"math/rand"
	"strconv"
	"testing"

	"github.com/google/uuid"
)

var randomAccountEmailNro = "joseph." + strconv.Itoa(rand.Intn(1000)) + "@gmail.com"

func TestCreateAccount(t *testing.T) {
	ad, err := NewAccountDao()

	if err != nil {
		t.Errorf("Construct AccountDao fail")
		log.Fatal(err)
	}

	account := map[string]string{
		"id":               uuid.NewString(),
		"name":             "Joseph Dao Lastname",
		"email":            randomAccountEmailNro,
		"cpf":              "01234567890",
		"isPassenger":      "true",
		"carplate":         "",
		"date":             "2023-12-07 23:28:18.203",
		"verificationCode": uuid.NewString(),
	}

	saved, err := ad.save(account)

	if err != nil {
		t.Errorf("Create account return a error %s", err)
	}

	if !saved {
		t.Errorf("Fail to create account")
	}

	savedAccount, err := ad.getByEmail(account["email"])

	if err != nil {
		t.Errorf("Get create account error %s", err)
	}

	if savedAccount["email"] != account["email"] {
		t.Errorf("Expected email " + account["email"])
	}
	if savedAccount["name"] != account["name"] {
		t.Errorf("Expected name " + account["name"])
	}
	if savedAccount["cpf"] != account["cpf"] {
		t.Errorf("Expected cpf " + account["cpf"])
	}
}
