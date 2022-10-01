package controller

import (
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
)

func (c *GameController) handleLeaveGameRoomRequest(userId engine.UserId, request *Request, response *Response) string {
	roomUid, ok := c.rooms.GetUidByUserId(userId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	if _, ok := c.rooms.RemoveUser(userId); !ok {
		return response.WithStatus(ResponseStatusFailed)
	}
	c.users.ChangeUserStatus(userId, users.UserStatusInLobby)
	c.broadcastRoomStatus(roomUid)
	return response.WithStatus(ResponseStatusOk)
}
