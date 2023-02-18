package controller

import (
	"jrpg-gang/domain"
	"jrpg-gang/engine"
)

type GameActionRequestData struct {
	domain.Action
}

func (c *GameController) handleGameActionRequest(playerId engine.PlayerId, request *Request, response *Response) string {
	data := parseRequestData(&GameActionRequestData{}, request.Data)
	if data == nil {
		return response.WithStatus(ResponseStatusMalformed)
	}
	wrapper, ok := c.engines.Find(playerId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	defer wrapper.Unlock()
	wrapper.Lock()
	result, broadcastPlayerIds, ok := wrapper.ExecuteUserAction(data.Action, playerId)
	if !ok {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	response.Data[DataKeyAction] = request.Data
	response.Data[DataKeyActionResult] = result
	if len(broadcastPlayerIds) > 0 {
		c.broadcastGameAction(broadcastPlayerIds, result)
	}
	return response.WithStatus(ResponseStatusOk)
}
