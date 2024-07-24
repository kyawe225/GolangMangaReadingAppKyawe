package repositories

import (
	"realPj/mangaReadingApp/modules/shared/models"
	"realPj/mangaReadingApp/utils"
	"time"
)

type ICommentRepository interface {
	ListAllByChapter(chapter_id string) (*[]models.Comment, error)
	Save(id string, user_id string, message string) (*models.Comment, error)
}

type CommentRepository struct {
}

func (repository CommentRepository) ListAllByChapter(chapter_id string) (*[]models.Comment, error) {
	var comments []models.Comment
	query := `
		select comment.* from comment join chapter on chapter.id = comment.chapter_id where chapter.publish_url = ?
	`
	rows, err := utils.DB.Query(query, chapter_id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var model models.Comment
		rows.Scan(&model.Id, &model.ChapterId, &model.UserId, &model.Message, &model.CreatedAt, &model.UpdatedAt)
		comments = append(comments, model)
	}
	return &comments, nil
}

func (repository CommentRepository) Save(id string, user_id string, message string) (*models.Comment, error) {
	var chapterId string
	query := `
		select chapter.id as chapter_id form chapter where publish_url = ?;
	`

	row := utils.DB.QueryRow(query, id)

	err := row.Scan(&chapterId)

	if err != nil {
		return nil, err
	}
	commentKey := utils.GenerateUUIDV7()
	insertQuery := `
		insert into comment(id,chapter_id,user_id,message)
		values(?,?,?,?)
	`
	_, err = utils.DB.Exec(insertQuery, commentKey, chapterId, user_id, message)
	if err != nil {
		return nil, err
	}
	model := models.Comment{
		Id:        commentKey,
		ChapterId: chapterId,
		UserId:    user_id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return &model, nil
}
