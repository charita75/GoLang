package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//student struct
type Student struct {
	Name   string
	Degree string
	City   string
}

//Person struct
type Person struct {
	Name string
	Age  int
	City string
}

func main() {

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	//Inserting a single document to MongoDB
	collection := client.Database("mydb").Collection("students")
	collection1 := client.Database("mydb").Collection("persons")

	charita := Student{"Charita", "MS", "Houston"}
	insertResult, err := collection.InsertOne(context.TODO(), charita)
	fmt.Println("Inserted a Single Document into first collection: ", insertResult.InsertedID)

	ryan := Person{"Ryan", 21, "Santa Clara"}
	james := Person{"James", 32, "Nairobi"}

	trainers := []interface{}{ryan, james}

	insertManyResult, err := collection1.InsertMany(context.TODO(), trainers)
	//insertResult1, err := collection1.InsertOne(context.TODO(), ryan)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple into second collection : ", insertManyResult.InsertedIDs)
	if err != nil {
		log.Fatal(err)
	}

}
