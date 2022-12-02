package controller

import "jrpg-gang/engine"

func (c *GameController) handlePlayerInfoRequest(userId engine.UserId, request *Request, response *Response) string {
	wrapper, ok := c.engines.Find(userId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	defer wrapper.RUnlock()
	wrapper.RLock()
	result, ok := wrapper.ReadPlayerInfo(userId)
	if !ok {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	response.Data[DataKeyPlayerInfo] = result
	return response.WithStatus(ResponseStatusOk)
}
