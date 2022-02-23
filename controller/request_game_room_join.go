package controller

import "jrpg-gang/util"

type JoinGameRoomRequest struct {
	Request
	Data struct {
		RoomUid uint `json:"roomUid"`
	} `json:"data"`
}

func parseJoinGameRoomRequest(requestRaw string) *JoinGameRoomRequest {
	if r, err := util.JsonToObject(&JoinGameRoomRequest{}, requestRaw); err != nil {
		return r.(*JoinGameRoomRequest)
	}
	return nil
}

func (c *GameController) handleJoinGameRoomRequest(requestRaw string, response *Response) string {
	request := parseJoinGameRoomRequest(requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMailformed)
	}
	if c.rooms.ExistsForUserId(request.UserId) {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	if !c.rooms.Has(request.Data.RoomUid) {
		return response.WithStatus(ResponseStatusNotFound)
	}
	user, _ := c.users.Get(request.UserId)
	if !c.rooms.AddUser(request.Data.RoomUid, user) {
		return response.WithStatus(ResponseStatusFailed)
	}
	response.Data[DataKeyRoom], _ = c.rooms.GetByUserId(user.id)
	return response.WithStatus(ResponseStatusOk)
}
