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
	host, _ := c.users.Get(playerId)
	mercenaries := room.GetMercenaryCodes()
	for _, code := range mercenaries {
		c.mercenaries.Refund(code, &host.Unit.Unit)
	}
	c.broadcastRoomStatus(room.Uid)
	c.broadcastUserStatus(playerIds)
	return response.WithStatus(ResponseStatusOk)
}
