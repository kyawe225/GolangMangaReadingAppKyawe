package user

import (
	"realPj/mangaReadingApp/modules/shared/repositories"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(group *gin.RouterGroup) {
	controller := newUserController(repositories.UserRepository{})
	group.GET("", controller.index)
	group.POST("/create", controller.save)
	group.DELETE("/delete/:id", controller.delete)
	group.PUT("/update/:id", controller.update)
	group.GET("/:id", controller.detail)
}
