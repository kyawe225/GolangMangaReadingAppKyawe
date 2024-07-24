package profile

import (
	"log"
	"net/http"
	"realPj/mangaReadingApp/modules/shared/dtos"
	"realPj/mangaReadingApp/modules/shared/repositories"
	"realPj/mangaReadingApp/utils"

	"github.com/gin-gonic/gin"
)

func newUserController(repository repositories.IAuthRepository) userController {
	return userController{
		repo: repository,
	}
}

type userController struct {
	repo repositories.IAuthRepository
}

func (controller userController) profile(context *gin.Context) {
	id := utils.ExtractLoggedUserId(context)
	chapter, err := controller.repo.Profile(id)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, dtos.NewResponseDto("NG", "Request Format Is Wrong", err))
		return
	}
	context.JSON(http.StatusOK, dtos.NewResponseDto("OK", "Successfully Fetched Mangas", chapter))
}
