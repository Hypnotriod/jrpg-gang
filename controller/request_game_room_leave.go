package controller

import "jrpg-gang/engine"

type LeaveGameRoomRequest struct {
	Request
}

func (c *GameController) handleLeaveGameRoomRequest(userId engine.UserId, requestRaw string, response *Response) string {
	request := parseRequest(&LeaveGameRoomRequest{}, requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMailformed)
	}
	if !c.rooms.ExistsForUserId(userId) {
		return response.WithStatus(ResponseStatusNotFound)
	}
	if !c.rooms.RemoveUser(userId) {
		return response.WithStatus(ResponseStatusFailed)
	}
	c.users.ChangeUserStatus(userId, UserStatusJoined)
	c.broadcastLobbyStatus()
	return response.WithStatus(ResponseStatusOk)
}
