package main

import (
	"calc/app"
	"log"
	"os"
)

func main() {
	app := app.Build()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
