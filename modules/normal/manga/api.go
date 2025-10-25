package manga

import (
	"realPj/mangaReadingApp/modules/shared/repositories"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine) {
	controller := newUserMangaController(repositories.MangaRepository{})
	engine.GET("/api/v1/manga/:id", controller.detail)
	engine.GET("/api/v1/manga/", controller.index)
}
