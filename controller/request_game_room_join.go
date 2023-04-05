package controller

import (
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
)

type JoinGameRoomRequestData struct {
	RoomUid uint `json:"roomUid"`
}

func (c *GameController) handleJoinGameRoomRequest(playerId engine.PlayerId, request *Request, response *Response) []byte {
	data := parseRequestData(&JoinGameRoomRequestData{}, request.Data)
	if data == nil {
		return response.WithStatus(ResponseStatusMalformed)
	}
	if c.rooms.ExistsForPlayerId(playerId) {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	if !c.rooms.Has(data.RoomUid) {
		return response.WithStatus(ResponseStatusNotFound)
	}
	user, _ := c.users.Get(playerId)
	if !c.rooms.AddUser(data.RoomUid, user) {
		return response.WithStatus(ResponseStatusFailed)
	}
	response.Data[DataKeyRoom] = c.rooms.GetRoomInfoByUid(data.RoomUid)
	c.users.ChangeUserStatus(playerId, users.UserStatusInRoom)
	c.broadcastRoomStatus(data.RoomUid)
	return response.WithStatus(ResponseStatusOk)
}
