package bookmark

import (
	"realPj/mangaReadingApp/modules/shared/repositories"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(group *gin.RouterGroup) {
	controller := newBookMarkController(repositories.BookMarkRespotiroy{})

	group.GET("/bookmark", controller.index)
	group.POST("/bookmark/save/:id", controller.save)
	group.DELETE("/bookmark/delete/:id", controller.delete)
}
