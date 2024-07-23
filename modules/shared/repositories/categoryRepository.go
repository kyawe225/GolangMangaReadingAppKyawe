package repositories

import (
	"realPj/mangaReadingApp/modules/shared/models"
	"realPj/mangaReadingApp/utils"

	"github.com/google/uuid"
)

type ICategoryRepository interface {
	GetAll() *[]models.Category
	Save(model *models.Category) error
	Update(id string, model *models.Category) error
	Delete(id string) error
}

type CategoryRepository struct {
}

func (repository CategoryRepository) GetAll() *[]models.Category {
	var categories []models.Category
	query := `
		select * from category;
	`
	resultRows, err := utils.DB.Query(query)

	if err != nil {
		panic(err)
	}

	for resultRows.Next() {
		var category models.Category

		err := resultRows.Scan(&category.Id, &category.Name, &category.Description, &category.CreatedAt, &category.UpdatedAt)

		if err != nil {
			panic(err)
		}

		categories = append(categories, category)
	}
	return &categories
}

func (repository CategoryRepository) Save(model *models.Category) error {
	id, err := uuid.NewV7()

	if err != nil {
		panic(err)
	}

	query := `
		insert into category(id,name,description)
		values(?,?,?)
	`

	_, err = utils.DB.Exec(query, id.String(), model.Name, model.Description)

	if err != nil {
		panic(err)
	}
	model.Id = id.String()

	return nil
}

func (repository CategoryRepository) Update(id string, model *models.Category) error {
	query := `
		update category set name = ?,description = ?
		where id = ?;
	`

	_, err := utils.DB.Exec(query, model.Name, model.Description, id)

	if err != nil {
		panic(err)
	}

	return nil
}

func (manga CategoryRepository) Delete(id string) error {
	query := `
		delete from category
		where id = ?;
	`
	_, err := utils.DB.Exec(query, id)

	if err != nil {
		return err
	}
	return nil
}
