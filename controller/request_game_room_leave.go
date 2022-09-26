package controller

import (
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
)

type LeaveGameRoomRequest struct {
	Request
}

func (c *GameController) handleLeaveGameRoomRequest(userId engine.UserId, requestRaw string, response *Response) string {
	request := parseRequest(&LeaveGameRoomRequest{}, requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMailformed)
	}
	roomUid, ok := c.rooms.GetUidByUserId(userId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	if _, ok := c.rooms.RemoveUser(userId); !ok {
		return response.WithStatus(ResponseStatusFailed)
	}
	c.users.ChangeUserStatus(userId, users.UserStatusJoined)
	c.broadcastRoomStatus(roomUid)
	return response.WithStatus(ResponseStatusOk)
}
