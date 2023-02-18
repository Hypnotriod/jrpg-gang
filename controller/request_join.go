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
	PlayerId engine.PlayerId      `json:"playerId,omitempty"`
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
	if matched, _ := regexp.MatchString(USER_NICKNAME_REGEX, data.Nickname); !matched {
		return engine.PlayerIdEmpty, response.WithStatus(ResponseStatusNotAllowed)
	}
	if _, ok := c.users.GetByNickname(data.Nickname); ok {
		return engine.PlayerIdEmpty, response.WithStatus(ResponseStatusAlreadyExists)
	}
	unit := c.unitsConfig.GetByCode(domain.UnitCode(data.Class))
	if unit == nil {
		return engine.PlayerIdEmpty, response.WithStatus(ResponseStatusMalformed)
	}
	unit.PrepareForUser()
	user := users.NewUser(data.Nickname, data.Class, unit)
	c.users.AddUser(user)
	response.fillUserStatus(user)
	return user.Id, response.WithStatus(ResponseStatusOk)
}
