package controller

import (
	"jrpg-gang/engine"
)

func (c *GameController) handleGameLeaveRequest(userId engine.UserId, request *Request, response *Response) string {
	result, broadcastUserIds, unit, unlock, ok := c.engines.LeaveGame(userId)
	if unlock != nil {
		defer unlock()
	}
	if !ok {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	c.users.ResetUser(userId)
	c.users.UpdateWithUnitOnGameComplete(userId, &unit)
	if len(broadcastUserIds) > 0 {
		c.broadcastGameAction(broadcastUserIds, result)
	}
	return response.WithStatus(ResponseStatusOk)
}
