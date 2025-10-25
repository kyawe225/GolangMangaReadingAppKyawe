package bookmark

import (
	"log"
	"net/http"
	"realPj/mangaReadingApp/modules/shared/dtos"
	"realPj/mangaReadingApp/modules/shared/repositories"
	"realPj/mangaReadingApp/utils"

	"github.com/gin-gonic/gin"
)

func newBookMarkController(repository repositories.IBookMarkRepository) bookmarkController {
	return bookmarkController{
		repo: repository,
	}
}

type bookmarkController struct {
	repo repositories.IBookMarkRepository
}

func (controller bookmarkController) index(context *gin.Context) {
	id := utils.ExtractLoggedUserId(context)
	chapter, err := controller.repo.ListAll(id)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, dtos.NewResponseDto("NG", "Request Format Is Wrong", err))
		return
	}
	context.JSON(http.StatusOK, dtos.NewResponseDto("OK", "Successfully Fetched Mangas", chapter))
}

func (controller bookmarkController) save(context *gin.Context) {
	id := utils.ExtractLoggedUserId(context)
	manga_publish_id := context.Param("id")

	if manga_publish_id == "" {
		context.JSON(http.StatusNotFound, dtos.NewResponseDto("NG", "Request Format Is Wrong", "Invalid Id"))
		return
	}

	err := controller.repo.BookMarkManga(manga_publish_id, id)

	if err != nil {
		context.JSON(http.StatusNotFound, dtos.NewResponseDto("NG", "Request Format Is Wrong", err))
		return
	}
	context.JSON(http.StatusOK, dtos.NewResponseDto("OK", "Successfully Fetched Mangas", "Successfully BookMarked"))
}

func (controller bookmarkController) delete(context *gin.Context) {
	id := utils.ExtractLoggedUserId(context)
	manga_publish_id := context.Param("id")

	if manga_publish_id == "" {
		context.JSON(http.StatusNotFound, dtos.NewResponseDto("NG", "Request Format Is Wrong", "Invalid Id"))
		return
	}

	err := controller.repo.Delete(manga_publish_id, id)

	if err != nil {
		context.JSON(http.StatusNotFound, dtos.NewResponseDto("NG", "Request Format Is Wrong", err))
		return
	}
	context.JSON(http.StatusOK, dtos.NewResponseDto("OK", "Successfully Fetched Mangas", "Successfully Remove BookMark"))
}
