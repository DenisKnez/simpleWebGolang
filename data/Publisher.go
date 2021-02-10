package data

import (
	"time"

	"github.com/google/uuid"
)

//Publisher publisher entity
type Publisher struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	DateFounded time.Time  `json:"date_founded"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-"`
	IsDeleted   bool       `json:"-"`
}
