package middleware

import (
	"myapp/data"

	"github.com/Jonaxn/swiftness"
)

type Middleware struct {
	App    *swiftness.Swiftness
	Models data.Models
}
