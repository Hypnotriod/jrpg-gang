package controller

import (
	"jrpg-gang/engine"
)

func (c *GameController) handleGameLeaveRequest(playerId engine.PlayerId, request *Request, response *Response) []byte {
	wrapper, ok := c.engines.Unregister(playerId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	wrapper.Lock()
	defer wrapper.Unlock()
	result, broadcastPlayerIds, unit, ok := wrapper.LeaveGame(playerId)
	if !ok {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	c.users.ResetUser(playerId)
	c.users.UpdateWithUnitOnGameComplete(playerId, &unit)
	if user, ok := c.users.Get(playerId); ok {
		c.persistUser(&user)
	}
	if len(broadcastPlayerIds) > 0 {
		c.broadcastGameAction(broadcastPlayerIds, result)
	}
	return response.WithStatus(ResponseStatusOk)
}
