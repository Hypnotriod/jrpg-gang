package controller

import (
	"jrpg-gang/controller/rooms"
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
)

type CreateGameRoomRequest struct {
	Request
	Data struct {
		Capacity   uint                     `json:"capacity"`
		ScenarioId rooms.GameRoomScenarioId `json:"scenarioId"`
	} `json:"data"`
}

func (c *GameController) handleCreateGameRoomRequest(userId engine.UserId, requestRaw string, response *Response) string {
	request := parseRequest(&CreateGameRoomRequest{}, requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMailformed)
	}
	if request.Data.Capacity == 0 || request.Data.Capacity > GAME_ROOM_MAX_CAPACITY {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	if c.rooms.ExistsForUserId(userId) {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	// todo: check available scenario
	hostUser, _ := c.users.Get(userId)
	c.rooms.Create(
		request.Data.Capacity,
		request.Data.ScenarioId,
		hostUser,
	)
	c.users.ChangeUserStatus(userId, users.UserStatusInRoom)
	roomInfo := c.rooms.GetRoomInfoByUserId(userId)
	response.Data[DataKeyRoom] = roomInfo
	c.broadcastRoomStatus(roomInfo.Uid)
	return response.WithStatus(ResponseStatusOk)
}
