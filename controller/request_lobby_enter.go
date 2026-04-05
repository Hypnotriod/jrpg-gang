package controller

import (
	"jrpg-gang/controller/chat"
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
)

func (c *GameController) handleEnterLobbyRequest(playerId engine.PlayerId, request *Request, response *Response) []byte {
	user, ok := c.users.Get(playerId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	status := user.Status
	if status == users.UserStatusInLobby {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	c.users.ChangeUserStatus(playerId, users.UserStatusInLobby)
	c.lobbyChat.AddParticipant(playerId, &chat.ChatParticipant{
		Nickname: user.Nickname,
	})
	return response.WithStatus(ResponseStatusOk)
}
