package middleware

import (
	"log"
	"net/http"
	consts "realPj/mangaReadingApp/const"
	"realPj/mangaReadingApp/modules/shared/dtos"
	"realPj/mangaReadingApp/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckLoginAdmin(context *gin.Context) {
	token := context.GetHeader(consts.AuthorizationHeaderKey)

	if utils.IsEmpty(&token) {
		log.Println("user does not logged in")
		context.AbortWithStatusJSON(http.StatusUnauthorized, dtos.NewResponseDto("NG", "Not Logged Token", "Not Logged In"))
		return
	}

	token = strings.Split(token, consts.AuthorizationTypeBearer)[1]
	token = strings.Trim(token, " ")
	dto, err := utils.ValidateToken(token)

	if err != nil {
		log.Println(err)
		context.AbortWithStatusJSON(http.StatusUnauthorized, dtos.NewResponseDto("NG", "Please Login Again", "Please Login Again"))
		return
	}

	if dto.Role != "admin" {
		context.AbortWithStatusJSON(http.StatusForbidden, dtos.NewResponseDto("NG", "Forbidden Access", "Please Login With Another Account"))
		return
	}

	context.Set("user", dto)

	context.Next()
}

func CheckLoginUser(context *gin.Context) {
	token := context.GetHeader(consts.AuthorizationHeaderKey)

	if utils.IsEmpty(&token) {
		log.Println("user does not logged in")
		context.AbortWithStatusJSON(http.StatusUnauthorized, dtos.NewResponseDto("NG", "Not Logged Token", "Not Logged In"))
		return
	}

	token = strings.Split(token, consts.AuthorizationTypeBearer)[1]
	token = strings.Trim(token, " ")
	dto, err := utils.ValidateToken(token)

	if err != nil {
		log.Println(err)
		context.AbortWithStatusJSON(http.StatusUnauthorized, dtos.NewResponseDto("NG", "Please Login Again", "Please Login Again"))
		return
	}

	context.Set("user", dto)

	context.Next()
}
