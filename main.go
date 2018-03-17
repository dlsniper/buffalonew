package main

import (
	"log"

	"github.com/dlsniper/buffalonew/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
