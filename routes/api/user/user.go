package user

import (
	"encoding/json"
	"net/url"
	"time"

	"github.com/gofrs/uuid"
	"github.com/torniker/goapp/app"
	"github.com/torniker/goapp/app/logger"
	"github.com/torniker/goapp/app/request"
	"github.com/torniker/goapp/app/response"
	"github.com/torniker/goapp/db"
	"github.com/torniker/goapp/model"
	"github.com/torniker/goapp/schema"
)

func Handler(c *app.Ctx) error {
	// if request method is POST call handleInsert
	c.POST(handleInsert)
	// if request method is GET call handleByID
	c.GET(handleByID)
	// else call handleElse
	c.ELSE(handleElse)
	// Do the logic built above
	c.Do()
	return nil
}

func handleElse(c *app.Ctx) error {
	return c.JSON([]string{})
}

func handleByID(c *app.Ctx) error {
	userID, err := uuid.FromString(c.CurrentPath.Next())
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
	decoder := json.NewDecoder(c.Request.Input())
	var uir userInsertRequest
	err := decoder.Decode(&uir)
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
	a := c.App
	u, err := url.Parse("/api/user/" + id.String())
	if err != nil {
		return err
	}
	subCtx := a.NewCtx(request.NewSub("GET", u, ""), response.NewSub())
	err = a.DefaultHandler(subCtx)
	if err != nil {
		return err
	}
	user := subCtx.Response.Output().(model.User)
	return c.JSON(user)
}
