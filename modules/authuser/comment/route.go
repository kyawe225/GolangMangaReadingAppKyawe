package comment

import (
	"realPj/mangaReadingApp/modules/shared/repositories"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(group *gin.RouterGroup) {
	controller := newCommentController(repositories.CommentRepository{})
	group.Group("/comment")
	group.GET("/:id", controller.index)
	group.GET("/:id", controller.save)
}
