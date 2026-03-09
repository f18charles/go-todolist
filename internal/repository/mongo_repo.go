package repository

import (
	"context"
	"todolist/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepo struct {
	Collection *mongo.Collection
}

func NewMongoRepo(c *mongo.Collection) *MongoRepo {
	return &MongoRepo{Collection: c}
}

func (r *MongoRepo) Create(todo models.Todo) error {
	_, err := r.Collection.InsertOne(context.TODO(), todo)
	return err
}

func (r *MongoRepo) GetByID(id string) (models.Todo, error) {
	var todo models.Todo

	err := r.Collection.FindOne(
		context.TODO(),
		bson.M{"_id": id},
	).Decode(&todo)

	return todo, err
}

func (r *MongoRepo) GetAll() ([]models.Todo, error) {
	cursor, err := r.Collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	var todos []models.Todo
	err = cursor.All(context.TODO(), &todos)

	return todos, err
}

func (r *MongoRepo) Update(todo models.Todo) error {
	_, err := r.Collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": todo.ID},
		bson.M{"$set": todo},
	)
	return err
}

func (r *MongoRepo) Delete(id string) error {
	_, err := r.Collection.DeleteOne(
		context.TODO(),
		bson.M{"_id": id},
	)
	return err
}
