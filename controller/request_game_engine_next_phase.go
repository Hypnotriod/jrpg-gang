package controller

import "jrpg-gang/engine"

type GameNextPhaseRequest struct {
	Request
}

func (c *GameController) handleGameNextPhaseRequest(userId engine.UserId, requestRaw string, response *Response) string {
	request := parseRequest(&GameNextPhaseRequest{}, requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMalformed)
	}
	result, broadcastUserIds, ok := c.engines.NextPhase(userId)
	if !ok {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	response.Data[DataKeyActionResult] = result
	if len(broadcastUserIds) > 0 {
		c.broadcastGameAction(broadcastUserIds, result)
	}
	return response.WithStatus(ResponseStatusOk)
}
