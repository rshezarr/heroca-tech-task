package main

import (
	"log"
	"salt_srv/internal/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := a.Start(); err != nil {
		log.Fatal(err)
	}
}
