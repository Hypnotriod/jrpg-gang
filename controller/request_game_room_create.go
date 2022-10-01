package controller

import (
	"jrpg-gang/controller/rooms"
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
)

type CreateGameRoomRequestData struct {
	Capacity   uint                     `json:"capacity"`
	ScenarioId rooms.GameRoomScenarioId `json:"scenarioId"`
}

func (c *GameController) handleCreateGameRoomRequest(userId engine.UserId, request *Request, response *Response) string {
	data := parseRequestData(&CreateGameRoomRequestData{}, request.Data)
	if data == nil {
		return response.WithStatus(ResponseStatusMalformed)
	}
	if data.Capacity == 0 || data.Capacity > GAME_ROOM_MAX_CAPACITY {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	if c.rooms.ExistsForUserId(userId) {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	// todo: check available scenario
	hostUser, _ := c.users.Get(userId)
	c.rooms.Create(
		data.Capacity,
		data.ScenarioId,
		hostUser,
	)
	c.users.ChangeUserStatus(userId, users.UserStatusInRoom)
	roomInfo := c.rooms.GetRoomInfoByUserId(userId)
	response.Data[DataKeyRoom] = roomInfo
	c.broadcastRoomStatus(roomInfo.Uid)
	return response.WithStatus(ResponseStatusOk)
}
