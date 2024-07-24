package repositories

import (
	"realPj/mangaReadingApp/modules/shared/models"
	"realPj/mangaReadingApp/utils"
)

type IBookMarkRepository interface {
	ListAll(user_id string) (*[]models.BookMark, error)
	BookMarkManga(mangaPublishId string, user_id string) error
	Delete(mangaPublishId string, user_id string) error
	ListAllWithoutUserId() (*[]models.BookMark, error)
}

type BookMarkRespotiroy struct {
}

func (repository BookMarkRespotiroy) BookMarkManga(mangaPublishId string, user_id string) error {
	query := `
		select id from mangas where publish_id = ?
	`
	var primaryKey string = ""
	rows := utils.DB.QueryRow(query, mangaPublishId)

	err := rows.Scan(&primaryKey)

	if err != nil {
		return err
	}
	id := utils.GenerateUUIDV7()

	insertQuery := `
		insert into bookmark(id,manga_id,user_id) values(?,?,?)
	`
	_, err = utils.DB.Exec(insertQuery, id, primaryKey, user_id)

	if err != nil {
		return err
	}

	return nil
}

func (repository BookMarkRespotiroy) ListAll(user_id string) (*[]models.BookMark, error) {
	var arr []models.BookMark
	query := `
	select bookmark.*,mangas.* from bookmark join mangas on mangas.id = bookmark.manga_id where user_id = ?
	`
	// var primaryKey
	rows, err := utils.DB.Query(query, user_id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var model models.BookMark
		rows.Scan(&model.Id, &model.MangaId, &model.UserId, &model.CreatedAt, &model.UpdatedAt, &model.Manga.Id, &model.Manga.Name, &model.Manga.Description, &model.Manga.PublishDate, &model.Manga.IsPublished, &model.Manga.PublishUrl, &model.Manga.UserId, &model.Manga.CreatedAt, &model.Manga.UpdatedAt)
		arr = append(arr, model)
	}
	return &arr, nil
}

func (repository BookMarkRespotiroy) ListAllWithoutUserId() (*[]models.BookMark, error) {
	var arr []models.BookMark
	var temp string
	query := `
	select bookmark.*,mangas.*,user.* from bookmark join mangas on mangas.id = bookmark.manga_id join users on users.id = bookmark.user_id;
	`
	// var primaryKey
	rows, err := utils.DB.Query(query)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var model models.BookMark
		rows.Scan(&model.Id, &model.MangaId, &model.UserId, &model.CreatedAt, &model.UpdatedAt, &model.Manga.Id, &model.Manga.Name, &model.Manga.Description, &model.Manga.PublishDate, &model.Manga.IsPublished, &model.Manga.PublishUrl, &model.Manga.UserId, &model.Manga.CreatedAt, &model.Manga.UpdatedAt, &model.User.Id, &model.User.Name, &model.User.Email, &temp, &model.User.BirthDate, &model.User.Role, &model.User.CreatedAt, &model.User.UpdatedAt)
		arr = append(arr, model)
	}
	return &arr, nil
}

func (repository BookMarkRespotiroy) Delete(mangaPublishId string, user_id string) error {
	query := `
		select id from mangas where publish_id = ?
	`
	var primaryKey string = ""
	rows := utils.DB.QueryRow(query, mangaPublishId)

	err := rows.Scan(&primaryKey)

	if err != nil {
		return err
	}

	deleteQ := `
		delete from bookmark where manga_id = ? and user_id =? ;	`
	_, err = utils.DB.Exec(deleteQ, primaryKey, user_id)

	if err != nil {
		return err
	}

	return nil
}
