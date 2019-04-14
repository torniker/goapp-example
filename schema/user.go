package schema

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/torniker/goapp-example/model"
)

type User struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Username  string    `db:"username" json:"username"`
	Password  string    `db:"password" json:"password"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

// Model converts schema User to model User
func (udb *User) Model() model.User {
	user := model.User{
		ID:        udb.ID,
		Username:  udb.Username,
		Password:  udb.Password,
		CreatedAt: udb.CreatedAt,
	}
	return user
}
