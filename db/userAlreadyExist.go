package db

import (
	"context"
	"time"

	"github.com/zedeme/webtesting/models"
	"go.mongodb.org/mongo-driver/bson"
)

//UserAlreadyExist is the function to check pre-credentials
func UserAlreadyExist(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("zedb")
	collection := db.Collection("users")

	filter := bson.M{"email": email}

	var result models.User

	err := collection.FindOne(ctx, filter).Decode(&result)
	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
