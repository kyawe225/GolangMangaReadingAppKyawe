package mangas

import (
	"realPj/mangaReadingApp/modules/shared/repositories"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.RouterGroup) {
	var mangaRepository repositories.IMangaRepository = repositories.MangaRepository{}
	mangaController := newMangaController(mangaRepository)

	group := engine.Group("/manga")

	group.GET("", mangaController.index)
	group.POST("/create", mangaController.save)
	group.DELETE("/delete/:id", mangaController.delete)
	group.PUT("/update/:id", mangaController.update)
	group.GET("/:id", mangaController.detail)
}
