package main

import (
	"log"
	"os"
)

func main() {
	app := NewApplication(":8080")

	if err := app.Run(app.Mount()); err != nil {
		log.Printf("Error: %s", err.Error())
		os.Exit(1)
	}
}
