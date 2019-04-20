package api

import (
	"github.com/torniker/wrap"
	"github.com/torniker/wrap-example/routes/api/user"
)

func Handler(c *wrap.Ctx) error {
	switch c.Request.Path().Next() {
	case "user":
		c.Next(user.Handler)
		return nil
	default:
		return c.NotFound()
	}
}
