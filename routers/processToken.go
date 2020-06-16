package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/zedeme/webtesting/db"
	"github.com/zedeme/webtesting/models"
)

//Email is the email of user
var Email string

//IDUser is the ID of user
var IDUser string

//ProcessToken is the function to process validate of JWT
func ProcessToken(token string) (*models.Claim, bool, string, error) {
	jwtPassword := []byte("Zed1547299445213244900060")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("invalid token")
	}

	token = strings.TrimSpace(splitToken[1])

	jwttoken, err := jwt.ParseWithClaims(token, claims, func(tkn *jwt.Token) (interface{}, error) {
		return jwtPassword, nil
	})

	if err == nil {
		_, founded, _ := db.UserAlreadyExist(claims.Email)
		if founded {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return claims, founded, IDUser, nil
	}
	if !jwttoken.Valid {
		return claims, false, string(""), errors.New("invalid token")
	}
	return claims, false, string(""), err
}
