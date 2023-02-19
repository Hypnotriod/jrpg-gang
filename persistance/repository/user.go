package repository

import (
	"context"
	"jrpg-gang/persistance/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	Repository[models.UserModel]
}

func NewUserRepository(collection *mongo.Collection) *UserRepository {
	r := &UserRepository{}
	r.collection = collection
	return r
}

func (r *UserRepository) FindByEmail(ctx context.Context, email models.UserEmail) (*models.UserModel, bool) {
	return r.FindOne(ctx, primitive.M{"email": email})
}
