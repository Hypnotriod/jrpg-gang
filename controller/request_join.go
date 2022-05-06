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

func (c *GameController) handleJoinRequest(requestRaw string, response *Response) string {
	request := parseRequest(&JoinRequest{}, requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMailformed)
	}
	if matched, _ := regexp.MatchString(USER_NICKNAME_REGEX, request.Data.Nickname); !matched {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	if _, ok := c.users.GetByNickname(request.Data.Nickname); ok {
		return response.WithStatus(ResponseStatusAlreadyExists)
	}
	unit := NewGameUnitByClass(request.Data.Class) // todo: test purpose only
	if unit == nil {
		return response.WithStatus(ResponseStatusMailformed)
	}
	user := NewUser(request.Data.Nickname,
		request.Data.Class,
		unit)
	c.users.AddUser(user)
	response.Data[DataKeyUserId] = user.id
	response.Data[DataKeyUnit] = unit
	return response.WithStatus(ResponseStatusOk)
}
