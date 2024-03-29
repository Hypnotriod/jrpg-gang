package model

import (
	"jrpg-gang/domain"
)

type UserModel struct {
	Model    `bson:",inline"`
	Email    UserEmail                         `bson:"email"`
	Class    domain.UnitClass                  `bson:"class,omitempty"`
	Nickname string                            `bson:"nickname,omitempty"`
	Picture  string                            `bson:"picture,omitempty"`
	Unit     *domain.Unit                      `bson:"-"`
	Units    map[domain.UnitClass]*domain.Unit `bson:"units,omitempty"`
}
