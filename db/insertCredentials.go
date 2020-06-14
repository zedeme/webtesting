package db

import (
	"context"
	"time"

	"github.com/zedeme/webtesting/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//InsertCredentials is the function for send credentials to DB
func InsertCredentials(newUser models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("zedb")
	collection := db.Collection("users")

	newUser.Password, _ = PasswordEncrypt(newUser.Password)

	newUser.ID = primitive.NewObjectID()

	_, err := collection.InsertOne(ctx, newUser)

	if err != nil {
		return "", false, err
	}

	return newUser.ID.String(), true, nil
}
