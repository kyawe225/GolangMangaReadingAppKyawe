package auth

import (
	"realPj/mangaReadingApp/modules/shared/repositories"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine) {
	controller := NewAuthController(repositories.AuthRepository{})
	engine.POST("/auth/login", controller.login)
	engine.POST("/auth/register", controller.register)
}
