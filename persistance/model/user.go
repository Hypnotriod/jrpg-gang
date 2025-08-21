package model

import (
	"jrpg-gang/domain"
	"net"
)

type UserModel struct {
	Model    `bson:",inline"`
	Email    UserEmail                         `bson:"email"`
	Class    domain.UnitClass                  `bson:"class,omitempty"`
	Nickname string                            `bson:"nickname,omitempty"`
	Picture  string                            `bson:"picture,omitempty"`
	Ip       net.IP                            `bson:"-"`
	Unit     *domain.Unit                      `bson:"-"`
	Units    map[domain.UnitClass]*domain.Unit `bson:"units,omitempty"`
}
