package controller

import "jrpg-gang/util"

type LeaveGameRoomRequest struct {
	Request
}

func parseLeaveGameRoomRequest(requestRaw string) *LeaveGameRoomRequest {
	if r, ok := util.JsonToObject(&LeaveGameRoomRequest{}, requestRaw); ok {
		return r.(*LeaveGameRoomRequest)
	}
	return nil
}

func (c *GameController) handleLeaveGameRoomRequest(requestRaw string, response *Response) string {
	request := parseLeaveGameRoomRequest(requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMailformed)
	}
	if !c.rooms.ExistsForUserId(request.UserId) {
		return response.WithStatus(ResponseStatusNotFound)
	}
	if !c.rooms.RemoveUser(request.UserId) {
		return response.WithStatus(ResponseStatusFailed)
	}
	return response.WithStatus(ResponseStatusOk)
}
