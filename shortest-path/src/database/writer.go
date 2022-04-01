package database

import (
    "context"
    "time"
    "fmt"
//    "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const CONNECT_URI = "mongodb://localhost:27017/?readPreference=primary&appname=MongoDB%20Compass%20Community&ssl=false"

func Connect() bool {
    client, err := mongo.NewClient(options.Client().ApplyURI(CONNECT_URI))
    if err != nil {
        fmt.Println("Error creating MongoDB client")
        return false
    }

    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
        fmt.Println("Error on connect")
        return false
    }
    defer client.Disconnect(ctx)

    err = client.Ping(ctx, readpref.Primary())
    if err != nil {
        fmt.Println("Error on ping")
        return false
    }

    return true
}

func Write(start_city string, end_city string, path [][]float32) {
    client, err := mongo.NewClient(options.Client().ApplyURI(CONNECT_URI))
    if err != nil {
        fmt.Println(err)
        return
    }

    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer client.Disconnect(ctx)

    err = client.Ping(ctx, readpref.Primary())
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(start_city + end_city)
//    collection := client.Database("CS575FP").Collection("paths")
}
