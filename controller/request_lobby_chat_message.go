package controller

import (
	"jrpg-gang/controller/chat"
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
)

func (c *GameController) handleLobbyChatMessageRequest(playerId engine.PlayerId, request *Request, response *Response) []byte {
	data := parseRequestData(&ChatMessageRequestData{}, request.Data)
	if data == nil {
		return response.WithStatus(ResponseStatusMalformed)
	}
	status := c.users.GetUserStatus(playerId)
	if !status.Test(users.UserStatusInLobby | users.UserStatusInRoom) {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	msg, err := c.lobbyChat.SendMessage(playerId, data.Message)
	if err != nil {
		if err == chat.ErrMessagerateLimit {
			return response.WithStatus(ResponseStatusNotAllowed)
		}
		return response.WithStatus(ResponseStatusMalformed)
	}
	response.Data[DataKeyMessage] = msg
	return response.WithStatus(ResponseStatusOk)
}
