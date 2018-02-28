package main

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"fmt"
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
)

func main() {
	client, err := mongo.NewClient("mongodb://localhost")
	if err != nil {
		fmt.Println(err)
	}
	db := client.Database("test")

	collection := db.Collection("testColl")

	result, err := collection.InsertOne(
		context.Background(),
		bson.NewDocument(
			bson.EC.String("item", "canvas"),
			bson.EC.Int32("qty", 100),
		))

	fmt.Printf("%#v\n", result)
	fmt.Printf("%#v\n", result.InsertedID)

	count, err := collection.Count(context.Background(), bson.NewDocument(
		bson.EC.Int32("qty", 100),
	))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%T %v\n", count, count)
}
