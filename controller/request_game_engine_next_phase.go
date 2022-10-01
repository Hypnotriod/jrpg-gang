package controller

import "jrpg-gang/engine"

func (c *GameController) handleGameNextPhaseRequest(userId engine.UserId, request *Request, response *Response) string {
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
