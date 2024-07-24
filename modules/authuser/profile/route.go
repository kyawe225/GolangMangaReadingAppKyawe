package profile

import (
	"realPj/mangaReadingApp/modules/shared/repositories"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(group *gin.RouterGroup) {
	controller := newUserController(repositories.AuthRepository{})

	group.GET("/profile", controller.profile)
}
