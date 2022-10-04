package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func LoadEnv(key string) interface{} {
	if os.Getenv(key) == "" {
		err := godotenv.Load(".env")

		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	value := os.Getenv(key)

	return value

}

func DB() *mongo.Database {
	clientOptions := options.Client()
	uri := LoadEnv("MONGODB_URI").(string)
	clientOptions.ApplyURI(uri)

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Panic("mongo client not connected : " + err.Error())
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Panic("mongo connection failed: " + err.Error())
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		fmt.Println("mongodb failed...")
		// log.Panic(err)
	}
	return client.Database("dashboard")
}
