package controller

import "jrpg-gang/util"

type CreateGameRoomRoomRequest struct {
	Request
	Data struct {
		Capacity uint `json:"capacity"`
	} `json:"data"`
}

func parseCreateGameRoomRequest(requestRaw string) *CreateGameRoomRoomRequest {
	if r, ok := util.JsonToObject(&CreateGameRoomRoomRequest{}, requestRaw); ok {
		return r.(*CreateGameRoomRoomRequest)
	}
	return nil
}

func (c *GameController) handleCreateGameRoomRequest(requestRaw string, response *Response) string {
	request := parseCreateGameRoomRequest(requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMailformed)
	}
	if c.rooms.ExistsForUserId(request.UserId) {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	if request.Data.Capacity == 0 || request.Data.Capacity > GAME_ROOM_MAX_CAPACITY {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	room := &GameRoom{
		HostId:   request.UserId,
		Capacity: request.Data.Capacity,
	}
	c.rooms.Add(room)
	return response.WithStatus(ResponseStatusOk)
}
