package controller

import (
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
)

func (c *GameController) handleLeaveGameRoomRequest(playerId engine.PlayerId, request *Request, response *Response) string {
	roomUid, ok := c.rooms.GetUidByPlayerId(playerId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	if _, ok := c.rooms.RemoveUser(playerId); !ok {
		return response.WithStatus(ResponseStatusFailed)
	}
	c.users.ChangeUserStatus(playerId, users.UserStatusInLobby)
	c.broadcastRoomStatus(roomUid)
	return response.WithStatus(ResponseStatusOk)
}
