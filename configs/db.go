package configs

import (
	"context"
	"log"

	"github.com/heyyakash/realtime-weather-aggregator/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	WeatherDataCollection    *mongo.Collection
	CalculatedDataCollection *mongo.Collection
)

func ConnectDB() {
	connectionString := helpers.GetEnv("MONGO_URL")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(connectionString).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	log.Println("Pinged your deployment. You successfully connected to MongoDB!")
	WeatherDataCollection = client.Database("realtime-weather-aggregator").Collection("weather-data")
	CalculatedDataCollection = client.Database("realtime-weather-aggregator").Collection("calculated-data")
}
