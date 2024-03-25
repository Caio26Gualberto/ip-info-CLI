package main

import (
	"cli/app"
	"log"
	"os"
)

func main() {
	application := app.Gerar()

	if exception := application.Run(os.Args); exception != nil {
		log.Fatal(exception)
	}
}
