package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model struct {
	Id        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

func (m *Model) OnCreate() {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
}

func (m *Model) OnUpdate() {
	m.UpdatedAt = time.Now()
}
