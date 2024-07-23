package models

import "time"

type Category struct {
	Id          string
	Name        string `json:"name" required:"true"`
	Description string `json:"description" required:"true"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
