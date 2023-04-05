package controller

import (
	"jrpg-gang/engine"
)

type GameNextPhaseRequestData struct {
	IsReady bool `json:"isReady"`
}

func (c *GameController) handleGameNextPhaseRequest(playerId engine.PlayerId, request *Request, response *Response) []byte {
	data := parseRequestData(&GameNextPhaseRequestData{}, request.Data)
	if data == nil {
		return response.WithStatus(ResponseStatusMalformed)
	}
	wrapper, ok := c.engines.Find(playerId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	defer wrapper.Unlock()
	wrapper.Lock()
	result, broadcastPlayerIds, ok := wrapper.ReadyForNextPhase(playerId, data.IsReady)
	if !ok {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	response.Data[DataKeyActionResult] = result
	if len(broadcastPlayerIds) > 0 {
		c.broadcastGameAction(broadcastPlayerIds, result)
	}
	return response.WithStatus(ResponseStatusOk)
}
