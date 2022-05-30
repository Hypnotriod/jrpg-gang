package controller

import (
	"jrpg-gang/engine"
	"regexp"
)

type JoinRequest struct {
	Request
	Data struct {
		Nickname string               `json:"nickname"`
		Class    engine.GameUnitClass `json:"class"`
	} `json:"data"`
}

func (c *GameController) handleJoinRequest(requestRaw string, response *Response) (engine.UserId, string) {
	request := parseRequest(&JoinRequest{}, requestRaw)
	if request == nil {
		return engine.UserIdEmpty, response.WithStatus(ResponseStatusMailformed)
	}
	if matched, _ := regexp.MatchString(USER_NICKNAME_REGEX, request.Data.Nickname); !matched {
		return engine.UserIdEmpty, response.WithStatus(ResponseStatusNotAllowed)
	}
	if _, ok := c.users.GetByNickname(request.Data.Nickname); ok {
		return engine.UserIdEmpty, response.WithStatus(ResponseStatusAlreadyExists)
	}
	unit := NewGameUnitByClass(request.Data.Class) // todo: test purpose only
	if unit == nil {
		return engine.UserIdEmpty, response.WithStatus(ResponseStatusMailformed)
	}
	user := NewUser(request.Data.Nickname,
		request.Data.Class,
		unit)
	c.users.AddUser(user)
	response.Data[DataKeyUserNickname] = user.Nickname
	response.Data[DataKeyUserId] = user.id
	response.Data[DataKeyUnit] = user.unit
	return user.id, response.WithStatus(ResponseStatusOk)
}
