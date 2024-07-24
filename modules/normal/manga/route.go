package manga

import (
	"net/http"
	"realPj/mangaReadingApp/modules/shared/dtos"
	"realPj/mangaReadingApp/modules/shared/repositories"

	"github.com/gin-gonic/gin"
)

func newUserMangaController(repo repositories.IMangaRepository) UserMangaController {
	return UserMangaController{
		repository: repo,
	}
}

type UserMangaController struct {
	repository repositories.IMangaRepository
}

func (controller *UserMangaController) index(context *gin.Context) {
	context.JSON(http.StatusOK, dtos.NewResponseDto("OK", "Successfully Fetched Mangas", controller.repository.GetList()))
}

func (controller *UserMangaController) detail(context *gin.Context) {
	id := context.Param("id")

	if id == "" {
		context.JSON(http.StatusOK, dtos.NewResponseDto("NG", "Request Format Is Wrong", "Wrong Id"))
		return
	}
	context.JSON(http.StatusOK, dtos.NewResponseDto("OK", "Successfully Fetched Mangas", controller.repository.FindByIdDto(id)))
}
