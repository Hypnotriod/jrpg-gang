package repository

import (
	"context"
	"jrpg-gang/persistance/models"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RepositoryModel interface {
	*models.UserModel
	OnCreate()
	OnUpdate()
}

type Repository[T RepositoryModel] struct {
	collection *mongo.Collection
}

func (r *Repository[T]) InsertOne(ctx context.Context, model T) primitive.ObjectID {
	model.OnCreate()
	result, err := r.collection.InsertOne(ctx, model)
	if err != nil {
		log.Error("Mongodb: InsertOne (", model, ") fail: ", err)
		return primitive.NilObjectID
	}
	if objectId, ok := result.InsertedID.(primitive.ObjectID); ok {
		return objectId
	}
	return primitive.NilObjectID
}
