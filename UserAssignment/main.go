package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

//UserInfo Struct
type UserInfo struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

//UserInterests Struct
type UserInterests struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Sport string             `json:"sport,omitempty" bson:"sport,omitempty"`
	Hobby string             `json:"hobby,omitempty" bson:"hobby,omitempty"`
}

//Creating User information
func CreateUserInfoEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var userinfo UserInfo
	_ = json.NewDecoder(request.Body).Decode(&userinfo)
	collection := client.Database("user").Collection("info")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, userinfo)
	json.NewEncoder(response).Encode(result)
}

//RCreating User Interests
func CreateUserInterestsEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var userinterests UserInterests
	_ = json.NewDecoder(request.Body).Decode(&userinterests)
	collection := client.Database("user").Collection("interests")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, userinterests)
	json.NewEncoder(response).Encode(result)
}

//Retrieving UserInfo
func GetUserInfoEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var info []UserInfo
	collection := client.Database("user").Collection("info")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var userinfo UserInfo
		cursor.Decode(&userinfo)
		info = append(info, userinfo)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(info)
}

//Retrieving Userinterests
func GetUserInterestsEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var interests []UserInterests
	collection := client.Database("user").Collection("interests")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var userinterests UserInterests
		cursor.Decode(&userinterests)
		interests = append(interests, userinterests)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(interests)
}

//Get Userinfo By ID
func GetUserInfoByIdEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var userinfo UserInfo
	collection := client.Database("user").Collection("info")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, UserInfo{ID: id}).Decode(&userinfo)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(userinfo)
}

func main() {
	fmt.Println("Starting the application..")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	router := mux.NewRouter()
	router.HandleFunc("/userinfo", CreateUserInfoEndpoint).Methods("POST")
	router.HandleFunc("/info", GetUserInfoEndpoint).Methods("GET")
	router.HandleFunc("/userinterests", CreateUserInterestsEndpoint).Methods("POST")
	router.HandleFunc("/interests", GetUserInterestsEndpoint).Methods("GET")
	router.HandleFunc("/userinfo/{id}", GetUserInfoByIdEndpoint).Methods("GET")
	http.ListenAndServe(":8091", router)
}
