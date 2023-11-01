package internal

import (
	"log"
	"math/rand"
	"strconv"
	"testing"
)

var randomAccountEmail = "joseph." + strconv.Itoa(rand.Intn(1000)) + "@gmail.com"

func TestSignup(t *testing.T) {

	s, err := NewAccountService()

	if err != nil {
		t.Errorf("Connection fail")
		log.Fatal(err)
	}

	input := map[string]string{
		"name":        "Joseph Lastname",
		"email":       randomAccountEmail,
		"cpf":         "01234567890",
		"isPassenger": "true",
		"carplate":    "",
	}

	output, err := s.Signup(input)

	if err != nil {
		t.Errorf("Signup return a error %s", err)
	}

	if len(output) == 0 {
		t.Errorf("Not create account")
	}
}

func TestSignupExistingEmail(t *testing.T) {

	s, _ := NewAccountService()

	input := map[string]string{
		"name":        "Joseph Second Account",
		"email":       randomAccountEmail,
		"cpf":         "01234567890",
		"isPassenger": "true",
		"carplate":    "",
	}

	output, _ := s.Signup(input)

	if len(output) > 0 {
		t.Errorf("Can't create account with a existing e-mail")
	}
}

func TestSignupInvalidName(t *testing.T) {

	s, _ := NewAccountService()

	randomId := rand.Intn(1000)
	input := map[string]string{
		"name":        "Joseph",
		"email":       "joseph." + strconv.Itoa(randomId) + "@gmail.com",
		"cpf":         "01234567890",
		"isPassenger": "true",
		"carplate":    "",
	}

	output, _ := s.Signup(input)

	if len(output) > 0 {
		t.Errorf("Can't create account with invalid name")
	}
}

func TestSignupInvalidEmail(t *testing.T) {

	s, _ := NewAccountService()

	input := map[string]string{
		"name":        "Joseph Lastname",
		"email":       "joseph@gmail",
		"cpf":         "01234567890",
		"isPassenger": "true",
		"carplate":    "",
	}

	output, _ := s.Signup(input)

	if len(output) > 0 {
		t.Errorf("Can't create account with invalid e-mail")
	}
}

func TestSignupInvalidCarplate(t *testing.T) {

	s, _ := NewAccountService()

	randomId := rand.Intn(1000)
	input := map[string]string{
		"name":        "Joseph",
		"email":       "joseph." + strconv.Itoa(randomId) + "@gmail.com",
		"cpf":         "01234567890",
		"isPassenger": "true",
		"carplate":    "AAA",
	}

	output, _ := s.Signup(input)

	if len(output) > 0 {
		t.Errorf("Can't create account with invalid Carplate")
	}
}

func TestSignupInvalidCpf(t *testing.T) {

	s, _ := NewAccountService()

	randomId := rand.Intn(1000)
	input := map[string]string{
		"name":        "Joseph",
		"email":       "joseph." + strconv.Itoa(randomId) + "@gmail.com",
		"cpf":         "958.187.055-00",
		"isPassenger": "true",
		"carplate":    "",
	}

	output, _ := s.Signup(input)

	if len(output) > 0 {
		t.Errorf("Can't create account with invalid CPF")
	}
}
