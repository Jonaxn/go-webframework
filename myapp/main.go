package main

import (
	"myapp/data"
	"myapp/handlers"

	"github.com/jonaxn/swiftness"
)

type application struct {
	App      *swiftness.Swiftness
	Handlers *handlers.Handlers
	Models   data.Models
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
