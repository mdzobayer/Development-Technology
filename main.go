package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CFProblem struct {
	ContestID      int      `json:"contestId,omitempty"`
	ProblemsetName string   `json:"problemsetName,omitempty"`
	Index          string   `json:"index,omitempty"`
	Name           string   `json:"name,omitempty"`
	Type           string   `json:"type,omitempty"`
	Points         float32  `json:"points,omitempty"`
	Rating         int      `json:"rating,omitempty"`
	Tags           []string `json:"tags,omitempty"`
}

type ProcessedCodeforcesProblem struct {
	ProblemID      string   `bson:"problemId,omitempty"`
	ContestID      int      `bson:"contestId,omitempty"`
	ProblemsetName string   `bson:"problemsetName,omitempty"`
	Index          string   `bson:"index,omitempty"`
	Name           string   `bson:"name,omitempty"`
	Type           string   `bson:"type,omitempty"`
	Points         float32  `bson:"points,omitempty"`
	Rating         int      `bson:"rating,omitempty"`
	Tags           []string `bson:"tags,omitempty"`
}

func ConvertCodeforcesProblem(jp CFProblem, bp *ProcessedCodeforcesProblem) {
	bp.ContestID = jp.ContestID
	bp.Index = jp.Index
	bp.Name = jp.Name
	bp.Points = jp.Points
	bp.ProblemID = "codeforces" + strconv.Itoa(jp.ContestID) + jp.Index
	bp.ProblemsetName = jp.ProblemsetName
	bp.Rating = jp.Rating
	bp.Tags = jp.Tags
	bp.Type = jp.Type
}

func AddCodeforcesProblem(jp CFProblem) {

	dbname := "harir-khobor"
	problemCollectionName := "judgeProblems"

	var BsonProblem ProcessedCodeforcesProblem
	ConvertCodeforcesProblem(jp, &BsonProblem)

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
	collection := client.Database(dbname).Collection(problemCollectionName)
	fmt.Println("Connected to Collections")

	// Inserting a single Document
	insertResult, err := collection.InsertOne(context.TODO(), BsonProblem)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

func UpdateCodeforcesProblem(jp CFProblem) {
	var BsonProblem ProcessedCodeforcesProblem

	ConvertCodeforcesProblem(jp, &BsonProblem)

	dbname := "harir-khobor"
	problemCollectionName := "judgeProblems"

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
	collection := client.Database(dbname).Collection(problemCollectionName)
	fmt.Println("Connected to Collections")

	filter := bson.D{{"problemId", "codeforces" + strconv.Itoa(jp.ContestID) + jp.Index}}
	fmt.Println("Filter Create Successful")

	// Inserting a single Document
	fmt.Println("Update rating = ", jp.Rating)
	update := bson.D{
		{"$set", bson.D{
			{"contestId", jp.ContestID},
			{"problemsetName", jp.ProblemsetName},
			{"index", jp.Index},
			{"name", jp.Name},
			{"type", jp.Type},
			{"points", jp.Points},
			{"rating", jp.Rating},
			{"tags", jp.Tags},
		}},
	}

	fmt.Println("Update Create SucceCFProblemssful")

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)

	fmt.Println("After update operation")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Matched %v documents and update %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")

	if updateResult.MatchedCount == 0 {
		AddCodeforcesProblem(jp)
	}
}

func main() {
	//fmt.Println("hello")

	var JsonProblem CFProblem
	//var BsonProblem ProcessedCodeforcesProblem

	JsonProblem.ContestID = 1169
	JsonProblem.Index = "C"
	JsonProblem.Name = "Pairs"
	JsonProblem.Points = 1000
	JsonProblem.Rating = 1500
	JsonProblem.Type = "PROGRAMMING"
	JsonProblem.Tags = []string{"graphs", "implementation"}

	UpdateCodeforcesProblem(JsonProblem)
}
