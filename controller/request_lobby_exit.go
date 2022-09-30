package controller

import (
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
)

type ExitLobbyRequest struct {
	Request
}

func (c *GameController) handleExitLobbyRequest(userId engine.UserId, requestRaw string, response *Response) string {
	request := parseRequest(&EnterLobbyRequest{}, requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMalformed)
	}
	status := c.users.GetUserStatus(userId)
	if status != users.UserStatusInLobby {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	c.users.ChangeUserStatus(userId, users.UserStatusJoined)
	return response.WithStatus(ResponseStatusOk)
}
