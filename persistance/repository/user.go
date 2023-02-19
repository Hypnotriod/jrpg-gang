package repository

import (
	"jrpg-gang/persistance/models"

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
