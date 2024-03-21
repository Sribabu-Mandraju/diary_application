package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBconnection() *mongo.Client {
	mongoDb := "mongodb+srv://sribabu:63037sribabu@atlascluster.k6u2oy9.mongodb.net/?retryWrites=true&w=majority"

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoDb))

	if err != nil {
		fmt.Println("unable to connect to database")
		return nil
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		fmt.Println("network error")
		return nil
	}

	fmt.Println("Successfully connected to database")
	return client

}


