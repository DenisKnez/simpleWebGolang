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
	ReleaseDate *time.Time `json:"release_date"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-"`
	IsDeleted   bool       `json:"-"`
	PublisherID *uuid.UUID `json:"publisher_id"`
}
