package repository

import (
	"context"
	"jrpg-gang/persistance/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	MongoDBRepository[model.UserModel]
}

func NewUserRepository(collection *mongo.Collection) *UserRepository {
	r := &UserRepository{}
	r.collection = collection
	return r
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*model.UserModel, bool) {
	filter := primitive.D{{Key: "email", Value: email}}
	return r.FindOne(ctx, filter)
}

func (r *UserRepository) FindByNickname(ctx context.Context, nickname string) (*model.UserModel, bool) {
	filter := primitive.D{{Key: "nickname", Value: nickname}}
	return r.FindOne(ctx, filter)
}

func (r *UserRepository) UpdateOneWithUnit(ctx context.Context, user model.UserModel) (int64, bool) {
	filter := primitive.D{{Key: "email", Value: user.Email}}
	fields := primitive.D{
		{Key: "class", Value: user.Class},
		{Key: "nickname", Value: user.Nickname},
		{Key: "unit", Value: user.Unit},
	}
	return r.UpdateOne(ctx, filter, fields)
}
