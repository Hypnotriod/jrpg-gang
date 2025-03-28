package repository

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ObjectId string

const (
	ObjectIdEmpty ObjectId = ""
)

type MongoDBRepositoryModelPointer[T any] interface {
	SetCreatedAt(t time.Time)
	SetUpdatedAt(t time.Time)
	*T
}

type MongoDBRepository[T any, P MongoDBRepositoryModelPointer[T]] struct {
	collection *mongo.Collection
}

func (r *MongoDBRepository[T, P]) InsertOne(ctx context.Context, model P) (ObjectId, bool) {
	t := time.Now()
	model.SetCreatedAt(t)
	model.SetUpdatedAt(t)
	result, err := r.collection.InsertOne(ctx, model)
	if err != nil {
		log.Error("Mongodb: InsertOne (", model, ") to collection:", r.collection.Name(), " fail: ", err)
		return ObjectIdEmpty, false
	}
	if objectId, ok := result.InsertedID.(primitive.ObjectID); ok {
		return ObjectId(objectId.Hex()), true
	}
	return ObjectIdEmpty, false
}

func (r *MongoDBRepository[T, P]) UpdateOneById(ctx context.Context, id ObjectId, fields bson.D) (int64, bool) {
	objectId, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		log.Error("Mongodb: UpdateOneById: ", id, " in collection:", r.collection.Name(), " fail: ", err)
		return 0, false
	}

	fields = append(fields, bson.E{Key: "updated_at", Value: time.Now()})
	filter := bson.D{{Key: "_id", Value: objectId}}
	update := bson.D{{Key: "$set", Value: fields}}
	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Error("Mongodb: UpdateOneById: ", id, " in collection:", r.collection.Name(), " fail: ", err)
		return 0, false
	}
	return result.MatchedCount, true
}

func (r *MongoDBRepository[T, P]) UpdateOne(ctx context.Context, filter bson.D, fields bson.D) (int64, bool) {
	fields = append(fields, bson.E{Key: "updated_at", Value: time.Now()})
	update := bson.D{{Key: "$set", Value: fields}}
	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Error("Mongodb: UpdateOne: ", filter, " in collection:", r.collection.Name(), " fail: ", err)
		return 0, false
	}
	return result.MatchedCount, true
}

func (r *MongoDBRepository[T, P]) FindOneById(ctx context.Context, id ObjectId, accumulator P) (P, bool) {
	objectId, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		log.Error("Mongodb: FindOneById: ", id, " in collection:", r.collection.Name(), " fail: ", err)
		return nil, false
	}

	filter := bson.D{{Key: "_id", Value: objectId}}
	err = r.collection.FindOne(ctx, filter).Decode(accumulator)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			log.Error("Mongodb: FindOneById: ", id, " in collection:", r.collection.Name(), " fail: ", err)
		}
		return nil, false
	}
	return accumulator, true
}

func (r *MongoDBRepository[T, P]) FindOne(ctx context.Context, filter bson.D, accumulator P) (P, bool) {
	err := r.collection.FindOne(ctx, filter).Decode(accumulator)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, true
		}
		log.Error("Mongodb: FindOne: ", filter, " in collection:", r.collection.Name(), " fail: ", err)
		return nil, false
	}
	return accumulator, true
}

func (r *MongoDBRepository[T, P]) DeleteOneById(ctx context.Context, id ObjectId) (int64, bool) {
	objectId, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		log.Error("Mongodb: DeleteOneById: ", id, " from collection:", r.collection.Name(), " fail: ", err)
		return 0, false
	}

	filter := bson.D{{Key: "_id", Value: objectId}}
	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Error("Mongodb: DeleteOneById: ", id, " from collection:", r.collection.Name(), " fail: ", err)
		return 0, false
	}
	return result.DeletedCount, true
}
