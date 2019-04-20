package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/torniker/wrap"
	"github.com/torniker/wrap-example/db"
	"github.com/torniker/wrap-example/routes"
	"github.com/torniker/wrap/logger"
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
	p := wrap.New()
	err = setup(p, cfg)
	if err != nil {
		logger.Error(err)
	}
	p.StartHTTP(":8989")
	// a.StartCLI()
}

func setup(p *wrap.Prog, cfg config) error {
	p.DefaultHandler = routes.Handler
	p.Env = cfg.Environment
	addr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresAddr, cfg.PostgresDB)
	return db.New(addr)
}
