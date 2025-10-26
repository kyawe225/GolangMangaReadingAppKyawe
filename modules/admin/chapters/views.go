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

// @BasePath /chapters/
// @Summary List All Chapter
// @Schemes
// @Description List All Chapter
// @Tags Chapters
// @Accept json
// @Produce json
// @Failure 400 {object} dtos.ResponseDto[string]
// @Router /chapters [get]
func (controller *chapterController) index(context *gin.Context) {
	var userId string
	value, ok := context.Get("user")
	if ok {
		userId = value.(dtos.RegisterDto).Id
	}
	context.JSON(http.StatusOK, dtos.NewResponseDto("OK", "Successfully fetch chapters", controller.repository.GetAll(userId)))
}

// @BasePath /chapters/
// @Summary Detail Chapter
// @Schemes
// @Description Detail Chapter
// @Tags Chapters
// @Accept json
// @Produce json
// @Failure 400 {object} dtos.ResponseDto[string]
// @Router /chapters/{id} [get]
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

// @BasePath /chapters/
// @Summary Save Chapter
// @Schemes Chapter
// @Description Save Chapter
// @Tags Chapters
// @Accept json
// @Produce json
// @Failure 400 {object} dtos.ResponseDto[string]
// @Router /chapters [post]
func (controller *chapterController) save(context *gin.Context) {

	var chapter models.Chapter

	err := context.ShouldBindBodyWithJSON(&chapter)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, dtos.NewResponseDto("NG", "Request Format Is Wrong", chapter))
		return
	}

	value, ok := context.Get("user")
	if ok {
		chapter.UserId = value.(*dtos.RegisterDto).Id
	}

	err = controller.repository.Save(&chapter)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, dtos.NewResponseDto("NG", "Something Wrong In Save", chapter))
		return
	}
	context.JSON(http.StatusCreated, dtos.NewResponseDto("Ok", "Saved Successfully", chapter))
}

// @BasePath /chapters/
// @Summary Update Chapter
// @Schemes Chapter
// @Description Update Chapter
// @Tags Chapters
// @Accept json
// @Produce json
// @Failure 400 {object} dtos.ResponseDto[string]
// @Router /chapters/{id} [put]
func (controller *chapterController) update(context *gin.Context) {
	id := context.Param("id")

	if id == "" {
		log.Println("wrong request")
		context.JSON(http.StatusBadRequest, dtos.NewResponseDto("NG", "Request Format Is Wrong", id))
		return
	}

	var userId string
	value, ok := context.Get("user")
	if ok {
		userId = value.(*dtos.RegisterDto).Id
	}
	var chapter models.Chapter

	err := context.ShouldBindBodyWithJSON(&chapter)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, dtos.NewResponseDto("NG", "Request Format Is Wrong", chapter))
		return
	}

	err = controller.repository.Update(id, userId, &chapter)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, dtos.NewResponseDto("NG", "Something Wrong In Save", chapter))
		return
	}
	context.JSON(http.StatusOK, dtos.NewResponseDto("Ok", "Saved Successfully", chapter))
}

// @BasePath /chapters/
// @Summary Delete Chapter
// @Schemes
// @Description Delete Chapter
// @Tags Chapters
// @Accept json
// @Produce json
// @Success 200 {object} dtos.ResponseDto[string]
// @Failure 400 {object} dtos.ResponseDto[string]
// @Router /chapters/{id} [delete]
func (controller *chapterController) delete(context *gin.Context) {

	id := context.Param("id")
	var userId string
	value, ok := context.Get("user")
	if ok {
		userId = value.(*dtos.RegisterDto).Id
	}

	if id == "" {
		log.Println("wrong request")
		context.JSON(http.StatusBadRequest, dtos.NewResponseDto("NG", "Request Format Is Wrong", id))
		return
	}

	err := controller.repository.Delete(id, userId)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, dtos.NewResponseDto("NG", "Something Wrong In Save", id))
		return
	}
	context.JSON(http.StatusOK, dtos.NewResponseDto("Ok", "Saved Successfully", id))
}
