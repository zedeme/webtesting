package db

import (
	"github.com/zedeme/webtesting/models"
	"golang.org/x/crypto/bcrypt"
)

//TryLogin is the function to login all normal users
func TryLogin(email, password string) (models.User, bool) {
	u, founded, _ := UserAlreadyExist(email)
	if founded == false {
		return u, false
	}
	passwordBytes := []byte(password)
	passwordDB := []byte(u.Password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)

	if err != nil {
		return u, false
	}
	return u, true

}
