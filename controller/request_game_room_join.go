package controller

import (
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
)

type JoinGameRoomRequest struct {
	Request
	Data struct {
		RoomUid uint `json:"roomUid"`
	} `json:"data"`
}

func (c *GameController) handleJoinGameRoomRequest(userId engine.UserId, requestRaw string, response *Response) string {
	request := parseRequest(&JoinGameRoomRequest{}, requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMailformed)
	}
	if c.rooms.ExistsForUserId(userId) {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	if !c.rooms.Has(request.Data.RoomUid) {
		return response.WithStatus(ResponseStatusNotFound)
	}
	user, _ := c.users.Get(userId)
	if !c.rooms.AddUser(request.Data.RoomUid, user) {
		return response.WithStatus(ResponseStatusFailed)
	}
	response.Data[DataKeyRoom] = c.rooms.GetRoomInfoByUid(request.Data.RoomUid)
	c.users.ChangeUserStatus(userId, users.UserStatusInRoom)
	c.broadcastRoomStatus(request.Data.RoomUid)
	return response.WithStatus(ResponseStatusOk)
}
