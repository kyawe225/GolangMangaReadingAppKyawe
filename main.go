package main

import (
	"realPj/mangaReadingApp/docs"
	"realPj/mangaReadingApp/modules/admin"
	"realPj/mangaReadingApp/modules/auth"
	"realPj/mangaReadingApp/modules/authuser"
	"realPj/mangaReadingApp/modules/normal"
	"realPj/mangaReadingApp/utils"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	utils.InitDB()

	docs.SwaggerInfo.BasePath = "/"

	server := gin.Default()

	admin.RegisterRoutes(server)
	auth.RegisterRoutes(server)
	normal.RegisterRoutes(server)
	authuser.RegisterRoutes(server)

	defer utils.DB.Close()
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	server.Run(":8000")
}
