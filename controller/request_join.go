package controller

import (
	"jrpg-gang/controller/users"
	"jrpg-gang/domain"
	"jrpg-gang/engine"
	"regexp"
)

type JoinRequestData struct {
	Nickname string               `json:"nickname"`
	Class    engine.GameUnitClass `json:"class"`
	UserId   engine.UserId        `json:"userId,omitempty"`
}

func (c *GameController) handleJoinRequest(request *Request, response *Response) (engine.UserId, string) {
	data := parseRequestData(&JoinRequestData{}, request.Data)
	if data == nil {
		return engine.UserIdEmpty, response.WithStatus(ResponseStatusMalformed)
	}
	if data.UserId != engine.UserIdEmpty {
		user, ok := c.users.Get(data.UserId)
		if !ok {
			return engine.UserIdEmpty, response.WithStatus(ResponseStatusNotAllowed)
		}
		user.IsOffline = false
		response.fillUserStatus(&user)
		return user.Id, response.WithStatus(ResponseStatusOk)
	}
	if matched, _ := regexp.MatchString(USER_NICKNAME_REGEX, data.Nickname); !matched {
		return engine.UserIdEmpty, response.WithStatus(ResponseStatusNotAllowed)
	}
	if _, ok := c.users.GetByNickname(data.Nickname); ok {
		return engine.UserIdEmpty, response.WithStatus(ResponseStatusAlreadyExists)
	}
	unit := c.unitsConfig.GetByCode(domain.UnitCode(data.Class))
	if unit == nil {
		return engine.UserIdEmpty, response.WithStatus(ResponseStatusMalformed)
	}
	unit.PrepareForUser()
	user := users.NewUser(data.Nickname, data.Class, unit)
	c.users.AddUser(user)
	response.fillUserStatus(user)
	return user.Id, response.WithStatus(ResponseStatusOk)
}
