package controller

import (
	"jrpg-gang/auth"
	"jrpg-gang/controller/users"
	"jrpg-gang/domain"
	"jrpg-gang/engine"
	"jrpg-gang/persistance/model"
	"regexp"
)

type JoinRequestData struct {
	Token    auth.AuthenticationToken `json:"token,omitempty"`
	Nickname string                   `json:"nickname,omitempty"`
	Class    engine.GameUnitClass     `json:"class,omitempty"`
	PlayerId engine.PlayerId          `json:"playerId,omitempty"`
}

func (c *GameController) handleJoinRequest(request *Request, response *Response) (engine.PlayerId, string) {
	data := parseRequestData(&JoinRequestData{}, request.Data)
	if data == nil {
		return engine.PlayerIdEmpty, response.WithStatus(ResponseStatusMalformed)
	}
	if data.PlayerId != engine.PlayerIdEmpty {
		user, ok := c.users.Get(data.PlayerId)
		if !ok {
			return engine.PlayerIdEmpty, response.WithStatus(ResponseStatusNotAllowed)
		}
		user.IsOffline = false
		response.fillUserStatus(&user)
		return user.Id, response.WithStatus(ResponseStatusOk)
	}
	userModel, ok := c.persistance.GetUserFromCache(data.Token)
	if userModel == nil || !ok {
		return engine.PlayerIdEmpty, response.WithStatus(ResponseStatusNotAllowed)
	}
	var unit *engine.GameUnit
	var nickname string
	var class engine.GameUnitClass
	if userModel.Unit != nil {
		unit = engine.NewGameUnit(userModel.Unit)
		c.itemsConfig.PopulateFromDescriptor(&unit.Inventory)
		nickname = userModel.Nickname
		class = userModel.Class
	} else {
		if matched, _ := regexp.MatchString(USER_NICKNAME_REGEX, data.Nickname); !matched {
			return engine.PlayerIdEmpty, response.WithStatus(ResponseStatusNotAllowed)
		}
		if ok := c.persistance.HasUserWithNickname(data.Nickname); ok {
			return engine.PlayerIdEmpty, response.WithStatus(ResponseStatusAlreadyExists)
		}
		unit = c.unitsConfig.GetByCode(domain.UnitCode(data.Class))
		nickname = data.Nickname
		class = data.Class
	}
	if unit == nil {
		return engine.PlayerIdEmpty, response.WithStatus(ResponseStatusMalformed)
	}
	unit.PrepareForUser()
	user := users.NewUser(nickname, userModel.Email, class, unit)
	if userModel.Unit == nil {
		c.persistUser(user)
	}
	c.users.AddUser(user)
	response.fillUserStatus(user)
	return user.Id, response.WithStatus(ResponseStatusOk)
}

func (c *GameController) persistUser(user *users.User) {
	unit := user.Unit.ToPersist()
	userModel := model.UserModel{
		Email:    user.Email,
		Nickname: user.Nickname,
		Class:    user.Class,
		Unit:     &unit.Unit,
	}
	c.persistance.UpdateUser(userModel)
}
