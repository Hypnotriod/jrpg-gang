package controller

import (
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
)

func (c *GameController) handleEnterLobbyRequest(userId engine.UserId, request *Request, response *Response) string {
	status := c.users.GetUserStatus(userId)
	if status == users.UserStatusInLobby {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	c.users.ChangeUserStatus(userId, users.UserStatusInLobby)
	return response.WithStatus(ResponseStatusOk)
}
