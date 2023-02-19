package repository

import (
	"context"
	"jrpg-gang/persistance/models"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RepositoryObjectId string

const (
	RepositoryObjectIdEmpty RepositoryObjectId = ""
)

type RepositoryModel interface {
	models.UserModel
}

type Repository[T RepositoryModel] struct {
	collection *mongo.Collection
}

func (r *Repository[T]) InsertOne(ctx context.Context, model T) RepositoryObjectId {
	result, err := r.collection.InsertOne(ctx, model)
	if err != nil {
		log.Error("Mongodb: InsertOne (", model, ") to collection:", r.collection.Name(), " fail: ", err)
		return RepositoryObjectIdEmpty
	}
	if objectId, ok := result.InsertedID.(primitive.ObjectID); ok {
		return RepositoryObjectId(objectId.Hex())
	}
	return RepositoryObjectIdEmpty
}

func (r *Repository[T]) UpdateOneById(ctx context.Context, id RepositoryObjectId, fields bson.D) int64 {
	objectId, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		log.Error("Mongodb: UpdateOneById: ", id, " in collection:", r.collection.Name(), " fail: ", err)
		return 0
	}

	fields = append(fields, bson.E{Key: "updated_at", Value: time.Now()})
	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": fields}
	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Error("Mongodb: UpdateOneById: ", id, " in collection:", r.collection.Name(), " fail: ", err)
		return 0
	}
	return result.MatchedCount
}
func (r *Repository[T]) UpdateOne(ctx context.Context, filter bson.M, fields bson.D) int64 {
	fields = append(fields, bson.E{Key: "updated_at", Value: time.Now()})
	update := bson.M{"$set": fields}
	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Error("Mongodb: UpdateOne: ", filter, " in collection:", r.collection.Name(), " fail: ", err)
		return 0
	}
	return result.MatchedCount
}

func (r *Repository[T]) FindOneById(ctx context.Context, id RepositoryObjectId, fields bson.D) *T {
	objectId, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		log.Error("Mongodb: FindOneById: ", id, " in collection:", r.collection.Name(), " fail: ", err)
		return nil
	}

	result := &T{}
	filter := bson.M{"_id": objectId}
	err = r.collection.FindOne(ctx, filter).Decode(result)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			log.Error("Mongodb: FindOneById: ", id, " in collection:", r.collection.Name(), " fail: ", err)
		}
		return nil
	}
	return result
}

func (r *Repository[T]) FindOne(ctx context.Context, filter bson.M, fields bson.D) *T {
	result := &T{}
	err := r.collection.FindOne(ctx, filter).Decode(result)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			log.Error("Mongodb: FindOneById: ", filter, " in collection:", r.collection.Name(), " fail: ", err)
		}
		return nil
	}
	return result
}

func (r *Repository[T]) DeleteOneById(ctx context.Context, id RepositoryObjectId) int64 {
	objectId, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		log.Error("Mongodb: DeleteOneById: ", id, " from collection:", r.collection.Name(), " fail: ", err)
		return 0
	}

	filter := bson.M{"_id": objectId}
	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Error("Mongodb: DeleteOneById: ", id, " from collection:", r.collection.Name(), " fail: ", err)
		return 0
	}
	return result.DeletedCount
}
