package controller

import (
	"jrpg-gang/engine"
)

type GameLeaveRequest struct {
	Request
}

func (c *GameController) handleGameLeaveRequest(userId engine.UserId, requestRaw string, response *Response) string {
	request := parseRequest(&GameLeaveRequest{}, requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMalformed)
	}
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
