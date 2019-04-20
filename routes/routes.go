package routes

import (
	"github.com/torniker/wrap"
	"github.com/torniker/wrap-example/routes/api"
)

func Handler(c *wrap.Ctx) error {
	if c.Request.Path().Next() == "api" {
		c.Next(api.Handler)
		return nil
	}
	return c.NotFound()
}
