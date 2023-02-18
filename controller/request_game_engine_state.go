package controller

import "jrpg-gang/engine"

func (c *GameController) handleGameStateRequest(playerId engine.PlayerId, request *Request, response *Response) string {
	wrapper, ok := c.engines.Find(playerId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	defer wrapper.RUnlock()
	wrapper.RLock()
	result, _, ok := wrapper.ReadGameState(playerId)
	if !ok {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	response.Data[DataKeyGameState] = result
	return response.WithStatus(ResponseStatusOk)
}
