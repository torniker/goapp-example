package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	app "github.com/torniker/goapp"
)

const key string = "postgres"

// New creates database object connects to it and stores in app.Store
func New(addr string) error {
	_ = pq.Efatal
	db, err := sqlx.Connect("postgres", addr)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	app.Instance().Store[key] = db
	return nil
}

func db() *sqlx.DB {
	return app.Instance().Store[key].(*sqlx.DB)
}
