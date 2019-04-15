package user

import (
	"time"

	"github.com/gofrs/uuid"
	app "github.com/torniker/goapp"
	"github.com/torniker/goapp-example/db"
	"github.com/torniker/goapp-example/model"
	"github.com/torniker/goapp-example/schema"
	"github.com/torniker/goapp/logger"
)

// Handler handles /api/user routes
func Handler(c *app.Ctx) error {
	// if request method is POST call handleInsert
	c.Create(handleInsert)
	// if request method is GET call handleByID
	c.Read(handleByID)
	return nil
}

func handleElse(c *app.Ctx) error {
	return c.JSON([]string{})
}

func handleByID(c *app.Ctx) error {
	userID, err := uuid.FromString(c.Request.Path().Next())
	if err != nil {
		logger.Warn(err)
		return c.NotFound()
	}
	user, err := db.UserByID(userID)
	if err != nil {
		logger.Error(err)
		return c.InternalError()
	}
	if user == nil {
		return c.NotFound()
	}
	return c.JSON(user.Model())
}

type userInsertRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func handleInsert(c *app.Ctx) error {
	var uir userInsertRequest
	err := c.Request.Bind(&uir)
	if err != nil {
		logger.Error(err)
		return err
	}
	id, err := uuid.NewV4()
	if err != nil {
		logger.Error(err)
		return err
	}
	userDB := schema.User{
		ID:        id,
		Username:  uir.Username,
		Password:  uir.Password,
		CreatedAt: time.Now(),
	}
	err = db.UserInsert(userDB)
	if err != nil {
		return err
	}
	var user model.User
	err = c.App.Call().Read("/api/user/" + id.String()).Flags(c.Request.Flags()).Bind(&user)
	if err != nil {
		return c.NotFound()
	}
	return c.JSON(user)
}
