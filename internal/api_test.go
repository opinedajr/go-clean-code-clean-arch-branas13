package internal

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"testing"
)

func TestApiSignup(t *testing.T) {

	randomId := rand.Intn(1000)
	jsonData := `{
		"name": "Joseph Testapi",
		"email": "joseph` + strconv.Itoa(randomId) + `@gmail.com",
		"cpf": "01234567890",
		"isPassenger": "true",
		"carplate": "" 
	}`
	payload := bytes.NewBuffer([]byte(jsonData))
	resp, err := http.Post("http://localhost:8000/signup", "application/json", payload)
	if err != nil {
		t.Errorf(err.Error())
	}

	output, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf(err.Error())
	}


	

	defer resp.Body.Close()

	if len(output) == 0 {
		t.Errorf("Not create account")
	}

}
