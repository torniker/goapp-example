package user

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/torniker/wrap"
	"github.com/torniker/wrap-example/db"
	"github.com/torniker/wrap-example/model"
	"github.com/torniker/wrap-example/schema"
	"github.com/torniker/wrap/logger"
)

// Handler handles /api/user routes
func Handler(c *wrap.Ctx) error {
	// if request method is POST call handleInsert
	c.Post(handleInsert)
	// if request method is GET call handleByID
	c.Get(handleByID)
	return nil
}

func handleByID(c *wrap.Ctx) error {
	userID, err := uuid.FromString(c.Request.Path().Next())
	if err != nil {
		logger.Warn(err)
		return c.NotFound()
	}
	user, err := db.UserByID(userID)
	if err != nil {
		return c.InternalError(err)
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

func handleInsert(c *wrap.Ctx) error {
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
	err = c.Prog.Call().Read("/api/user/" + id.String()).Flags(c.Request.Flags()).Bind(&user)
	if err != nil {
		return c.NotFound()
	}
	return c.JSON(user)
}
