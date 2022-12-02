package controller

import (
	"jrpg-gang/engine"
)

func (c *GameController) handleGameLeaveRequest(userId engine.UserId, request *Request, response *Response) string {
	wrapper, ok := c.engines.Unregister(userId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	defer wrapper.Unlock()
	wrapper.Lock()
	result, broadcastUserIds, unit, ok := wrapper.LeaveGame(userId)
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
