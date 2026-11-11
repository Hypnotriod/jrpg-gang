package controller

import (
	"jrpg-gang/engine"

	log "github.com/sirupsen/logrus"
)

func (c *GameController) handleDeleteAccountRequest(playerId engine.PlayerId, request *Request, response *Response) []byte {
	user, ok := c.users.Get(playerId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	if result := c.persistance.DeleteUserByEmail(user.Email); !result {
		return response.WithStatus(ResponseStatusError)
	}
	log.Info("The user account has been deleted: ", user.Id)
	c.Leave(playerId)
	return response.WithStatus(ResponseStatusOk)
}
