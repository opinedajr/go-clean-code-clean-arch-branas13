package internal

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type Api struct {
	listenAddress string
	accoutService Service
}

func NewApi(listenAddress string, accoutService Service) *Api {
	return &Api{
		listenAddress: listenAddress,
		accoutService: accoutService,
	}
}

func (api *Api) Run() {
	router := chi.NewRouter()
	router.Post("/signup", func(w http.ResponseWriter, r *http.Request) {

		var input map[string]string

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			fmt.Println("Error:", err.Error())
		}

		output, err := api.accoutService.Signup(input)

		if err != nil {
			fmt.Println("Error:", err.Error())
		}

		w.Write([]byte(output))
	})
	router.Get("/accounts/{id}", func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		account, err := api.accoutService.GetAccount(idParam)
		if err != nil {
			fmt.Println("Error:", err.Error())
		}
		accoutJson, _ := json.Marshal(account)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(accoutJson))
	})

	fmt.Println("JSON API Server running on port:", api.listenAddress)
	http.ListenAndServe(api.listenAddress, router)
}
