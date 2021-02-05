package data

import (
	"time"

	"github.com/google/uuid"
)

//Book book entity
type Book struct {
	ID          uuid.UUID  `json:"-"`
	Title       string     `json:"title"`
	Author      string     `json:"author"`
	ReleaseDate time.Time  `json:"releaseDate"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-"`
	IsDeleted   bool       `json:"-"`
}
