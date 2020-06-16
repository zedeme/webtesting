package db

import (
	"context"
	"fmt"
	"time"

	"github.com/zedeme/webtesting/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SearchProfile is the function to search user profile
func SearchProfile(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("zedb")
	collection := db.Collection("users")

	var profile models.User

	objID, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.M{
		"_id": objID,
	}

	err := collection.FindOne(ctx, filter).Decode(profile)

	profile.Password = ""

	if err != nil {
		fmt.Println("Error profile dont found ", err.Error())
		return profile, err
	}
	return profile, nil
}
