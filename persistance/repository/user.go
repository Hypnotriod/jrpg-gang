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
	return r.FindOne(ctx, primitive.M{"email": email})
}
