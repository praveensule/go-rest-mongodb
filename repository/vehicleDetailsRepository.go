package repository

import (
	"context"
	"fmt"
	. "go-rest-mongodb/config"
	"go-rest-mongodb/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type VehicleDetailsRepository struct {}

var config Config
var collection = new(mongo.Collection)
const VehicleDetailsCollection = "VehicleDetails"

func init() {
	config.Read()

	// Connect to DB
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Database.Uri))
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database(config.Database.DatabaseName).Collection(VehicleDetailsCollection)
}

// Get all VehicleDetails
func (p *VehicleDetailsRepository) FindAll() ([]models.VehicleDetails, error) {
	var vehicleDetails []models.VehicleDetails

	findOptions := options.Find()
	findOptions.SetLimit(100)

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	// Finding multiple documents returns a cursor
	cur, err := collection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Iterate through the cursor
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result models.VehicleDetails
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		vehicleDetails = append(vehicleDetails, result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return vehicleDetails, err
}

// Create a new VehicleDetails
func (p *VehicleDetailsRepository) Insert(vehicleDetails models.VehicleDetails) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collection.InsertOne(ctx, &vehicleDetails)
	fmt.Println("Inserted a single document: ", result.InsertedID)
	return result.InsertedID, err
}

// Delete an existing VehicleDetails
func (p *VehicleDetailsRepository) Delete(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collection.DeleteOne(ctx, filter)
	fmt.Println("Deleted a single document: ", result.DeletedCount)
	return err
}
