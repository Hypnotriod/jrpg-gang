package controller

import "jrpg-gang/engine"

func (c *GameController) handleGameStateRequest(userId engine.UserId, request *Request, response *Response) string {
	wrapper, ok := c.engines.Find(userId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	defer wrapper.RUnlock()
	wrapper.RLock()
	result, _, ok := wrapper.ReadGameState(userId)
	if !ok {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	response.Data[DataKeyGameState] = result
	return response.WithStatus(ResponseStatusOk)
}
