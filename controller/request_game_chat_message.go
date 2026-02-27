package controller

import (
	"jrpg-gang/engine"
)

type ChatMessageRequestData struct {
	Message string `json:"message"`
}

func (c *GameController) handleGameChatMessageRequest(playerId engine.PlayerId, request *Request, response *Response) []byte {
	data := parseRequestData(&ChatMessageRequestData{}, request.Data)
	if data == nil {
		return response.WithStatus(ResponseStatusMalformed)
	}
	wrapper, ok := c.engines.Find(playerId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	msg := wrapper.SendChatMessage(playerId, data.Message)
	response.Data[DataKeyMessage] = msg
	return response.WithStatus(ResponseStatusOk)
}
