package repositories

import (
	"realPj/mangaReadingApp/modules/shared/dtos"
	"realPj/mangaReadingApp/modules/shared/models"
	"realPj/mangaReadingApp/utils"

	"github.com/google/uuid"
)

type IMangaRepository interface {
	GetAll(user_id string) *[]models.Manga
	Save(model *models.Manga) error
	Update(id string, user_id string, model *models.Manga) error
	Delete(id string, user_id string) error
	FindById(id string) *models.Manga
	FindByIdDto(id string) *dtos.MangaDetailDto
	GetList() *[]dtos.MangaDto
}

type MangaRepository struct {
}

func (manga MangaRepository) GetAll(user_id string) *[]models.Manga {
	var rows []models.Manga
	query := `select mangas.*, category.*
	from mangas
	join manga_category on manga_category.manga_id = mangas.id
	join category on manga_category.category_id = category.id
	where user_id = $1`
	resultRows, err := utils.DB.Query(query, user_id)
	if err != nil {
		panic(err)
	}
	for resultRows.Next() {
		var row models.Manga
		err = resultRows.Scan(&row.Id, &row.Name, &row.Description, &row.PublishDate, &row.IsPublished, &row.PublishUrl, &row.UserId, &row.CreatedAt, &row.UpdatedAt, &row.Category.Id, &row.Category.Name, &row.Category.Description, &row.Category.CreatedAt, &row.Category.UpdatedAt)
		if err != nil {
			panic(err)
		}
		rows = append(rows, row)
	}
	return &rows
}

func (manga MangaRepository) FindById(id string) *models.Manga {
	var row models.Manga
	query := `select mangas.*,category.*
	from mangas
	join manga_category on manga_category.manga_id = mangas.id
	join category on manga_category.category_id = category.id
	where mangas.id = $1
	limit 1;`
	resultRow := utils.DB.QueryRow(query, id)
	err := resultRow.Scan(&row.Id, &row.Name, &row.Description, &row.PublishDate, &row.IsPublished, &row.PublishUrl, &row.UserId, &row.CreatedAt, &row.UpdatedAt, &row.Category.Id, &row.Category.Name, &row.Category.Description, &row.Category.CreatedAt, &row.Category.UpdatedAt)
	if err != nil {
		panic(err)
	}
	return &row
}

func (manga MangaRepository) Save(model *models.Manga) error {
	uniqueId, _ := uuid.NewV7()
	id := uniqueId.String()
	if model.IsPublished {
		model.PublishUrl = uuid.NewString()
	}
	query := `
		insert into mangas(id,name,description,publish_date,is_published,published_url,user_id)
		values ($1,$2,$3,$4,$5,$6,$7)
	`
	_, err := utils.DB.Exec(query, id, model.Name, model.Description, model.PublishDate, model.IsPublished, model.PublishUrl, model.UserId)

	if err != nil {
		panic(err)
	}

	uniqueId, _ = uuid.NewV7()
	mangaCategoryId := uniqueId.String()

	query = `insert into manga_category(id,manga_id,category_id) values(?,?,?)`

	_, err = utils.DB.Exec(query, mangaCategoryId, id, model.CategoryId)

	if err != nil {
		panic(err)
	}
	model.Id = id
	return nil
}

func (manga MangaRepository) Update(id string, user_id string, model *models.Manga) error {
	query := `
		update mangas set name=$1 , description =$2 , publish_date = $3, is_published = $4,published_url =$5
		where id = $6 and user_id = $7
	`
	_, err := utils.DB.Exec(query, model.Name, model.Description, model.PublishDate, model.IsPublished, model.PublishUrl, id, user_id)

	if err != nil {
		return err
	}

	query = `delete from manga_category where manga_id = $1;`

	_, err = utils.DB.Exec(query, id)
	if err != nil {
		return err
	}

	uniqueId, _ := uuid.NewV7()
	mangaCategoryId := uniqueId.String()
	query = `insert into manga_category(id,manga_id,category_id) values($1,$2,$3)`

	_, err = utils.DB.Exec(query, mangaCategoryId, id, model.CategoryId)

	if err != nil {
		panic(err)
	}

	return nil
}

func (manga MangaRepository) Delete(id string, user_id string) error {
	query := `
		delete from mangas
		where id = $1 and user_id = $2;
	`
	_, err := utils.DB.Exec(query, id, user_id)

	if err != nil {
		return err
	}
	return nil
}

/*
This is for public Api
*/
func (manga MangaRepository) GetList() *[]dtos.MangaDto {
	var rows []dtos.MangaDto
	var temp, temp1 string
	query := `select mangas.*, category.id as category_id,category.name as category_name
	from mangas
	join manga_category on manga_category.manga_id = mangas.id
	join category on manga_category.category_id = category.id
	where is_published = 1`
	resultRows, err := utils.DB.Query(query)
	if err != nil {
		panic(err)
	}
	for resultRows.Next() {
		var row dtos.MangaDto
		err = resultRows.Scan(&temp, &row.Name, &row.Description, &row.PublishDate, &row.IsPublished, &row.PublishUrl, &row.CreatedAt, &row.UpdatedAt, &temp1, &row.CategoryId, &row.CategoryName)
		if err != nil {
			panic(err)
		}
		rows = append(rows, row)
	}
	return &rows
}

func (manga MangaRepository) FindByIdDto(id string) *dtos.MangaDetailDto {
	var row dtos.MangaDetailDto
	temp, temp1 := "", ""
	query := `select mangas.*, category.id as category_id,category.name as category_name
	from mangas
	join manga_category on manga_category.manga_id = mangas.id
	join category on manga_category.category_id = category.id
	where published_url = $1
	limit 1;`
	resultRow := utils.DB.QueryRow(query, id)
	err := resultRow.Scan(&temp, &row.Manga.Name, &row.Manga.Description, &row.Manga.PublishDate, &row.Manga.IsPublished, &row.Manga.PublishUrl, &row.Manga.CreatedAt, &row.Manga.UpdatedAt, &temp1, &row.Manga.CategoryId, &row.Manga.CategoryName)
	if err != nil {
		panic(err)
	}

	query = `
		select * from chapter
		where manga_id = $1;
	`
	rows, err := utils.DB.Query(query, temp)
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var chapterDto dtos.ChapterDto
		err := rows.Scan(&temp1, &chapterDto.MangaPublishedId, &chapterDto.ChapterName, &chapterDto.Name, &chapterDto.Description, &chapterDto.IsPublished, &chapterDto.PublishUrl, &chapterDto.CreatedAt, &chapterDto.UpdatedAt, &temp)

		if err != nil {
			panic(err)
		}
		chapterDto.MangaPublishedId = id
		row.Chapters = append(row.Chapters, chapterDto)
	}

	return &row
}
