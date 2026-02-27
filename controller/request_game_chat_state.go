package controller

import (
	"jrpg-gang/engine"
)

func (c *GameController) handleGameChatStateRequest(playerId engine.PlayerId, request *Request, response *Response) []byte {
	wrapper, ok := c.engines.Find(playerId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	response.Data[DataKeyChat] = wrapper.ChatState()
	return response.WithStatus(ResponseStatusOk)
}
