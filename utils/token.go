package utils

import (
	"errors"
	"realPj/mangaReadingApp/modules/shared/dtos"
	"realPj/mangaReadingApp/modules/shared/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var mySigningKey []byte = []byte("helloworld")

func GenerateToken(user *models.User) (string, error) {

	// Create the Claims
	claims := &jwt.MapClaims{
		"ExpiresAt": jwt.NewNumericDate(time.Now().Add(60 * time.Minute)),
		"Issuer":    "kyawe",
		"name":      user.Name,
		"userId":    user.Id,
		"role":      user.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(mySigningKey)
	return signedToken, err
}

func ValidateToken(tokenString string) (*dtos.RegisterDto, error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("signing method difference")
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return &dtos.RegisterDto{
			Id:   claims["userId"].(string),
			Name: claims["name"].(string),
			Role: claims["role"].(string),
		}, nil
	} else {
		return nil, err
	}
}
