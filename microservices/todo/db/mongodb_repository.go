package repository

import (
	"context"
	"errors"
	"todo/db/mongocli"

	pb "todo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBRepository struct {
	dbclient *mongo.Client
}

func NewMongoDBRepository() Repository {
	return &MongoDBRepository{
		dbclient: mongocli.GetDB(),
	}
}

func (repo *MongoDBRepository) Add(element pb.AddItemRequest) error {
	collection := repo.dbclient.Database("local").Collection("test")

	_, err := collection.InsertOne(context.TODO(), element)
	if err != nil {
		return err
	}

	// example: get oid.Hex
	// if oid, ok := insertResult.InsertedID.(primitive.ObjectID); ok {
	// 	return c.JSON(http.StatusCreated, map[string]interface{}{
	// 		"id": oid.Hex(),
	// 	})
	// }

	return nil
}

func (repo *MongoDBRepository) Delete(id string) error {
	collection := repo.dbclient.Database("local").Collection("test")

	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	result, err := collection.DeleteOne(context.TODO(), bson.M{"_id": idPrimitive})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("Id not found")
	}

	return nil
}

func (repo *MongoDBRepository) FindAll() (pb.ListResponse, error) {

	collection := repo.dbclient.Database("local").Collection("test")

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return pb.ListResponse{}, err
	}

	var elements pb.ListResponse
	if err := cursor.All(context.Background(), &elements.Items); err != nil {
		return pb.ListResponse{}, err
	}

	return elements, nil
}
