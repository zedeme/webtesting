package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User is the normal struct for Users
type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username    string             `bson:"username" json:"username,omitempty"`
	Password    string             `bson:"password" json:"password,omitempty"`
	FirtsName   string             `bson:"firtsname" json:"firtsname,omitempty"`
	LastName    string             `bson:"lastname" json:"lastname,omitempty"`
	Email       string             `bson:"email" json:"email,omitempty"`
	Birthday    time.Time          `bson:"birthday" json:"birthday,omitempty"`
	Avatar      string             `bson:"avatar" json:"avatar,omitempty"`
	Banner      string             `bson:"banner" json:"banner,omitempty"`
	Description string             `bson:"description" json:"description,omitempty"`
	Ubication   string             `bson:"ubication" json:"ubication,omitempty"`
	WebSite     string             `bson:"website" json:"website,omitempty"`
}
