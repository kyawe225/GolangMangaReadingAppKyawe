package models

import "time"

type Comment struct {
	Id        string
	ChapterId string
	UserId    string
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
