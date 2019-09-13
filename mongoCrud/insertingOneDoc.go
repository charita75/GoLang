// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// //Person struct
// type student struct {
// 	Name   string
// 	Degree string
// 	City   string
// }

// func main() {

// 	// Set client options
// 	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

// 	// Connect to MongoDB
// 	client, err := mongo.Connect(context.TODO(), clientOptions)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Check the connection
// 	err = client.Ping(context.TODO(), nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Connected to MongoDB!")

// 	//Inserting a single document to MongoDB
// 	collection := client.Database("mydb").Collection("students")

// 	charita := student{"Charita", "MS", "Houston"}

// 	insertResult, err := collection.InsertOne(context.TODO(), charita)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)

// }
