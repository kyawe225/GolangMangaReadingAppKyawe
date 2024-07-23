package dtos

import "time"

type LoginDto struct {
	Email    string `json:"email" requrie:"true"`
	Password string `json:"password" require:"true"`
}

type RegisterDto struct {
	Id        string
	Name      string    `json:"name" require:"true"`
	Email     string    `json:"email" requrie:"true"`
	Password  string    `json:"password" require:"true"`
	BirthDate time.Time `json:"birthdate" require:"true"`
	Role      string
}
type ProfileDto struct {
	Id        string    `json:"id"`
	Name      string    `json:"name" require:"true"`
	Email     string    `json:"email" requrie:"true"`
	Password  string    `json:"password" require:"true"`
	BirthDate time.Time `json:"birthdate" require:"true"`
}
