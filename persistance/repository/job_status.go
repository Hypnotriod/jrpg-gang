package repository

import (
	"context"
	"jrpg-gang/persistance/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type JobStatusRepository struct {
	MongoDBRepository[*model.JobStatusModel]
}

func NewJobStatusRepository(collection *mongo.Collection) *JobStatusRepository {
	r := &JobStatusRepository{}
	r.collection = collection
	return r
}

func (r *JobStatusRepository) FindByUserId(ctx context.Context, userId model.UserId) (*model.JobStatusModel, bool) {
	filter := primitive.D{{Key: "user_id", Value: userId}}
	return r.FindOne(ctx, filter, &model.JobStatusModel{})
}

func (r *JobStatusRepository) UpdateOrInsertOne(ctx context.Context, model *model.JobStatusModel) (int64, bool) {
	filter := primitive.D{{Key: "user_id", Value: model.UserId}}
	fields := primitive.D{
		{Key: "is_in_progress", Value: model.IsInProgress},
		{Key: "is_complete", Value: model.IsComplete},
		{Key: "completion_time", Value: model.CompletionTime},
		{Key: "code", Value: model.Code},
		{Key: "countdown", Value: model.Countdown},
	}
	matchedCount, ok := r.UpdateOne(ctx, filter, fields)
	if !ok || matchedCount != 0 {
		return matchedCount, ok
	}
	if _, ok := r.InsertOne(ctx, model); ok {
		return 1, true
	}
	return 0, false
}
