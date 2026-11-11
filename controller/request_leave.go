package controller

import "jrpg-gang/engine"

func (c *GameController) handleLeaveRequest(playerId engine.PlayerId, request *Request, response *Response) []byte {
	c.Leave(playerId)
	return response.WithStatus(ResponseStatusOk)
}
