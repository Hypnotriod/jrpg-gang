package controller

import "jrpg-gang/engine"

func (c *GameController) handlePlayerInfoRequest(userId engine.UserId, request *Request, response *Response) string {
	result, ok := c.engines.PlayerInfo(userId)
	if !ok {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	response.Data[DataKeyPlayerInfo] = result
	return response.WithStatus(ResponseStatusOk)
}
