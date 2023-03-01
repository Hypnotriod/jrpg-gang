package repository

import (
	"context"
	"jrpg-gang/persistance/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type JobStatusRepository struct {
	MongoDBRepository[model.JobStatusModel]
}

func NewJobStatusRepository(collection *mongo.Collection) *JobStatusRepository {
	r := &JobStatusRepository{}
	r.collection = collection
	return r
}

func (r *JobStatusRepository) FindByEmail(ctx context.Context, email string) (*model.JobStatusModel, bool) {
	filter := primitive.D{{Key: "email", Value: email}}
	return r.FindOne(ctx, filter, &model.JobStatusModel{})
}

func (r *JobStatusRepository) UpdateOrInsertOne(ctx context.Context, model model.JobStatusModel) (int64, bool) {
	filter := primitive.D{{Key: "email", Value: model.Email}}
	fields := primitive.D{
		{Key: "is_in_progress", Value: model.IsInProgress},
		{Key: "is_complete", Value: model.IsComplete},
		{Key: "completion_time", Value: model.CompletionTime},
		{Key: "countdown", Value: model.Countdown},
	}
	matchedCount, ok := r.UpdateOne(ctx, filter, fields)
	if !ok || matchedCount != 0 {
		return matchedCount, ok
	}
	model.OnCreate()
	if _, ok := r.InsertOne(ctx, model); ok {
		return 1, true
	}
	return 0, false
}
