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

func Write(start_city string, end_city string, path [][]float32, client *mongo.Client, ctx context.Context) {
    var id string

    if end_city < start_city {
        id = end_city + start_city
    } else {
        id = start_city + end_city
    }

    collection := client.Database("CS575FP").Collection("paths")

    insert_res, err := collection.InsertOne(ctx, bson.D{
        {Key: "id", Value: id},
        {Key: "path", Value: path},
    })

    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Inserted: ", insert_res.InsertedID)
}

func Read(client *mongo.Client, ctx context.Context) {
    collection := client.Database("CS575FP").Collection("paths")
    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        fmt.Println(err)
        return
    }

    var paths []bson.M
    if err = cursor.All(ctx, &paths); err != nil {
        fmt.Println(err)
        return
    }

}

func Disconnect(client *mongo.Client, ctx context.Context) {
    client.Disconnect(ctx)
}
