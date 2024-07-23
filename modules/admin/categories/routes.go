package categories

import (
	"realPj/mangaReadingApp/modules/shared/repositories"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.RouterGroup) {
	categoryController := newCategoriesController(repositories.CategoryRepository{})
	group := engine.Group("/categories")
	group.GET("/", categoryController.index)
	group.POST("/create", categoryController.save)
	group.PUT("/update/:id", categoryController.update)
	group.DELETE("/delete/:id", categoryController.delete)
	group.GET("/:id", categoryController.index)
}
