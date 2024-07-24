package chapter

import (
	"realPj/mangaReadingApp/modules/shared/repositories"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine) {
	controller := newChapterController(repositories.ChapterRepository{})
	engine.GET("/api/v1/chapter/:id", controller.detail)
	engine.GET("/api/v1/chapter/", controller.index)
}
