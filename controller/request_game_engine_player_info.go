package controller

import "jrpg-gang/engine"

type PlayerInfoRequest struct {
	Request
}

func (c *GameController) handlePlayerInfoRequest(userId engine.UserId, requestRaw string, response *Response) string {
	request := parseRequest(&PlayerInfoRequest{}, requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMailformed)
	}
	result, ok := c.engines.PlayerInfo(userId)
	if !ok {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	response.Data[DataKeyPlayerInfo] = result
	return response.WithStatus(ResponseStatusOk)
}
