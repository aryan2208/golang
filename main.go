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

func GetPostsEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var users []Posts
	collection := client.Database("aryan2208").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var posts
		cursor.Decode(&person)
		posts = append(users, posts)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(users)
}
func GetUsersEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var users []Users
	collection := client.Database("aryan2208").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var users
		cursor.Decode(&users)
		users = append(users, users)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(users)
}

func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{}
    	var limit int64 = 10
    	var page int64 = 1
    	collection := client.Database("myaggregate").Collection("stocks")
    	projection := bson.D{
    		{"id", 10001},
    		
    	}

        var products []Product
    	paginatedData_posts, err := New(collection).Context(ctx).Limit(limit).Page(page).Sort("price", -1).Select(projection).Filter(filter).Decode(&posts).Find()
	paginatedData_users, err := New(collection).Context(ctx).Limit(limit).Page(page).Sort("price", -1).Select(projection).Filter(filter).Decode(&users).Find()
    	if err != nil {
    		panic(err)
    	}
    
    	
    	fmt.Printf("Normal Find Data: %+v\n", posts)
    
    	
    	fmt.Printf("Normal find pagination info: %+v\n", paginatedData.Pagination)
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

