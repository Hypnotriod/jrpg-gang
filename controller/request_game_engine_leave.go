package controller

import (
	"jrpg-gang/engine"
)

func (c *GameController) handleGameLeaveRequest(userId engine.UserId, request *Request, response *Response) string {
	result, broadcastUserIds, unit, ok := c.engines.LeaveGame(userId)
	if !ok {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	c.users.ResetUser(userId)
	c.users.UpdateWithUnitOnGameComplete(userId, &unit)
	if len(broadcastUserIds) > 0 {
		c.broadcastGameState(broadcastUserIds, result)
	}
	return response.WithStatus(ResponseStatusOk)
}
