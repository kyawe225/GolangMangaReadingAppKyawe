package models

import "time"

type ChapterPictures struct {
	Id          string
	PictureData string `json:"picture_data" require:"true"`
	ChapterId   string
	Serial      int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
