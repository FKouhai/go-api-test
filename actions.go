package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	c "main/config"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	//	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"
)

type Message struct {
	Status  string `json: "status"`
	Message string `json: "message"`
}

func (this *Message) setStatus(data string) {
	this.Status = data
}
func (this *Message) setMessage(data string) {
	this.Message = data
}

func getSession(config *c.Config) *mongo.Client {
	ctx := context.TODO()
	config = c.ReadConfig()
	session, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoHost))
	//	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	return session
}

func responseMovie(w http.ResponseWriter, status int, results Movie) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)
}
func responseMovies(w http.ResponseWriter, status int, results []Movie) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)
}

var collection = getSession(c.ReadConfig()).Database(c.ReadConfig().MongoDB).Collection(c.ReadConfig().MongoDB)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo")
}

func MovieList(w http.ResponseWriter, r *http.Request) {
	var results []Movie
	ctx := context.TODO()
	fmt.Fprint(w, "Listado de peliculas")
	res, err := collection.Find(ctx, bson.D{{}}, options.Find())
	if err != nil {
		panic(err)
	} else {
		for res.Next(context.TODO()) {
			var movie Movie
			err := res.Decode(&movie)
			if err != nil {
				log.Fatalln(err)
			}
			results = append(results, movie)
		}
		fmt.Println("Resultados: ", results)
	}
	responseMovies(w, 200, results)
}

func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]
	fmt.Println(movie_id)
	oid, _ := primitive.ObjectIDFromHex(movie_id)
	fmt.Println(oid)
	results := Movie{}
	err := collection.FindOne(context.Background(), bson.M{"_id": oid}).Decode(&results)
	if err != nil {
		panic(err)
	}
	responseMovie(w, 200, results)
}

func MovieAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var movie_data Movie
	err := decoder.Decode(&movie_data)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	ctx := context.TODO()
	if _, err := collection.InsertOne(ctx, movie_data); err != nil {
		w.WriteHeader(500)
	}
	responseMovie(w, 200, movie_data)

}

func MovieRemove(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]
	decoder := json.NewDecoder(r.Body)
	oid, _ := primitive.ObjectIDFromHex(movie_id)
	var movie_data Movie

	err := decoder.Decode(&movie_data)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	document := bson.M{"_id": oid}

	_, err1 := collection.DeleteOne(context.Background(), document)
	if err1 != nil {
		panic(err1)
	}

	//results := Message{"Success", "The movie with ID" + movie_id + "Ha sido borrada correctamente"}
	message := new(Message)
	message.setStatus("success")
	message.setMessage("The movie with ID" + movie_id + "has been removed")
	results := message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}

func MovieUpdate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]
	decoder := json.NewDecoder(r.Body)
	oid, _ := primitive.ObjectIDFromHex(movie_id)
	var movie_data Movie

	err := decoder.Decode(&movie_data)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	document := bson.M{"_id": oid}

	change := bson.M{"$set": movie_data}

	_, err1 := collection.UpdateOne(context.Background(), document, change)
	if err1 != nil {
		panic(err1)
	}
	responseMovie(w, 200, movie_data)
}
