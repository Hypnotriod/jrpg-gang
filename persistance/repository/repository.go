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

type MongoDBObjectId string

const (
	MongoDBObjectIdEmpty MongoDBObjectId = ""
)

type MongoDBRepositoryModel interface {
	models.UserModel
}

type MongoDBRepository[T MongoDBRepositoryModel] struct {
	collection *mongo.Collection
}

func (r *MongoDBRepository[T]) InsertOne(ctx context.Context, model T) (MongoDBObjectId, bool) {
	result, err := r.collection.InsertOne(ctx, model)
	if err != nil {
		log.Error("Mongodb: InsertOne (", model, ") to collection:", r.collection.Name(), " fail: ", err)
		return MongoDBObjectIdEmpty, false
	}
	if objectId, ok := result.InsertedID.(primitive.ObjectID); ok {
		return MongoDBObjectId(objectId.Hex()), true
	}
	return MongoDBObjectIdEmpty, false
}

func (r *MongoDBRepository[T]) UpdateOneById(ctx context.Context, id MongoDBObjectId, fields bson.D) (int64, bool) {
	objectId, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		log.Error("Mongodb: UpdateOneById: ", id, " in collection:", r.collection.Name(), " fail: ", err)
		return 0, false
	}

	fields = append(fields, bson.E{Key: "updated_at", Value: time.Now()})
	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": fields}
	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Error("Mongodb: UpdateOneById: ", id, " in collection:", r.collection.Name(), " fail: ", err)
		return 0, false
	}
	return result.MatchedCount, true
}
func (r *MongoDBRepository[T]) UpdateOne(ctx context.Context, filter bson.M, fields bson.D) (int64, bool) {
	fields = append(fields, bson.E{Key: "updated_at", Value: time.Now()})
	update := bson.M{"$set": fields}
	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Error("Mongodb: UpdateOne: ", filter, " in collection:", r.collection.Name(), " fail: ", err)
		return 0, false
	}
	return result.MatchedCount, true
}

func (r *MongoDBRepository[T]) FindOneById(ctx context.Context, id MongoDBObjectId) (*T, bool) {
	objectId, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		log.Error("Mongodb: FindOneById: ", id, " in collection:", r.collection.Name(), " fail: ", err)
		return nil, false
	}

	result := &T{}
	filter := bson.M{"_id": objectId}
	err = r.collection.FindOne(ctx, filter).Decode(result)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			log.Error("Mongodb: FindOneById: ", id, " in collection:", r.collection.Name(), " fail: ", err)
		}
		return nil, false
	}
	return result, true
}

func (r *MongoDBRepository[T]) FindOne(ctx context.Context, filter bson.M) (*T, bool) {
	result := &T{}
	err := r.collection.FindOne(ctx, filter).Decode(result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, true
		}
		log.Error("Mongodb: FindOneById: ", filter, " in collection:", r.collection.Name(), " fail: ", err)
		return nil, false
	}
	return result, false
}

func (r *MongoDBRepository[T]) DeleteOneById(ctx context.Context, id MongoDBObjectId) (int64, bool) {
	objectId, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		log.Error("Mongodb: DeleteOneById: ", id, " from collection:", r.collection.Name(), " fail: ", err)
		return 0, false
	}

	filter := bson.M{"_id": objectId}
	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Error("Mongodb: DeleteOneById: ", id, " from collection:", r.collection.Name(), " fail: ", err)
		return 0, false
	}
	return result.DeletedCount, true
}
