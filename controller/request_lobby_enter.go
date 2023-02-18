package controller

import (
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
)

func (c *GameController) handleEnterLobbyRequest(playerId engine.PlayerId, request *Request, response *Response) string {
	status := c.users.GetUserStatus(playerId)
	if status == users.UserStatusInLobby {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	c.users.ChangeUserStatus(playerId, users.UserStatusInLobby)
	return response.WithStatus(ResponseStatusOk)
}
