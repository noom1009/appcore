package database

import (
    "context" // เพิ่มการ import นี้
    "log"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB(uri string) *mongo.Client {
    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
    if err != nil {
        log.Fatalf("MongoDB connection error: %v", err)
    }
    return client
}
