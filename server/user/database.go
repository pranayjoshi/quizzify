package user

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/joho/godotenv"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// func init() {
// 	if err := godotenv.Load(); err != nil {
// 		log.Print("No .env file found")
// 	}
// }

// var client *mongo.Client
// var err error

// func ConnectDB() (*mongo.Client, context.Context, error) {
// 	uri := os.Getenv("URI")
// 	fmt.Println(uri)

// 	opts := options.Client().ApplyURI(uri)
// 	// Create a new client and connect to the server
// 	client, err := mongo.Connect(context.TODO(), opts)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer func() {
// 		if err = client.Disconnect(context.TODO()); err != nil {
// 			panic(err)
// 		}
// 	}()
// 	// Send a ping to confirm a successful connection
// 	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
// 	return client, nil, nil
// }
