package authuser

import (
	"realPj/mangaReadingApp/modules/authuser/profile"
	"realPj/mangaReadingApp/modules/shared/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine) {
	group := engine.Group("/api/v1/user")

	group.Use(middleware.CheckLoginUser)

	profile.RegisterRoutes(group)

}
