package main

import (
	"fmt"
	"log"

	"github.com/opinedajr/go-clean-code-clean-arch-branas13/internal"
)

func main() {
	fmt.Println("Clean Code and Clean Arch Branas.io - 13")

	accountService, err := internal.NewAccountService()

	if err != nil {
		log.Fatal(err)
	}

	api := internal.NewApi(":8000", accountService)
	api.Run()
}
