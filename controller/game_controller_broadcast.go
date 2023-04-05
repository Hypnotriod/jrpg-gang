package controller

import (
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
)

func (c *GameController) RegisterBroadcaster(broadcaster GameControllerBroadcaster) {
	c.broadcaster = broadcaster
}

func (c *GameController) broadcastGameAction(playerIds []engine.PlayerId, result *engine.GameEvent) {
	response := NewResponse()
	response.Type = RequestGameAction
	response.Data[DataKeyActionResult] = result
	c.broadcaster.BroadcastGameMessageSync(playerIds, response.WithStatus(ResponseStatusOk))
}

func (c *GameController) broadcastUsersStatus(playerIds []engine.PlayerId) {
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
	response.Data[DataKeyUsersCount] = c.users.TotalCount()
	playerIds := c.users.GetIdsByStatus(users.UserStatusInLobby|users.UserStatusInRoom, true)
	c.broadcaster.BroadcastGameMessageAsync(playerIds, response.WithStatus(ResponseStatusOk))
}

func (c *GameController) BroadcastGameMessageSync(playerIds []engine.PlayerId, message []byte) {
}

func (c *GameController) BroadcastGameMessageAsync(playerIds []engine.PlayerId, message []byte) {
}
