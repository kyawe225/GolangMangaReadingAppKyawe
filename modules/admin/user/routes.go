package user

import (
	"realPj/mangaReadingApp/modules/shared/repositories"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(group *gin.RouterGroup) {
	controller := newUserController(repositories.UserRepository{})
	userGroup := group.Group("/user")
	userGroup.GET("/", controller.index)
	userGroup.POST("/create", controller.save)
	userGroup.DELETE("/delete/:id", controller.delete)
	userGroup.PUT("/update/:id", controller.update)
	userGroup.GET("/:id", controller.detail)
}
