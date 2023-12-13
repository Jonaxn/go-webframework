package swiftness

import "net/http"

func (c *Swiftness) SessionLoad(next http.Handler) http.Handler {
	return c.Session.LoadAndSave(next)
}
