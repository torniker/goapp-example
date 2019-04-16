package main

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	app "github.com/torniker/goapp"
	"github.com/torniker/goapp/logger"
	"github.com/torniker/goapp/request"
	"github.com/torniker/goapp/response"
)

type (
	api struct {
		Name     string `json:"name"`
		Request  req    `json:"request"`
		Response resp   `json:"response"`
	}

	resp struct {
		Status int    `json:"status"`
		Type   string `json:"type"`
	}

	req struct {
		Action string `json:"action"`
		URI    string `json:"uri"`
		Input  string `json:"input"`
	}
)

func TestApp(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Fatal("Error loading .env file")
	}
	cfg := config{
		Environment:      app.Testing,
		PostgresAddr:     os.Getenv("POSTGRES_ADDRESS"),
		PostgresDB:       os.Getenv("POSTGRES_DB"),
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
	}
	a := app.New()
	err = setup(a, cfg)
	if err != nil {
		logger.Error(err)
	}
	file, err := ioutil.ReadFile("docs/api.json")
	if err != nil {
		logger.Error(err)
		return
	}
	tests := []api{}
	err = json.Unmarshal([]byte(file), &tests)
	if err != nil {
		logger.Error(err)
		return
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url, err := url.Parse(test.Request.URI)
			if err != nil {
				t.Fatal(err)
			}
			req := request.Req{}
			action := request.NewActionFromString(test.Request.Action)
			req.SetAction(action).SetPath(url)
			if test.Request.Input != "" {
				req.SetInput(strings.NewReader(test.Request.Input))
			}
			ctx := a.NewCtx(req, response.NewResponse())
			err = a.DefaultHandler(ctx)
			if err != nil {
				ctx.Error(err)
			}
			if ctx.Response.Status() != test.Response.Status {
				t.Errorf("Response Status code %v do not match: %v", ctx.Response.Status(), test.Response.Status)
			}
			o := ctx.Response.Output()
			resptype := reflect.ValueOf(o).Type().String()
			if resptype != test.Response.Type {
				t.Errorf("Response type %v do not match: %v", resptype, test.Response.Type)
			}
		})
	}

}
