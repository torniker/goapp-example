package routes

import (
	"github.com/torniker/goapp/app"
	"github.com/torniker/goapp/routes/api"
)

func Handler(c *app.Ctx) error {
	if c.CurrentPath.Next() == "api" {
		c.Next(api.Handler)
		return nil
	}
	return c.NotFound()
}
