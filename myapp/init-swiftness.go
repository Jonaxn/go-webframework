package main

import (
	"log"
	"os"

	"github.com/Jonaxn/swiftness"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// init swiftness
	cel := &swiftness.Swiftness{}
	err = cel.New(path)
	if err != nil {
		log.Fatal(err)
	}

	cel.AppName = "myapp"
	cel.Debug = true

	app := &application{
		App: cel,
	}

	return app
}
