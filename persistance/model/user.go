package model

import (
	"jrpg-gang/domain"
)

type UserModel struct {
	Model    `bson:",inline"`
	Email    string           `bson:"email"`
	Class    domain.UnitClass `bson:"class,omitempty"`
	Nickname string           `bson:"nickname,omitempty"`
	Picture  string           `bson:"picture,omitempty"`
	Unit     *domain.Unit     `bson:"unit,omitempty"`
}
