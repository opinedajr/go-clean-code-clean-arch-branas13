package internal

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"testing"
)

var randomAccountEmail = "joseph." + strconv.Itoa(rand.Intn(1000)) + "@gmail.com"
var newAccountId = ""
var service AccountService

func TestNewAccountService(t *testing.T) {
	ad, err := NewAccountDao()

	if err != nil {
		t.Errorf("Fail initiate AccountDao")
		log.Fatal(err)
	}

	s, err := NewAccountService(*ad)

	if err != nil {
		t.Errorf("Connection fail")
		log.Fatal(err)
	}

	service = *s
}

func TestSignup(t *testing.T) {

	input := map[string]string{
		"name":        "Joseph Lastname",
		"email":       randomAccountEmail,
		"cpf":         "01234567890",
		"isPassenger": "true",
		"carplate":    "",
	}

	output, err := service.Signup(input)

	if err != nil {
		t.Errorf("Signup return a error %s", err)
	}

	if len(output) == 0 {
		t.Errorf("Not create account")
	}

	newAccountId = output
}

func TestSignupExistingEmail(t *testing.T) {

	input := map[string]string{
		"name":        "Joseph Second Account",
		"email":       randomAccountEmail,
		"cpf":         "01234567890",
		"isPassenger": "true",
		"carplate":    "",
	}

	output, _ := service.Signup(input)

	if len(output) > 0 {
		t.Errorf("Can't create account with a existing e-mail")
	}
}

func TestSignupInvalidName(t *testing.T) {

	randomId := rand.Intn(1000)
	input := map[string]string{
		"name":        "Joseph",
		"email":       "joseph." + strconv.Itoa(randomId) + "@gmail.com",
		"cpf":         "01234567890",
		"isPassenger": "true",
		"carplate":    "",
	}

	output, _ := service.Signup(input)

	if len(output) > 0 {
		t.Errorf("Can't create account with invalid name")
	}
}

func TestSignupInvalidEmail(t *testing.T) {

	input := map[string]string{
		"name":        "Joseph Lastname",
		"email":       "joseph@gmail",
		"cpf":         "01234567890",
		"isPassenger": "true",
		"carplate":    "",
	}

	output, _ := service.Signup(input)

	if len(output) > 0 {
		t.Errorf("Can't create account with invalid e-mail")
	}
}

func TestSignupInvalidCarplate(t *testing.T) {

	randomId := rand.Intn(1000)
	input := map[string]string{
		"name":        "Joseph",
		"email":       "joseph." + strconv.Itoa(randomId) + "@gmail.com",
		"cpf":         "01234567890",
		"isPassenger": "true",
		"carplate":    "AAA",
	}

	output, _ := service.Signup(input)

	if len(output) > 0 {
		t.Errorf("Can't create account with invalid Carplate")
	}
}

func TestSignupInvalidCpf(t *testing.T) {

	randomId := rand.Intn(1000)
	input := map[string]string{
		"name":        "Joseph",
		"email":       "joseph." + strconv.Itoa(randomId) + "@gmail.com",
		"cpf":         "958.187.055-00",
		"isPassenger": "true",
		"carplate":    "",
	}

	output, _ := service.Signup(input)

	if len(output) > 0 {
		t.Errorf("Can't create account with invalid CPF")
	}
}

func TestGetAccount(t *testing.T) {

	output, error := service.GetAccount(newAccountId)

	if error != nil {
		t.Errorf(error.Error())
	}

	if len(fmt.Sprint(output["name"])) == 0 {
		t.Errorf("Not retrieve account")
	}
}
