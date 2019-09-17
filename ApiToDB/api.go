// package api

// import (
// 	"context"
// 	"encoding/json"
// 	"net/http"
// 	"time"

// 	"github.com/gorilla/mux"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// //Exported
// func CreatePersonEndpoint(response http.ResponseWriter, request *http.Request) {
// 	response.Header().Set("content-type", "application/json")
// 	var person Person
// 	_ = json.NewDecoder(request.Body).Decode(&person)
// 	collection := client.Database("thepolyglotdeveloper").Collection("people")
// 	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
// 	result, _ := collection.InsertOne(ctx, person)
// 	json.NewEncoder(response).Encode(result)
// }

// //Exported
// func GetPersonEndpoint(response http.ResponseWriter, request *http.Request) {
// 	response.Header().Set("content-type", "application/json")
// 	params := mux.Vars(request)
// 	id, _ := primitive.ObjectIDFromHex(params["id"])
// 	var person Person
// 	collection := client.Database("thepolyglotdeveloper").Collection("people")
// 	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
// 	err := collection.FindOne(ctx, Person{ID: id}).Decode(&person)
// 	if err != nil {
// 		response.WriteHeader(http.StatusInternalServerError)
// 		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
// 		return
// 	}
// 	json.NewEncoder(response).Encode(person)
// }

// //Exported
// func GetPeopleEndpoint(response http.ResponseWriter, request *http.Request) {
// 	response.Header().Set("content-type", "application/json")
// 	var people []Person
// 	collection := client.Database("thepolyglotdeveloper").Collection("people")
// 	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
// 	cursor, err := collection.Find(ctx, bson.M{})
// 	if err != nil {
// 		response.WriteHeader(http.StatusInternalServerError)
// 		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
// 		return
// 	}
// 	defer cursor.Close(ctx)
// 	for cursor.Next(ctx) {
// 		var person Person
// 		cursor.Decode(&person)
// 		people = append(people, person)
// 	}
// 	if err := cursor.Err(); err != nil {
// 		response.WriteHeader(http.StatusInternalServerError)
// 		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
// 		return
// 	}
// 	json.NewEncoder(response).Encode(people)
// }
