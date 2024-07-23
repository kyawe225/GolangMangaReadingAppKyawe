package auth

import (
	"log"
	"net/http"
	"realPj/mangaReadingApp/modules/shared/dtos"
	"realPj/mangaReadingApp/modules/shared/repositories"

	"github.com/gin-gonic/gin"
)

func NewAuthController(repository repositories.IAuthRepository) AuthController {
	return AuthController{
		repo: repository,
	}
}

type AuthController struct {
	repo repositories.IAuthRepository
}

// @BasePath /auth/

// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /auth/login [post]
func (controller *AuthController) login(context *gin.Context) {
	var model dtos.LoginDto
	err := context.ShouldBindBodyWithJSON(&model)
	if err != nil {
		context.JSON(http.StatusBadRequest, dtos.NewResponseDto("NG", "Request Format Not Right", model))
		return
	}
	token, err := controller.repo.Login(&model)
	if err != nil {
		context.JSON(http.StatusBadRequest, dtos.NewResponseDto("NG", "Incorrect Email Or Password", model))
		return
	}
	context.JSON(http.StatusOK, dtos.NewResponseDto("OK", "Login Successfully", map[string]string{"Login_Token": (*token)}))
}

// @BasePath /auth/

// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /auth/register [post]
func (controller *AuthController) register(context *gin.Context) {
	var model dtos.RegisterDto
	err := context.ShouldBindBodyWithJSON(&model)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, dtos.NewResponseDto("NG", "Request Format Not Right", model))
		return
	}
	err = controller.repo.Register(&model)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, dtos.NewResponseDto("NG", "Please Try Again Later", model))
		return
	}
	context.JSON(http.StatusOK, dtos.NewResponseDto("OK", "Login Successfully", model))
}
