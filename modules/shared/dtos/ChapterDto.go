package dtos

import (
	"realPj/mangaReadingApp/modules/shared/models"
	"time"
)

type ChapterDto struct {
	MangaPublishedId string
	ChapterName      string
	Name             string
	Description      string
	IsPublished      bool
	PublishUrl       string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type ChapterDetailDto struct {
	MangaPublishedId string
	ChapterName      string
	Name             string
	Description      string
	IsPublished      bool
	PublishUrl       string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	ChapterPics      []models.ChapterPictures
}
