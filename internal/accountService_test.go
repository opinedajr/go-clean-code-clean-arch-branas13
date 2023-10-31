package internal

import (
	"log"
	"math/rand"
	"strconv"
	"testing"
)

func TestSignup(t *testing.T) {

	s, err := NewAccountService()

	if err != nil {
		t.Errorf("Connection fail")
		log.Fatal(err)
	}

	randomId := rand.Intn(1000)
	input := map[string]string{
		"name":        "Joseph",
		"email":       "joseph." + strconv.Itoa(randomId) + "@gmail.com",
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
