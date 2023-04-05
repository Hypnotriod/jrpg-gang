package controller

import (
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
)

func (c *GameController) handleDestroyGameRoomRequest(playerId engine.PlayerId, request *Request, response *Response) []byte {
	if !c.rooms.ExistsForHostId(playerId) {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	room, ok := c.rooms.PopByHostId(playerId)
	if !ok {
		return response.WithStatus(ResponseStatusFailed)
	}
	playerIds := room.GetPlayerIds()
	for _, playerId := range playerIds {
		c.users.ChangeUserStatus(playerId, users.UserStatusInLobby)
	}
	c.broadcastRoomStatus(room.Uid)
	c.broadcastUsersStatus(playerIds)
	return response.WithStatus(ResponseStatusOk)
}
