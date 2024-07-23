package chapters

import (
	"log"
	"net/http"
	"realPj/mangaReadingApp/modules/shared/dtos"
	"realPj/mangaReadingApp/modules/shared/models"
	"realPj/mangaReadingApp/modules/shared/repositories"

	"github.com/gin-gonic/gin"
)

type chapterController struct {
	repository repositories.IChapterRepository
}

func newChapterController(repository repositories.IChapterRepository) chapterController {
	return chapterController{
		repository: repository,
	}
}

func (controller *chapterController) index(context *gin.Context) {
	context.JSON(http.StatusOK, dtos.NewResponseDto("OK", "Successfully fetch chapters", controller.repository.GetAll()))
}

func (controller *chapterController) detail(context *gin.Context) {
	id := context.Param("id")

	if id == "" {
		context.JSON(http.StatusBadRequest, dtos.NewResponseDto("NG", "Request Format Is Wrong", "Wrong Id"))
		return
	}
	chapter, err := controller.repository.FindById(id)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, dtos.NewResponseDto("NG", "Request Format Is Wrong", "Not Found"))
		return
	}
	context.JSON(http.StatusOK, dtos.NewResponseDto("OK", "Successfully Fetched Mangas", chapter))
}

func (controller *chapterController) save(context *gin.Context) {

	var chapter models.Chapter

	err := context.ShouldBindBodyWithJSON(&chapter)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, dtos.NewResponseDto("NG", "Request Format Is Wrong", chapter))
		return
	}

	err = controller.repository.Save(&chapter)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, dtos.NewResponseDto("NG", "Something Wrong In Save", chapter))
		return
	}
	context.JSON(http.StatusCreated, dtos.NewResponseDto("Ok", "Saved Successfully", chapter))
}

func (controller *chapterController) update(context *gin.Context) {
	id := context.Param("id")

	if id == "" {
		log.Println("wrong request")
		context.JSON(http.StatusBadRequest, dtos.NewResponseDto("NG", "Request Format Is Wrong", id))
		return
	}

	var chapter models.Chapter

	err := context.ShouldBindBodyWithJSON(&chapter)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, dtos.NewResponseDto("NG", "Request Format Is Wrong", chapter))
		return
	}

	err = controller.repository.Update(id, &chapter)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, dtos.NewResponseDto("NG", "Something Wrong In Save", chapter))
		return
	}
	context.JSON(http.StatusOK, dtos.NewResponseDto("Ok", "Saved Successfully", chapter))
}

func (controller *chapterController) delete(context *gin.Context) {

	id := context.Param("id")

	if id == "" {
		log.Println("wrong request")
		context.JSON(http.StatusBadRequest, dtos.NewResponseDto("NG", "Request Format Is Wrong", id))
		return
	}

	err := controller.repository.Delete(id)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, dtos.NewResponseDto("NG", "Something Wrong In Save", id))
		return
	}
	context.JSON(http.StatusOK, dtos.NewResponseDto("Ok", "Saved Successfully", id))
}
