package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
}

func (m *Model) SetCreatedAt(t time.Time) {
	m.CreatedAt = t
}

func (m *Model) SetUpdatedAt(t time.Time) {
	m.UpdatedAt = t
}
