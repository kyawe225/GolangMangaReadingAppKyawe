package normal

import (
	"realPj/mangaReadingApp/modules/normal/chapter"
	"realPj/mangaReadingApp/modules/normal/manga"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine) {
	manga.RegisterRoutes(engine)
	chapter.RegisterRoutes(engine)
}
