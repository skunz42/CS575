package database

import (
    "context"
    "time"
    "fmt"
    "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const CONNECT_URI = "mongodb://localhost:27017/?readPreference=primary&appname=MongoDB%20Compass%20Community&ssl=false"

type Schema struct {
    Id string `bson: "id"`
    Path [][]float32 `bson: "path"`
}

func Connect() (*mongo.Client, context.Context) {
    client, err := mongo.NewClient(options.Client().ApplyURI(CONNECT_URI))
    if err != nil {
        fmt.Println("Error creating MongoDB client")
        return nil, nil
    }

    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
        fmt.Println("Error on connect")
        return nil, nil
    }

    err = client.Ping(ctx, readpref.Primary())
    if err != nil {
        fmt.Println("Error on ping")
        return nil, nil
    }

    return client, ctx
}

func Write(id string, path [][]float32, client *mongo.Client, ctx context.Context) {

    collection := client.Database("CS575FP").Collection("paths")

    to_ins := Schema{Id: id, Path:path}

    insert_res, err := collection.InsertOne(ctx, to_ins)

    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Inserted: ", insert_res.InsertedID)
}

func Read(client *mongo.Client, ctx context.Context, search_id string) ([][]float32) {
    var paths Schema

    collection := client.Database("CS575FP").Collection("paths")

    err := collection.FindOne(ctx, bson.M{"id":search_id}).Decode(&paths)
    if err != nil {
        fmt.Println("Path has not yet been computed")
        return nil
    }

    fmt.Println("Found cached path")

    return paths.Path
}

func Disconnect(client *mongo.Client, ctx context.Context) {
    client.Disconnect(ctx)
}
