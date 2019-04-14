package api

import (
	"github.com/torniker/goapp/app"
	"github.com/torniker/goapp/routes/api/user"
)

func Handler(c *app.Ctx) error {
	switch c.CurrentPath.Next() {
	case "user":
		c.Next(user.Handler)
		return nil
	default:
		return c.NotFound()
	}
}
