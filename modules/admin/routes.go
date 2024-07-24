package admin

import (
	"realPj/mangaReadingApp/modules/admin/categories"
	"realPj/mangaReadingApp/modules/admin/chapters"
	"realPj/mangaReadingApp/modules/admin/mangas"
	"realPj/mangaReadingApp/modules/admin/user"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine) {
	engine.GET("/admin", index)
	group := engine.Group("/admin")
	// group.Use(middleware.CheckLoginUser)
	mangas.RegisterRoutes(group)
	categories.RegisterRoutes(group)
	chapters.RegisterRoutes(group)
	user.RegisterRoutes(group)
}
