package controller

import (
	// "context"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"myMongoApi/model"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/v2/bson"
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

//Mongo helper

//insert 1 record

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.TODO(), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("successfull inserted", inserted.InsertedID)
}

// update one record
func updateOneMovie(movieId string) {
	id, err := bson.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count: ", result.ModifiedCount)

}

func deleteOneMovie(movieId string) {
	id, err := bson.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted count: ", deleteCount.DeletedCount)
}

func deleteManyMovie() int64 {
	deleteresult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted count: ", deleteresult.DeletedCount)
	return deleteresult.DeletedCount

}

func getAllMovie() []bson.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies []bson.M
	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cur.Close(context.Background())
	return movies
}

//actual controllers

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetAllMovies")
	w.Header().Set("content-type", "application/x-www-form-urlencode")
	allMovies := getAllMovie()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/x-www-form-urlencode")
	w.Header().Set("aAllow-Control-Allow-Methods", "POST")
	var movie model.Netflix
	json.NewDecoder(r.Body).Decode(&movie)

	insertOneMovie(movie)

	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/x-www-form-urlencode")
	w.Header().Set("aAllow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)

	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/x-www-form-urlencode")
	w.Header().Set("aAllow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
func DeleteManyMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/x-www-form-urlencode")
	w.Header().Set("aAllow-Control-Allow-Methods", "DELETE")

	count := deleteManyMovie()
	json.NewEncoder(w).Encode(count)
}
