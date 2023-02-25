package controller

import (
	"jrpg-gang/controller/users"
	"jrpg-gang/persistance/model"
)

func (c *GameController) persistUser(user *users.User) bool {
	unit := user.Unit.ToPersist()
	userModel := model.UserModel{
		Email:    user.Email,
		Nickname: user.Nickname,
		Class:    user.Class,
		Unit:     &unit.Unit,
	}
	return c.persistance.UpdateUser(userModel)
}
