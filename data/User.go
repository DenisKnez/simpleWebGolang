package data

import (
	"time"

	"github.com/google/uuid"
)

//User user entity
type User struct {
	ID        uuid.UUID  `json:"-"`
	Name      string     `json:"name"`
	Lastname  string     `json:"lastname"`
	Age       int16      `json:"age"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
	IsDeleted bool       `json:"-"`
}
