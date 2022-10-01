package controller

import (
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
)

type JoinGameRoomRequestData struct {
	RoomUid uint `json:"roomUid"`
}

func (c *GameController) handleJoinGameRoomRequest(userId engine.UserId, request *Request, response *Response) string {
	data := parseRequestData(&JoinGameRoomRequestData{}, request.Data)
	if data == nil {
		return response.WithStatus(ResponseStatusMalformed)
	}
	if c.rooms.ExistsForUserId(userId) {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	if !c.rooms.Has(data.RoomUid) {
		return response.WithStatus(ResponseStatusNotFound)
	}
	user, _ := c.users.Get(userId)
	if !c.rooms.AddUser(data.RoomUid, user) {
		return response.WithStatus(ResponseStatusFailed)
	}
	response.Data[DataKeyRoom] = c.rooms.GetRoomInfoByUid(data.RoomUid)
	c.users.ChangeUserStatus(userId, users.UserStatusInRoom)
	c.broadcastRoomStatus(data.RoomUid)
	return response.WithStatus(ResponseStatusOk)
}
