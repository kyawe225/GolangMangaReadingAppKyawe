package models

import "time"

type User struct {
	Id        string
	Name      string    `json:"name" require:"true"`
	Email     string    `json:"email" requrie:"true"`
	Password  string    `json:"password" require:"true"`
	BirthDate time.Time `json:"birthdate" require:"true"`
	Role      string    `json:"role" requrie:"true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
