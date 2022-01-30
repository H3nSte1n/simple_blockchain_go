package db

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

//Create creating a task in a mongo or document db
func Create[T any](object T, model string) (*int32, error) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	result, err := client.Database(model).Collection(model).InsertOne(ctx, object)
	if err != nil {
		log.Printf("Could not create %q: %v", model, err)
		return nil, err
	}

	oid := result.InsertedID.(int32)
	return &oid, nil
}

//Create creating a task in a mongo or document db
func GetList[T any](model string) ([]T, error) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	cur, err := client.Database(model).Collection(model).Find(ctx, bson.D{})

	if err != nil {
		log.Printf("Could not fetch %q: %v", model, err)
		return nil, err
	}

	res := make([]T, 0)
	if err := cur.All(ctx, &res); err != nil {
		log.Printf("Could not fetch %q: %v", model, err)
		return nil, err
	}

	return res, err
}

func ResetCollection(model string) error {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	_, err := client.Database(model).Collection(model).DeleteMany(ctx, bson.D{})
	if err != nil {
		log.Printf("Could not fetch %q: %v", model, err)
		return err
	}

	return nil
}
