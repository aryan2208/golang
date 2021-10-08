package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

type Users struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	name string             `json:"name,omitempty" bson:"name,omitempty"`
	email  string             `json:"email,omitempty" bson:"email,omitempty"`
	password string				`json:"email,omitempty" bson:"email,omitempty"`
}

type Posts struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	caption string             `json:"caption,omitempty" bson:"caption,omitempty"`
	url  string             `json:"url,omitempty" bson:"url,omitempty"`
	ts string				`json:"url,omitempty" bson:"url,omitempty"`
}

func CreateusersEndpoint(response http.ResponseWriter, request *http.Request) {}
func GetusersEndpoint(response http.ResponseWriter, request *http.Request) { }
func GetusersEndpoint(response http.ResponseWriter, request *http.Request) { }


func CreateusersEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var users 
	_ = json.NewDecoder(request.Body).Decode(&users)
	collection := client.Database("thepolyglotdeveloper").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, users)
	json.NewEncoder(response).Encode(result)
}

func CreatePostsEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var posts 
	_ = json.NewDecoder(request.Body).Decode(&users)
	collection := client.Database("thepolyglotdeveloper").Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, posts)
	json.NewEncoder(response).Encode(result)
}

func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	router := mux.NewRouter()
	router.HandleFunc("/users", CreateusersEndpoint).Methods("POST")
	router.HandleFunc("/{id}", GetusersEndpoint).Methods("GET")
	router.HandleFunc("/users/{id}", GetusersEndpoint).Methods("GET")
	router.HandleFunc("/posts/{id}", GetusersEndpoint).Methods("GET")
	router.HandleFunc("/users/posts/{id}", GetusersEndpoint).Methods("GET")
	http.ListenAndServe(":12345", router)
}

