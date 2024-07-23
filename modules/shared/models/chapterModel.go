package models

import "time"

type Chapter struct {
	Id          string
	MangaId     string `json:"manga_id" require:"true"`
	ChapterName string `json:"chapter_name" require:"true"`
	Name        string `json:"name" require:"true"`
	Description string `json:"description" require:"true"`
	IsPublished bool   `json:"is_published" require:"true"`
	PublishUrl  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ChapterPics []ChapterPictures `json:"chapter_pictures" require:"true"`
	UserId      string
	User        User
}
