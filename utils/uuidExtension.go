package utils

import (
	"log"

	"github.com/google/uuid"
)

func GenerateUUIDV7() string {
	uuidS, err := uuid.NewV7()
	if err != nil {
		log.Println(err)
	}
	return uuidS.String()
}
