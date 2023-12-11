package main

import (
	"github.com/Jonaxn/swiftness"
)

type application struct {
	App *swiftness.Swiftness
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
