package dtos

import (
	"time"
)

type MangaDto struct {
	Name         string
	Description  string
	PublishDate  time.Time
	IsPublished  bool
	PublishUrl   string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	CategoryId   string
	CategoryName string
}

type MangaDetailDto struct {
	Manga    MangaDto
	Chapters []ChapterDto
}
