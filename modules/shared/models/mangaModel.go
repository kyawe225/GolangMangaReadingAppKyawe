package models

import (
	"time"
)

type Manga struct {
	Id          string
	Name        string    `json:"name" required:"true"`
	Description string    `json:"description" required:"true"`
	PublishDate time.Time `json:"publish_date" required:"true"`
	IsPublished bool      `json:"is_published" required:"true"`
	PublishUrl  string    `json:"published_url" required:"true"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Category    Category
	CategoryId  string `json:"category_id" require:"true"`
}
