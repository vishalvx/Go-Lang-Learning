package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vishalvx/dbconnect/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://vishalvx:<password>@try-out.g5ikcxe.mongodb.net/?retryWrites=true&w=majority"
const dbName = "Netflix"
const collectionName = "WatchList"

var collection *mongo.Collection

func init() {

	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOption := options.Client().ApplyURI(connectionString).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("Connected to DB")

	collection = client.Database(dbName).Collection(collectionName)
	log.Println(collection.Name() + " Connected successfully")
}

func insertOneMovie(movie model.Netflix) {

	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("Movie inserted with id of ", inserted.InsertedID)
}

func updateOneMovie(movieId string) {
	id, _err := primitive.ObjectIDFromHex(movieId)
	if _err != nil {
		log.Fatal("Error fail to convert string into object id: ", _err)
		return
	}

	filter := bson.M{"_id": id}
	updates := bson.M{"$set": bson.M{"watched": true}}

	updated, _err := collection.UpdateOne(context.Background(), filter, updates)

	if _err != nil {
		log.Fatal(_err)
		return
	}
	log.Print("update watched to true", updated.ModifiedCount)

}

func deleteOneMovie(movieId string) {
	id, _err := primitive.ObjectIDFromHex(movieId)
	if _err != nil {
		log.Fatal("Error fail to convert string into object id: ", _err)
		return
	}

	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("delete count: ", deleteResult.DeletedCount)

}

func deleteAllMovies() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
		return 0
	}

	log.Println("delete count: ", deleteResult.DeletedCount)

	return deleteResult.DeletedCount
}

func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
		return nil
	}

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		if err := cursor.Decode(&movie); err != nil {
			log.Fatal(err)
			return nil
		}
		movies = append(movies, movie)
	}

	defer cursor.Close(context.Background())
	return movies
}

func GetAllMovies(res http.ResponseWriter, req *http.Request) {
	log.Print("INFO - Get All Movie Endpoint Hit ---")

	res.Header().Set("Content-Type", "application/json")
	allMovies := getAllMovies()

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(allMovies)
}

func CreateOneMovie(res http.ResponseWriter, req *http.Request) {
	log.Print("INFO - Create One Movie Endpoint Hit ---")

	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(req.Body).Decode(&movie)

	insertOneMovie(movie)

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(movie)
}

func MarkedAsWatched(res http.ResponseWriter, req *http.Request) {
	log.Print("INFO - Update movie to marked as watched Endpoint Hit ---")

	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(req)
	updateOneMovie(params["id"])

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode("update movie status to marked as read")
}

func DeleteOneMovie(res http.ResponseWriter, req *http.Request) {
	log.Print("INFO -  Delete One Movie Endpoint Hit ---")

	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(req)
	deleteOneMovie(params["id"])

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode("Delete One movie.")
}

func DeleteAllMovies(res http.ResponseWriter, req *http.Request) {
	log.Print("INFO -  Delete All Movies Endpoint Hit ---")

	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	deleteAllMovies()

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode("Delete All movies.")
}

func GetHome(res http.ResponseWriter, req *http.Request) {
	log.Print("INFO -  Delete All Movies Endpoint Hit ---")

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode("Hello world!!!")
}
