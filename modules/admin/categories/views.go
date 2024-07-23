package categories

import (
	"log"
	"net/http"
	"realPj/mangaReadingApp/modules/shared/dtos"
	"realPj/mangaReadingApp/modules/shared/models"
	"realPj/mangaReadingApp/modules/shared/repositories"

	"github.com/gin-gonic/gin"
)

func newCategoriesController(repo repositories.ICategoryRepository) CategoriesController {
	return CategoriesController{
		repository: repo,
	}
}

type CategoriesController struct {
	repository repositories.ICategoryRepository
}

func (controller *CategoriesController) index(context *gin.Context) {
	context.JSON(http.StatusOK, dtos.NewResponseDto("OK", "Categories are successfully fetched", controller.repository.GetAll()))
}

func (controller *CategoriesController) save(context *gin.Context) {

	var categories models.Category

	err := context.ShouldBindBodyWithJSON(&categories)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, dtos.NewResponseDto("NG", "Request Format Is Wrong", categories))
		return
	}

	err = controller.repository.Save(&categories)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, dtos.NewResponseDto("NG", "Something Wrong In Save", categories))
		return
	}
	context.JSON(http.StatusCreated, dtos.NewResponseDto("Ok", "Saved Successfully", categories))
}

func (controller *CategoriesController) update(context *gin.Context) {

	var category models.Category

	id := context.Param("id")

	if id == "" {
		context.JSON(http.StatusBadRequest, dtos.NewResponseDto("NG", "Request Format Is Wrong", "string should be empty"))
		return
	}

	err := context.ShouldBindBodyWithJSON(&category)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, dtos.NewResponseDto("NG", "Request Format Is Wrong", category))
		return
	}

	err = controller.repository.Update(id, &category)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, dtos.NewResponseDto("NG", "Something Wrong In Save", category))
		return
	}
	context.JSON(http.StatusCreated, dtos.NewResponseDto("Ok", "Saved Successfully", category))
}

func (controller *CategoriesController) delete(context *gin.Context) {

	id := context.Param("id")

	if id != "" {
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
	context.JSON(http.StatusCreated, dtos.NewResponseDto("Ok", "Saved Successfully", id))
}
