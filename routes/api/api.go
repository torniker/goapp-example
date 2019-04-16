package api

import (
	app "github.com/torniker/goapp"
	"github.com/torniker/goapp-example/routes/api/user"
)

func Handler(c *app.Ctx) error {
	switch c.Request.Path().Next() {
	case "user":
		c.Next(user.Handler)
		return nil
	default:
		return c.NotFound()
	}
}
