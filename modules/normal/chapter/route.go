package chapter

import (
	"log"
	"net/http"
	"realPj/mangaReadingApp/modules/shared/dtos"
	"realPj/mangaReadingApp/modules/shared/repositories"

	"github.com/gin-gonic/gin"
)

type UserchapterController struct {
	repository repositories.IChapterRepository
}

func newChapterController(repository repositories.IChapterRepository) UserchapterController {
	return UserchapterController{
		repository: repository,
	}
}

func (controller *UserchapterController) index(context *gin.Context) {
	context.JSON(http.StatusOK, dtos.NewResponseDto("OK", "Successfully fetch chapters", controller.repository.GetListDto()))
}

func (controller *UserchapterController) detail(context *gin.Context) {
	id := context.Param("id")

	if id == "" {
		context.JSON(http.StatusBadRequest, dtos.NewResponseDto("NG", "Request Format Is Wrong", "Wrong Id"))
		return
	}
	chapter, err := controller.repository.FindByIdDto(id)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, dtos.NewResponseDto("NG", "Request Format Is Wrong", "Not Found"))
		return
	}
	context.JSON(http.StatusOK, dtos.NewResponseDto("OK", "Successfully Fetched Mangas", chapter))
}
