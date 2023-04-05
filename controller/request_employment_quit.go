package controller

import (
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
)

func (c *GameController) handleQuitJobRequest(playerId engine.PlayerId, request *Request, response *Response) []byte {
	user, ok := c.users.Get(playerId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	jobStatus, ok := c.employment.QuitJob(&user)
	if !ok {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	if jobStatus != nil {
		c.persistJobStatus(user.Email, *jobStatus)
	}
	c.users.ChangeUserStatus(playerId, users.UserStatusJoined)
	return response.WithStatus(ResponseStatusOk)
}
