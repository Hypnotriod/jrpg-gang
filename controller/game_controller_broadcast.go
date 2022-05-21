package controller

import "jrpg-gang/engine"

func (c *GameController) RegisterBroadcaster(broadcaster GameControllerBroadcaster) {
	c.broadcaster = broadcaster
}
func (c *GameController) broadcastGameAction(userIds []engine.UserId, result *engine.GameEvent) {
	response := NewResponse()
	response.Type = RequestGameAction
	response.Data[DataKeyActionResult] = result
	c.broadcaster.BroadcastGameMessage(userIds, response.WithStatus(ResponseStatusOk))
}

func (c *GameController) broadcastGameState(userIds []engine.UserId, state *engine.GameEvent) {
	response := NewResponse()
	response.Type = RequestGameState
	response.Data[DataKeyGameState] = state
	c.broadcaster.BroadcastGameMessage(userIds, response.WithStatus(ResponseStatusOk))
}

func (c *GameController) broadcastLobbyStatus() {
	response := NewResponse()
	response.Type = RequestLobbyStatus
	response.Data[DataKeyRooms] = c.rooms.ResponseList()
	response.Data[DataKeyUsersCount] = c.users.TotalCount()
	userIds := c.users.GetIdsByStatus(UserStatusJoined | UserStatusInRoom)
	c.broadcaster.BroadcastGameMessage(userIds, response.WithStatus(ResponseStatusOk))
}

func (c *GameController) BroadcastGameMessage(userIds []engine.UserId, message string) {
}
