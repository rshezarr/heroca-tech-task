package main

import (
	"log"
	"user_svc/internal/app"
)

func main() {
	app, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := app.RunApp(); err != nil {
		log.Fatal(err)
	}
}
