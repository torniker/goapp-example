package routes

import (
	app "github.com/torniker/goapp"
	"github.com/torniker/goapp-example/routes/api"
)

func Handler(c *app.Ctx) error {
	if c.CurrentPath.Next() == "api" {
		c.Next(api.Handler)
		return nil
	}
	return c.NotFound()
}
