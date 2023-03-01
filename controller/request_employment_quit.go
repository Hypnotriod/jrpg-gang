package controller

import "jrpg-gang/engine"

func (c *GameController) handleQuitJobRequest(playerId engine.PlayerId, request *Request, response *Response) string {
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
	return response.WithStatus(ResponseStatusOk)
}
