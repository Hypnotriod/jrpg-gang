package controller

import (
	"jrpg-gang/domain"
)

type GameActionRequest struct {
	Request
	Data domain.Action `json:"data"`
}

func (c *GameController) handleGameActionRequest(requestRaw string, response *Response) string {
	request := parseRequest(&GameActionRequest{}, requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMailformed)
	}
	result, broadcastUserIds, ok := c.engines.ExecuteUserAction(request.Data, request.UserId)
	if !ok {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	response.Data[DataKeyActionResult] = result
	if len(broadcastUserIds) > 0 {
		c.broadcastGameAction(broadcastUserIds, result)
	}
	return response.WithStatus(ResponseStatusOk)
}
