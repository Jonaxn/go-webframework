package main

import (
	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"

	"github.com/Jonaxn/swiftness"
)

type application struct {
	App        *swiftness.Swiftness
	Handlers   *handlers.Handlers
	Models     data.Models
	Middleware *middleware.Middleware
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
