package comment

import (
	"net/http"
	"realPj/mangaReadingApp/modules/shared/dtos"
	"realPj/mangaReadingApp/modules/shared/repositories"
	"realPj/mangaReadingApp/utils"

	"github.com/gin-gonic/gin"
)

func newCommentController(repository repositories.ICommentRepository) commentController {
	return commentController{
		repo: repository,
	}
}

type commentController struct {
	repo repositories.ICommentRepository
}

func (comment *commentController) index(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		context.AbortWithStatusJSON(http.StatusBadRequest, dtos.NewResponseDto("OK", "No Id Found", "Invalid Id"))
	}
	data, err := comment.repo.ListAllByChapter(id)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, dtos.NewResponseDto("NG", "Some error occour", err))
		return
	}
	context.JSON(http.StatusOK, dtos.NewResponseDto("OK", "Successfully fetched comments", data))
}

func (comment *commentController) save(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		context.AbortWithStatusJSON(http.StatusBadRequest, dtos.NewResponseDto("NG", "No Id Found", "Invalid Id"))
		return
	}
	var message string
	context.ShouldBindBodyWithJSON(&message)
	userId := utils.ExtractLoggedUserId(context)

	model, err := comment.repo.Save(id, userId, message)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, dtos.NewResponseDto("NG", "Somthing Wrong", err))
		return
	}
	context.JSON(http.StatusOK, dtos.NewResponseDto("OK", "Successfully Inserted", (*model)))
}
