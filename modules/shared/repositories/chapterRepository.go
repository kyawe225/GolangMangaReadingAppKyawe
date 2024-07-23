package repositories

import (
	"log"
	"realPj/mangaReadingApp/modules/shared/dtos"
	"realPj/mangaReadingApp/modules/shared/models"
	"realPj/mangaReadingApp/utils"

	"github.com/google/uuid"
)

type IChapterRepository interface {
	GetAll() *[]models.Chapter
	Save(model *models.Chapter) error
	Update(id string, model *models.Chapter) error
	Delete(id string) error
	FindById(id string) (*models.Chapter, error)
	GetListDto() *[]dtos.ChapterDto
	FindByIdDto(id string) (*dtos.ChapterDetailDto, error)
}

type ChapterRepository struct {
}

func (repository ChapterRepository) GetAll() *[]models.Chapter {
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

		err = resultRows.Scan(&chapter.Id, &chapter.MangaId, &chapter.Name, &chapter.ChapterName, &chapter.Description, &chapter.IsPublished, &chapter.PublishUrl, &chapter.CreatedAt, &chapter.UpdatedAt)

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
		err := rows.Scan(&chapter.Id, &chapter.MangaId, &chapter.ChapterName, &chapter.Name, &chapter.Description, &chapter.IsPublished, &chapter.PublishUrl, &chapter.CreatedAt, &chapter.UpdatedAt)
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
	query := `
	insert into chapter(id,manga_id,chapter_name,name,description,is_published,publish_url)
	values(?,?,?,?,?,?,?);
	`
	id := utils.GenerateUUIDV7()
	var publishString string = ""
	if model.IsPublished {
		publishString = utils.GenerateUUIDV7()
	}

	_, err := utils.DB.Exec(query, id, model.MangaId, model.ChapterName, model.Name, model.Description, model.IsPublished, publishString)
	if err != nil {
		panic(err)
	}
	for index, value := range model.ChapterPics {

		imgId := utils.GenerateUUIDV7()
		query = `
			insert into chapter_pictures(id,chapter_id,picture_data,serial)
			values(?,?,?,?)
		`
		_, err = utils.DB.Exec(query, imgId, id, value.PictureData, index+1)
		if err != nil {
			log.Println(err)
			continue
		}
		model.ChapterPics[index].Id = imgId
		model.ChapterPics[index].ChapterId = id
		model.ChapterPics[index].Serial = int64(index + 1)
	}

	model.Id = id

	return nil
}

func (repository ChapterRepository) Update(id string, model *models.Chapter) error {
	query := `
		update chapter set chapter_name = ?, name=? , description =? , is_published = ?,publish_url =?
		where id = ?
	`
	if !model.IsPublished {
		model.PublishUrl = ""
	} else if model.IsPublished && utils.IsEmpty(&model.PublishUrl) {
		model.PublishUrl = utils.GenerateUUIDV7()
	}
	_, err := utils.DB.Exec(query, model.ChapterName, model.Name, model.Description, model.IsPublished, model.PublishUrl, id)

	if err != nil {
		return err
	}

	deleteQuery := `
		delete from chapter_pictures where chapter_id = ? 
	`

	_, err = utils.DB.Exec(deleteQuery, id)

	if err != nil {
		return err
	}

	for index, value := range model.ChapterPics {
		uuidId, err := uuid.NewV7()

		if err != nil {
			panic(err)
		}

		imgId := uuidId.String()
		query = `
			insert into chapter_pictures(id,chapter_id,picture_data,serial)
			values(?,?,?,?)
		`
		_, err = utils.DB.Exec(query, imgId, id, value.PictureData, index+1)
		if err != nil {
			log.Println(err)
			continue
		}
		model.ChapterPics[index].Id = imgId
		model.ChapterPics[index].ChapterId = id
		model.ChapterPics[index].Serial = int64(index + 1)
	}

	return nil
}

func (repository ChapterRepository) Delete(id string) error {
	query := `
		delete from chapter
		where id = ?;
	`
	_, err := utils.DB.Exec(query, id)

	if err != nil {
		return err
	}
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

		err = resultRows.Scan(&temp, &temp2, &chapter.Name, &chapter.ChapterName, &chapter.Description, &chapter.IsPublished, &chapter.PublishUrl, &chapter.CreatedAt, &chapter.UpdatedAt, &chapter.MangaPublishedId)

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
		err = rows.Scan(&temp, &temp2, &chapter.Name, &chapter.ChapterName, &chapter.Description, &chapter.IsPublished, &chapter.PublishUrl, &chapter.CreatedAt, &chapter.UpdatedAt, &chapter.MangaPublishedId)
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
