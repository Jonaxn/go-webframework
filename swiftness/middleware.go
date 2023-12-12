package swiftness

import "net/http"

func (c *Swiftness) SessionLoad(next http.Handler) http.Handler {
	c.InfoLog.Println("SessionLoad called")
	return c.Session.LoadAndSave(next)
}
