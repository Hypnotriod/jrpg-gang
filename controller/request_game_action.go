package controller

import (
	"jrpg-gang/domain"
	"jrpg-gang/engine"
)

type GameActionRequest struct {
	Request
	Data domain.Action `json:"data"`
}

func (c *GameController) handleGameActionRequest(userId engine.UserId, requestRaw string, response *Response) string {
	request := parseRequest(&GameActionRequest{}, requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMailformed)
	}
	result, ok := c.engines.ExecuteUserAction(request.Data, userId)
	if !ok {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	response.Data[DataKeyActionResult] = result
	return response.WithStatus(ResponseStatusOk)
}
