package db

import (
	"github.com/gofrs/uuid"
	"github.com/torniker/goapp/schema"
)

// UserInsert inserts user
func UserInsert(udb schema.User) error {
	_, err := db().NamedExec(`
		INSERT INTO users
			(id,
			username,
			password,
			created_at)
		VALUES
			(:id,
			:username,
			:password,
			:created_at)`, udb)
	if err != nil {
		return err
	}
	return nil
}

// UserByUsername gets user with provided username from postgres
func UserByUsername(username string) (*schema.User, error) {
	var udbs []schema.User
	err := db().Select(&udbs, "SELECT id, username, password, created_at FROM users WHERE username=$1", username)
	if err != nil {
		return nil, err
	}
	if len(udbs) == 0 {
		return nil, nil
	}
	return &udbs[0], nil
}

// UserByID gets user with provided id from postgres
func UserByID(id uuid.UUID) (*schema.User, error) {
	var udbs []schema.User
	err := db().Select(&udbs, "SELECT id, username, password, created_at FROM users WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	if len(udbs) == 0 {
		return nil, nil
	}
	return &udbs[0], nil
}
