package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/zedeme/webtesting/models"
)

//GeneratorJWT is the Json Web Token generator
func GeneratorJWT(u models.User) (string, error) {
	JWTPassword := []byte("Zed1547299445213244900060")

	payload := jwt.MapClaims{
		"email":       u.Email,
		"username":    u.Username,
		"firtsname":   u.FirtsName,
		"lastname":    u.LastName,
		"birthday":    u.Birthday,
		"description": u.Description,
		"website":     u.WebSite,
		"_id":         u.ID.Hex(),
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(JWTPassword)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil

}
