package model

import (
	"time"
)

type Model struct {
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

func (m *Model) OnCreate() {
	t := time.Now()
	m.CreatedAt = t
	m.UpdatedAt = t
}
