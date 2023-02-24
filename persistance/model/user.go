package model

import (
	"jrpg-gang/domain"
	"jrpg-gang/engine"
)

type UserModel struct {
	Model
	Email    string               `bson:"email"`
	Class    engine.GameUnitClass `json:"class,omitempty"`
	Nickname string               `bson:"nickname,omitempty"`
	Picture  string               `bson:"picture,omitempty"`
	Unit     *domain.Unit         `bson:"unit,omitempty"`
}
