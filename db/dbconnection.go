package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCN is the var to connect Mongodb
var MongoCN = ConnectDB()

var clientOptions = options.Client().ApplyURI("mongodb+srv://sebastian:Zed15472994@zedb-6kra0.mongodb.net/<dbname>?retryWrites=true&w=majority")

//ConnectDB is the connect DB function
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("All its fine, we are connected with DB...")

	return client
}

//ConnectionCheck is the connection check function
func ConnectionCheck() bool {

	err := MongoCN.Ping(context.TODO(), nil)

	if err != nil {
		return false
	}
	return true
}
