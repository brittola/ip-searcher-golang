package main

import (
	"log"
	"os"
	"serversearch/app"
)

func main() {
	app := app.Build()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
