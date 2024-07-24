package utils

import (
	"realPj/mangaReadingApp/modules/shared/dtos"

	"github.com/gin-gonic/gin"
)

func IsEmpty(s *string) bool {
	return (*s) == ""
}

func ExtractLoggedUserId(context *gin.Context) string {
	var userId string
	value, ok := context.Get("user")
	if ok {
		userId = value.(dtos.RegisterDto).Id
	}
	return userId
}
