package model

import (
	"time"

	"github.com/gofrs/uuid"
)

// User is struct for moving user data between micro-services and frontend
type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}
