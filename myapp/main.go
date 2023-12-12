package main

import (
	"myapp/handlers"

	"github.com/Jonaxn/swiftness"
)

type application struct {
	App      *swiftness.Swiftness
	Handlers *handlers.Handlers
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
