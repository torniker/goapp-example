package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/torniker/goapp/app"
	"github.com/torniker/goapp/app/logger"
	"github.com/torniker/goapp/db"
	"github.com/torniker/goapp/routes"
)

type config struct {
	Environment      string
	PostgresAddr     string
	PostgresDB       string
	PostgresUser     string
	PostgresPassword string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file")
		return
	}
	cfg := config{
		Environment:      os.Getenv("ENV"),
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
	a.StartHTTP(":8989")
	// a.StartCLI()
}

func setup(a *app.App, cfg config) error {
	a.DefaultHandler = routes.Handler
	a.Env = cfg.Environment
	addr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresAddr, cfg.PostgresDB)
	return db.New(addr)
}
