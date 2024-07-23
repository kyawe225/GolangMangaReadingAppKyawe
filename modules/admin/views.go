package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func index(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello World"})
}
