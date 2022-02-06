package controller

import "jrpg-gang/util"

type DestroyGameRoomRequest struct {
	Request
}

func parseDestroyGameRoomRequest(requestRaw string) *DestroyGameRoomRequest {
	if r, ok := util.JsonToObject(&DestroyGameRoomRequest{}, requestRaw); ok {
		return r.(*DestroyGameRoomRequest)
	}
	return nil
}

func (c *GameController) handleDestroyGameRoomRequest(requestRaw string, response *Response) string {
	request := parseDestroyGameRoomRequest(requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMailformed)
	}
	if !c.rooms.ExistsForHostId(request.UserId) {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	if !c.rooms.RemoveByHostId(request.UserId) {
		return response.WithStatus(ResponseStatusFailed)
	}
	return response.WithStatus(ResponseStatusOk)
}
