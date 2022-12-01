package controller

import (
	"jrpg-gang/engine"
)

type GameNextPhaseRequestData struct {
	IsReady bool `json:"isReady"`
}

func (c *GameController) handleGameNextPhaseRequest(userId engine.UserId, request *Request, response *Response) string {
	data := parseRequestData(&GameNextPhaseRequestData{}, request.Data)
	if data == nil {
		return response.WithStatus(ResponseStatusMalformed)
	}
	result, broadcastUserIds, ok := c.engines.ReadyForNextPhase(userId, data.IsReady)
	if !ok {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	response.Data[DataKeyActionResult] = result
	if len(broadcastUserIds) > 0 {
		c.broadcastGameAction(broadcastUserIds, result)
	}
	return response.WithStatus(ResponseStatusOk)
}
