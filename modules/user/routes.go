package user

import (
	"realPj/mangaReadingApp/modules/user/chapter"
	"realPj/mangaReadingApp/modules/user/manga"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine) {
	manga.RegisterRoutes(engine)
	chapter.RegisterRoutes(engine)
}
