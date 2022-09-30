package controller

import (
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
)

type EnterLobbyRequest struct {
	Request
}

func (c *GameController) handleEnterLobbyRequest(userId engine.UserId, requestRaw string, response *Response) string {
	request := parseRequest(&EnterLobbyRequest{}, requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMalformed)
	}
	status := c.users.GetUserStatus(userId)
	if status == users.UserStatusInLobby {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	c.users.ChangeUserStatus(userId, users.UserStatusInLobby)
	return response.WithStatus(ResponseStatusOk)
}
