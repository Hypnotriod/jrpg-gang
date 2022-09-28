package controller

import (
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
)

func (c *GameController) RegisterBroadcaster(broadcaster GameControllerBroadcaster) {
	c.broadcaster = broadcaster
}

func (c *GameController) broadcastGameAction(userIds []engine.UserId, result *engine.GameEvent) {
	response := NewResponse()
	response.Type = RequestGameAction
	response.Data[DataKeyActionResult] = result
	c.broadcaster.BroadcastGameMessageSync(userIds, response.WithStatus(ResponseStatusOk))
}

func (c *GameController) broadcastGameState(userIds []engine.UserId, state *engine.GameEvent) {
	response := NewResponse()
	response.Type = RequestGameState
	response.Data[DataKeyGameState] = state
	c.broadcaster.BroadcastGameMessageSync(userIds, response.WithStatus(ResponseStatusOk))
}

func (c *GameController) broadcastUsersStatus(userIds []engine.UserId) {
	for _, userId := range userIds {
		user, ok := c.users.Get(userId)
		if !ok {
			continue
		}
		response := NewResponse()
		response.Type = RequestUserStatus
		response.fillUserStatus(&user)
		c.broadcaster.BroadcastGameMessageAsync([]engine.UserId{userId}, response.WithStatus(ResponseStatusOk))
	}
}

func (c *GameController) broadcastRoomStatus(uid uint) {
	response := NewResponse()
	response.Type = RequestRoomStatus
	response.Data[DataKeyRoom] = c.rooms.GetRoomInfoByUid(uid)
	response.Data[DataKeyUsersCount] = c.users.TotalCount()
	userIds := c.users.GetIdsByStatus(users.UserStatusInLobby|users.UserStatusInRoom, true)
	c.broadcaster.BroadcastGameMessageAsync(userIds, response.WithStatus(ResponseStatusOk))
}

func (c *GameController) BroadcastGameMessageSync(userIds []engine.UserId, message string) {
}

func (c *GameController) BroadcastGameMessageAsync(userIds []engine.UserId, message string) {
}
