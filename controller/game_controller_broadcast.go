package controller

import (
	"jrpg-gang/controller/chat"
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
)

func (c *GameController) RegisterBroadcaster(broadcaster GameControllerBroadcaster) {
	c.broadcaster = broadcaster
}

func (c *GameController) broadcastGameChatMessage(playerIds []engine.PlayerId, message *chat.ChatMessage) {
	response := NewResponse()
	response.Type = RequestGameChatMessage
	response.Data[DataKeyMessage] = message
	c.broadcaster.BroadcastGameMessageSync(playerIds, response.WithStatus(ResponseStatusOk))
}

func (c *GameController) broadcastLobbyChatMessage(playerIds []engine.PlayerId, message *chat.ChatMessage) {
	response := NewResponse()
	response.Type = RequestLobbyChatMessage
	response.Data[DataKeyMessage] = message
	c.broadcaster.BroadcastGameMessageSync(playerIds, response.WithStatus(ResponseStatusOk))
}

func (c *GameController) broadcastLobbyChatParticipant(playerIds []engine.PlayerId, playerId engine.PlayerId, participant *chat.ChatParticipant) {
	response := NewResponse()
	response.Type = RequestLobbyChatParticipant
	response.Data[DataKeyPlayerId] = playerId
	response.Data[DataKeyParticipant] = participant
	c.broadcaster.BroadcastGameMessageSync(playerIds, response.WithStatus(ResponseStatusOk))
}

func (c *GameController) broadcastGameAction(playerIds []engine.PlayerId, result *engine.GameEvent) {
	response := NewResponse()
	response.Type = RequestGameAction
	response.Data[DataKeyActionResult] = result
	c.broadcaster.BroadcastGameMessageSync(playerIds, response.WithStatus(ResponseStatusOk))
}

func (c *GameController) broadcastServerStatus(playerIds []engine.PlayerId) {
	response := NewResponse()
	response.Type = RequestServerStatus
	response.Data[DataKeyUsersNumber] = c.users.Total()
	for _, playerId := range playerIds {
		if _, ok := c.users.Get(playerId); !ok {
			continue
		}
		c.broadcaster.BroadcastGameMessageAsync([]engine.PlayerId{playerId}, response.WithStatus(ResponseStatusOk))
	}
}

func (c *GameController) broadcastUserStatus(playerIds []engine.PlayerId) {
	for _, playerId := range playerIds {
		user, ok := c.users.Get(playerId)
		if !ok {
			continue
		}
		response := NewResponse()
		response.Type = RequestUserStatus
		response.fillUserStatus(&user)
		c.broadcaster.BroadcastGameMessageAsync([]engine.PlayerId{playerId}, response.WithStatus(ResponseStatusOk))
	}
}

func (c *GameController) broadcastRoomStatus(uid uint) {
	response := NewResponse()
	response.Type = RequestRoomStatus
	response.Data[DataKeyRoom] = c.rooms.GetRoomInfoByUid(uid)
	playerIds := c.users.GetIdsByStatus(users.UserStatusInLobby|users.UserStatusInRoom, true)
	c.broadcaster.BroadcastGameMessageAsync(playerIds, response.WithStatus(ResponseStatusOk))
}

func (c *GameController) BroadcastGameMessageSync(playerIds []engine.PlayerId, message []byte) {
}

func (c *GameController) BroadcastGameMessageAsync(playerIds []engine.PlayerId, message []byte) {
}
