package controller

import "jrpg-gang/engine"

type DestroyGameRoomRequest struct {
	Request
}

func (c *GameController) handleDestroyGameRoomRequest(userId engine.UserId, requestRaw string, response *Response) string {
	request := parseRequest(&DestroyGameRoomRequest{}, requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMailformed)
	}
	if !c.rooms.ExistsForHostId(userId) {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	room, ok := c.rooms.PopByHostId(userId)
	if !ok {
		return response.WithStatus(ResponseStatusFailed)
	}
	for _, userId := range room.GetUserIds() {
		c.users.ChangeUserStatus(userId, UserStatusJoined)
	}
	c.broadcastLobbyStatus()
	return response.WithStatus(ResponseStatusOk)
}
