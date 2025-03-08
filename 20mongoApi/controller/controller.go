package controller

import (
	// "context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionString = "mongodb+srv://sk1181408:sk1181@cluster0.on1re.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

const dbName = "netflix"
const colName = "watchList"

//most IMPORTANT

var collection *mongo.Collection

//connect with mongodb

func init() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(connectionString).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connection success")

	// defer func(){
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 			panic(err)
	// 	}
	// }

	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("collection instance is ready")

}
