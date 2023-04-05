package controller

import "jrpg-gang/engine"

func (c *GameController) handlePlayerInfoRequest(playerId engine.PlayerId, request *Request, response *Response) []byte {
	wrapper, ok := c.engines.Find(playerId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	defer wrapper.RUnlock()
	wrapper.RLock()
	result, ok := wrapper.ReadPlayerInfo(playerId)
	if !ok {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	response.Data[DataKeyPlayerInfo] = result
	return response.WithStatus(ResponseStatusOk)
}
