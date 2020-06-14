package db

import (
	"golang.org/x/crypto/bcrypt"
)

//PasswordEncrypt is the function to encrypt the password for DB
func PasswordEncrypt(password string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}
