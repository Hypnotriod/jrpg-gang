package controller

import "jrpg-gang/engine"

type GameStateRequest struct {
	Request
}

func (c *GameController) handleGameStateRequest(userId engine.UserId, requestRaw string, response *Response) string {
	request := parseRequest(&GameStateRequest{}, requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMailformed)
	}
	result, _, ok := c.engines.GameState(userId)
	if !ok {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	response.Data[DataKeyGameState] = result
	return response.WithStatus(ResponseStatusOk)
}
