package models

import "time"

type BookMark struct {
	Id        string
	MangaId   string
	UserId    string
	CreatedAt time.Time
	UpdatedAt time.Time
	Manga     Manga
	User      User
}
