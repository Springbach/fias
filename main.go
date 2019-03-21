package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

const (
	mongodbURL = "mongodb://localhost:4320"
)

func init() {
	fmt.Println("Attemp to connect to MongoDB!")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI(mongodbURL))

	// Check the connection
	ctx, _ = context.WithTimeout(context.Background(), 20*time.Second)
	err := client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("SUCCESS! Connected to MongoDB.")
}

func main() {
	headersList := map[string]string{"Content-type": "application/json"}
	//GET list/regions
	http.Handle("/list/regions", GET(SetHeaders(List(Regions{}), headersList)))

	//GET list/citiies URL query reg REQUIRED
	http.Handle("/list/cities", GET(SetHeaders(QValid(List(Cities{}), []string{"reg"}), headersList)))

	//GET list/streets URL query reg && city REQUIRED
	http.Handle("/list/streets", GET(SetHeaders(QValid(List(Streets{}), []string{"reg", "city"}), headersList)))

	//GET aggregate city, streets by zipcode// query zip REQUIRED
	http.Handle("/list/territory", GET(SetHeaders(QValid(List(Territory{}), []string{"zip"}), headersList)))

	//GET aggregate houses by zipcode and street// query zip and street REQUIRED
	http.Handle("/list/houses", GET(SetHeaders(QValid(List(Houses{}), []string{"zip"}), headersList)))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
