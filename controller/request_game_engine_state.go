package controller

import "jrpg-gang/engine"

func (c *GameController) handleGameStateRequest(userId engine.UserId, request *Request, response *Response) string {
	result, _, ok := c.engines.GameState(userId)
	if !ok {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	response.Data[DataKeyGameState] = result
	return response.WithStatus(ResponseStatusOk)
}
