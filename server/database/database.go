package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func Connect() (*mongo.Client, context.Context, context.CancelFunc) {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	uri := fmt.Sprintf("mongodb+srv://%s:%s@cluster1.ie0rc.mongodb.net/?retryWrites=true&w=majority", username, password)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal(err)
	}

	return client, ctx, cancel
}

func TestConnection(client *mongo.Client, ctx context.Context) {
	err := client.Ping(ctx, nil)

	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		fmt.Println("Connected!")
	}
}
