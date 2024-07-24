package repositories

import (
	"log"
	"realPj/mangaReadingApp/modules/shared/dtos"
	"realPj/mangaReadingApp/modules/shared/models"
	"realPj/mangaReadingApp/utils"

	"github.com/google/uuid"
)

type IChapterRepository interface {
	GetAll(user_id string) *[]models.Chapter
	Save(model *models.Chapter) error
	Update(id string, user_id string, model *models.Chapter) error
	Delete(id string, user_id string) error
	FindById(id string) (*models.Chapter, error)
	GetListDto() *[]dtos.ChapterDto
	FindByIdDto(id string) (*dtos.ChapterDetailDto, error)
}

type ChapterRepository struct {
}

func (repository ChapterRepository) GetAll(user_id string) *[]models.Chapter {
	var chapters []models.Chapter

	query := `
		select * from chapter;
	`
	resultRows, err := utils.DB.Query(query)

	if err != nil {
		panic(err)
	}
	for resultRows.Next() {
		var chapter models.Chapter

		err = resultRows.Scan(&chapter.Id, &chapter.MangaId, &chapter.Name, &chapter.ChapterName, &chapter.Description, &chapter.IsPublished, &chapter.PublishUrl, &chapter.UserId, &chapter.CreatedAt, &chapter.UpdatedAt)

		if err != nil {
			panic(err)
		}

		chapters = append(chapters, chapter)
	}
	return &chapters
}

func (repository ChapterRepository) FindById(id string) (*models.Chapter, error) {
	var chapter models.Chapter
	query := `
		select * from chapter where id = ?;
	`

	rows, err := utils.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		err := rows.Scan(&chapter.Id, &chapter.MangaId, &chapter.ChapterName, &chapter.Name, &chapter.Description, &chapter.IsPublished, &chapter.PublishUrl, &chapter.UserId, &chapter.CreatedAt, &chapter.UpdatedAt)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
	}

	pictureQuery := `
		select * from chapter_pictures where chapter_id = ?;
	`
	row1, err := utils.DB.Query(pictureQuery, id)
	if err != nil {
		return nil, err
	}
	i := 0
	for row1.Next() {
		var chapterPics models.ChapterPictures
		err = row1.Scan(&chapterPics.ChapterId, &chapterPics.PictureData, &chapterPics.Id, &chapterPics.CreatedAt, &chapterPics.UpdatedAt, &chapterPics.Serial)
		if err != nil {
			return nil, err
		}
		chapter.ChapterPics = append(chapter.ChapterPics, chapterPics)
		i++
	}
	defer rows.Close()
	return &chapter, nil
}

func (repositroy ChapterRepository) Save(model *models.Chapter) error {
	transaction, errs := utils.DB.Begin()

	if errs != nil {
		panic(errs)
	}

	query := `
	insert into chapter(id,manga_id,chapter_name,name,description,is_published,publish_url,user_id)
	values(?,?,?,?,?,?,?,?);
	`
	id := utils.GenerateUUIDV7()
	var publishString string = ""
	if model.IsPublished {
		publishString = utils.GenerateUUIDV7()
	}

	_, err := transaction.Exec(query, id, model.MangaId, model.ChapterName, model.Name, model.Description, model.IsPublished, publishString, model.UserId)
	if err != nil {
		transaction.Rollback()
		panic(err)
	}
	for index, value := range model.ChapterPics {

		imgId := utils.GenerateUUIDV7()
		query = `
			insert into chapter_pictures(id,chapter_id,picture_data,serial,user_id)
			values(?,?,?,?,?)
		`
		_, err = transaction.Exec(query, imgId, id, value.PictureData, index+1, model.UserId)
		if err != nil {
			transaction.Rollback()
			log.Println(err)
			continue
		}
		model.ChapterPics[index].Id = imgId
		model.ChapterPics[index].ChapterId = id
		model.ChapterPics[index].Serial = int64(index + 1)
	}
	err = transaction.Commit()
	if err != nil {
		panic(err)
	}
	model.Id = id

	return nil
}

func (repository ChapterRepository) Update(id string, user_id string, model *models.Chapter) error {
	transaction, errs := utils.DB.Begin()

	if errs != nil {
		panic(errs)
	}
	query := `
		update chapter set chapter_name = ?, name=? , description =? , is_published = ?,publish_url =?
		where id = ? and user_id =?
	`
	if !model.IsPublished {
		model.PublishUrl = ""
	} else if model.IsPublished && utils.IsEmpty(&model.PublishUrl) {
		model.PublishUrl = utils.GenerateUUIDV7()
	}
	_, err := transaction.Exec(query, model.ChapterName, model.Name, model.Description, model.IsPublished, model.PublishUrl, id, user_id)

	if err != nil {
		transaction.Rollback()
		return err
	}

	deleteQuery := `
		delete from chapter_pictures where chapter_id = ? 
	`

	_, err = transaction.Exec(deleteQuery, id)

	if err != nil {
		return err
	}

	for index, value := range model.ChapterPics {
		uuidId, err := uuid.NewV7()

		if err != nil {
			transaction.Rollback()
			panic(err)
		}

		imgId := uuidId.String()
		query = `
			insert into chapter_pictures(id,chapter_id,picture_data,serial,user_id)
			values(?,?,?,?,?)
		`
		_, err = transaction.Exec(query, imgId, id, value.PictureData, index+1, user_id)
		if err != nil {
			transaction.Rollback()
			log.Println(err)
			continue
		}
		model.ChapterPics[index].Id = imgId
		model.ChapterPics[index].ChapterId = id
		model.ChapterPics[index].Serial = int64(index + 1)
	}
	transaction.Commit()
	if err != nil {
		panic(err)
	}
	return nil
}

func (repository ChapterRepository) Delete(id string, user_id string) error {

	transaction, errs := utils.DB.Begin()

	if errs != nil {
		panic(errs)
	}

	query := `
		delete from chapter
		where id = ? and user_id = ?;
	`
	_, err := transaction.Exec(query, id, user_id)

	if err != nil {
		transaction.Rollback()
		return err
	}
	deleteQuery := `
		delete from chapter_pictures where chapter_id = ? 
	`

	_, err = transaction.Exec(deleteQuery, id)

	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return nil
}

func (repository ChapterRepository) GetListDto() *[]dtos.ChapterDto {
	var chapters []dtos.ChapterDto
	var temp, temp2 string

	query := `
		select chapter.*,mangas.published_url as manga_published_id from chapter
		join mangas on mangas.id = chapter.manga_id  and mangas.is_published = 1
		where chapter.is_published = 1;
	`
	resultRows, err := utils.DB.Query(query)

	if err != nil {
		panic(err)
	}
	for resultRows.Next() {
		var chapter dtos.ChapterDto

		err = resultRows.Scan(&temp, &temp2, &chapter.Name, &chapter.ChapterName, &chapter.Description, &chapter.IsPublished, &chapter.PublishUrl, &temp2, &chapter.CreatedAt, &chapter.UpdatedAt, &chapter.MangaPublishedId)

		if err != nil {
			panic(err)
		}

		chapters = append(chapters, chapter)
	}
	return &chapters
}

func (repository ChapterRepository) FindByIdDto(id string) (*dtos.ChapterDetailDto, error) {
	var chapter dtos.ChapterDetailDto
	var temp, temp2 string
	query := `
		select chapter.*,mangas.published_url as manga_published_id from chapter
		join mangas on mangas.id = chapter.manga_id and mangas.is_published = 1
		where chapter.publish_url = ?;
	`

	rows, err := utils.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		err = rows.Scan(&temp, &temp2, &chapter.Name, &chapter.ChapterName, &chapter.Description, &chapter.IsPublished, &chapter.PublishUrl, &temp2, &chapter.CreatedAt, &chapter.UpdatedAt, &chapter.MangaPublishedId)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
	}

	pictureQuery := `
		select * from chapter_pictures where chapter_id = ?;
	`
	row1, err := utils.DB.Query(pictureQuery, temp)
	if err != nil {
		return nil, err
	}
	i := 0
	for row1.Next() {
		var chapterPics models.ChapterPictures
		err = row1.Scan(&chapterPics.ChapterId, &chapterPics.PictureData, &chapterPics.Id, &temp2, &chapterPics.CreatedAt, &chapterPics.UpdatedAt, &chapterPics.Serial)
		if err != nil {
			return nil, err
		}
		chapter.ChapterPics = append(chapter.ChapterPics, chapterPics)
		i++
	}
	defer rows.Close()
	return &chapter, nil
}
