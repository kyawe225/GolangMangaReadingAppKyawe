package user

import (
	"log"
	"net/http"
	"realPj/mangaReadingApp/modules/shared/dtos"
	"realPj/mangaReadingApp/modules/shared/models"
	"realPj/mangaReadingApp/modules/shared/repositories"

	"github.com/gin-gonic/gin"
)

func newUserController(repository repositories.IUserRepository) UserController {
	return UserController{
		repo: repository,
	}
}

type UserController struct {
	repo repositories.IUserRepository
}

// @BasePath /user/
// @Summary List All User
// @Schemes
// @Description List All User
// @Tags Users
// @Accept json
// @Produce json
// @Failure 400 {object} dtos.ResponseDto[string]
// @Router /user [get]
func (user UserController) index(context *gin.Context) {
	list, err := user.repo.List()
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, dtos.NewResponseDto("OK", "Successfully Fetched the user data", list))
	}
	context.JSON(http.StatusOK, dtos.NewResponseDto("OK", "Successfully Fetched the user data", list))
}

// @BasePath /user/
// @Summary Save User
// @Schemes
// @Description Save User
// @Tags Users
// @Accept json
// @Produce json
// @Failure 400 {object} dtos.ResponseDto[string]
// @Router /user [post]
func (user UserController) save(context *gin.Context) {
	var model models.User
	err := context.ShouldBindBodyWithJSON(&model)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusOK, dtos.NewResponseDto("NG", "Request Format Wrong", "Please fill required fields"))
		return
	}

	err = user.repo.Save(&model)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusOK, dtos.NewResponseDto("NG", "Request Format Wrong", "Saved Failed Please resend and check other fields"))
		return
	}

	context.JSON(http.StatusCreated, dtos.NewResponseDto("NG", "Request Format Wrong", "Saved Successfully"))
}

// @BasePath /user/
// @Summary Update User
// @Schemes
// @Description Update User
// @Tags Users
// @Accept json
// @Produce json
// @Failure 400 {object} dtos.ResponseDto[string]
// @Router /user/{id} [put]
func (user UserController) update(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		context.AbortWithStatusJSON(http.StatusOK, dtos.NewResponseDto("NG", "Request Format Wrong", "Invalid Id"))
		return
	}
	var model models.User
	err := context.ShouldBindBodyWithJSON(&model)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusOK, dtos.NewResponseDto("NG", "Request Format Wrong", "Please fill required fields"))
		return
	}

	err = user.repo.Update(id, &model)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusOK, dtos.NewResponseDto("NG", "Request Format Wrong", err))
		return
	}

	context.JSON(http.StatusOK, dtos.NewResponseDto("OK", "Saved Successfully", "Saved Successfully"))
}

// @BasePath /user/
// @Summary Delete User
// @Schemes
// @Description Delete User
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} dtos.ResponseDto[string]
// @Failure 400 {object} dtos.ResponseDto[string]
// @Router /user/{id} [delete]
func (user UserController) delete(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		context.AbortWithStatusJSON(http.StatusOK, dtos.NewResponseDto("NG", "Request Format Wrong", "Invalid Id"))
		return
	}
	err := user.repo.Delete(id)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusOK, dtos.NewResponseDto("NG", "Request Format Wrong", err))
		return
	}

	context.JSON(http.StatusOK, dtos.NewResponseDto("OK", "Delete Successfully", "Delete Successfully"))
}

// @BasePath /user/
// @Summary Detail User
// @Schemes
// @Description Detail User
// @Tags Users
// @Accept json
// @Produce json
// @Failure 400 {object} dtos.ResponseDto[string]
// @Router /user/{id} [get]
func (controller *UserController) detail(context *gin.Context) {
	id := context.Param("id")

	if id == "" {
		context.JSON(http.StatusBadRequest, dtos.NewResponseDto("NG", "Request Format Is Wrong", "Wrong Id"))
		return
	}
	chapter, err := controller.repo.FindById(id)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, dtos.NewResponseDto("NG", "Request Format Is Wrong", err))
		return
	}
	context.JSON(http.StatusOK, dtos.NewResponseDto("OK", "Successfully Fetched Mangas", chapter))
}
