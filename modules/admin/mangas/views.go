package mangas

import (
	"log"
	"net/http"
	"realPj/mangaReadingApp/modules/shared/dtos"
	"realPj/mangaReadingApp/modules/shared/models"
	"realPj/mangaReadingApp/modules/shared/repositories"

	"github.com/gin-gonic/gin"
)

func newMangaController(repo repositories.IMangaRepository) MangaController {
	return MangaController{
		repository: repo,
	}
}

type MangaController struct {
	repository repositories.IMangaRepository
}

// @BasePath /manga/
// @Summary List All Mangas
// @Schemes
// @Description List all the mangas
// @Tags Mangas
// @Accept json
// @Produce json
// @Failure 400 {object} dtos.ResponseDto[string]
// @Router /mangas [get]
func (controller *MangaController) index(context *gin.Context) {
	var userId string
	value, ok := context.Get("user")
	if ok {
		userId = value.(*dtos.RegisterDto).Id
	}
	context.JSON(http.StatusOK, dtos.NewResponseDto("OK", "Successfully Fetched Mangas", controller.repository.GetAll(userId)))
}

// @BasePath /manga/
// @Summary Detail Manga
// @Schemes
// @Description Detail the manga
// @Tags Mangas
// @Accept json
// @Produce json
// @Failure 400 {object} dtos.ResponseDto[string]
// @Router /mangas/{id} [get]
func (controller *MangaController) detail(context *gin.Context) {
	id := context.Param("id")

	if id == "" {
		context.JSON(http.StatusOK, dtos.NewResponseDto("NG", "Request Format Is Wrong", "Wrong Id"))
		return
	}
	context.JSON(http.StatusOK, dtos.NewResponseDto("OK", "Successfully Fetched Mangas", controller.repository.FindById(id)))
}

// @BasePath /manga/
// @Summary Create Manga
// @Schemes Manga
// @Description Create Manga
// @Tags Mangas
// @Accept json
// @Produce json
// @Failure 400 {object} dtos.ResponseDto[string]
// @Router /mangas [post]
func (controller *MangaController) save(context *gin.Context) {

	var manga models.Manga

	err := context.ShouldBindBodyWithJSON(&manga)
	value, ok := context.Get("user")
	if ok {
		manga.UserId = value.(*dtos.RegisterDto).Id
	}
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, dtos.NewResponseDto("NG", "Request Format Is Wrong", manga))
		return
	}

	err = controller.repository.Save(&manga)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, dtos.NewResponseDto("NG", "Something Wrong In Save", manga))
		return
	}
	context.JSON(http.StatusCreated, dtos.NewResponseDto("Ok", "Saved Successfully", manga))
}

// @BasePath /manga/
// @Summary Update Manga
// @Schemes Manga
// @Description Update Manga
// @Tags Mangas
// @Accept json
// @Produce json
// @Failure 400 {object} dtos.ResponseDto[string]
// @Router /mangas/{id} [put]
func (controller *MangaController) update(context *gin.Context) {
	id := context.Param("id")

	if id == "" {
		log.Println("wrong request")
		context.JSON(http.StatusBadRequest, dtos.NewResponseDto("NG", "Request Format Is Wrong", id))
		return
	}

	var manga models.Manga

	err := context.ShouldBindBodyWithJSON(&manga)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, dtos.NewResponseDto("NG", "Request Format Is Wrong", manga))
		return
	}
	var userId string
	value, ok := context.Get("user")
	if ok {
		userId = value.(*dtos.RegisterDto).Id
	}

	err = controller.repository.Update(id, userId, &manga)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, dtos.NewResponseDto("NG", "Something Wrong In Save", manga))
		return
	}
	context.JSON(http.StatusCreated, dtos.NewResponseDto("Ok", "Saved Successfully", manga))
}

// @BasePath /manga/
// @Summary Delete Manga
// @Schemes
// @Description Delete Manga
// @Tags Mangas
// @Accept json
// @Produce json
// @Success 200 {object} dtos.ResponseDto[string]
// @Failure 400 {object} dtos.ResponseDto[string]
// @Router /mangas/{id} [delete]
func (controller *MangaController) delete(context *gin.Context) {

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

	err := controller.repository.Delete(id, userId)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, dtos.NewResponseDto("NG", "Something Wrong In Save", id))
		return
	}
	context.JSON(http.StatusCreated, dtos.NewResponseDto("Ok", "Saved Successfully", id))
}
