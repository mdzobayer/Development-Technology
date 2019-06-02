// Source: https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial
// Source: https://github.com/tfogo/mongodb-go-tutorial
// Official: https://godoc.org/go.mongodb.org/mongo-driver/mongo
// Large Example: https://godoc.org/go.mongodb.org/mongo-driver/mongo

// Indexing Source: https://docs.mongodb.com/manual/indexes/

package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	// Create collections to database
	collection := client.Database("test").Collection("trainers")

	//---------------Insert----------------

	// Prepareing a Document
	ash := Trainer{"Ash", 10, "Pallet Town"}
	ash2 := Trainer{"Ash2", 20, "Pallet Town"}
	misty := Trainer{"Misty", 15, "Cerulean City"}
	brock := Trainer{"Brock", 20, "Pewter City"}

	// Inserting a single Document
	insertResult, err := collection.InsertOne(context.TODO(), ash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	// Inserting multiple document
	trainers := []interface{}{ash2, misty, brock}

	insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

	//-----------------Update---------------------

	filter := bson.D{{"name", "Ash"}}

	// update := bson.D{
	// 	{"$set", bson.D{
	// 		{"age", 50},
	// 	}},
	// }
	// fmt.Println("filter and update success")
	//updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("Matched %v documents and update %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	//------------------Find Documents--------------

	var result Trainer

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n", result)

	findOptions := options.Find()
	findOptions.SetLimit(0)

	// Filter with logical and (&&) operation
	// filter = bson.D{
	// 	{"$and", []interface{}{
	// 		bson.D{{"name", "Ash"}},
	// 		bson.D{{"age", 20}}}}}

	// Filter with logical or (||) operation
	// filter = bson.D{
	// 	{"$or", []interface{}{
	// 		bson.D{{"name", "Ash"}},
	// 		bson.D{{"age", 20}}}}}

	// Filter with less than (<) operator
	// filter = bson.D{
	// 	{"age", bson.D{{"$lt", 20}}}}

	// Filter with less than equal (<=) operator
	// filter = bson.D{
	// 	{"age", bson.D{{"$lte", 15}}}}

	// FIlter with grater than (>) operator
	filter = bson.D{
		{"age", bson.D{{"$gt", 10}}}}

	// Filter with grater than equal(>=) operator
	// filter = bson.D{
	// 	{"age", bson.D{{"$gte", 15}}}}

	var results []*Trainer

	cur, err := collection.Find(context.TODO(), filter, findOptions)

	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem Trainer
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	for x, res := range results {
		fmt.Printf("Found doument %+v item %+v\n", x, res)
	}

	cur.Close(context.TODO())
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)

	//--------------Delete Documents---------------
	deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{{"name", "Misty"}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

	//--------------Close Connections---------------

	// Collections Closed
	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")

}
