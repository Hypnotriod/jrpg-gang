package controller

import (
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
)

type CreateGameRoomRequestData struct {
	Capacity   uint                  `json:"capacity"`
	ScenarioId engine.GameScenarioId `json:"scenarioId"`
}

func (c *GameController) handleCreateGameRoomRequest(playerId engine.PlayerId, request *Request, response *Response) string {
	data := parseRequestData(&CreateGameRoomRequestData{}, request.Data)
	if data == nil {
		return response.WithStatus(ResponseStatusMalformed)
	}
	if data.Capacity == 0 || data.Capacity > GAME_ROOM_MAX_CAPACITY {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	if c.rooms.ExistsForPlayerId(playerId) {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	if !c.scenarioConfig.Has(data.ScenarioId) {
		return response.WithStatus(ResponseStatusNotFound)
	}
	hostUser, _ := c.users.Get(playerId)
	c.rooms.Create(
		data.Capacity,
		data.ScenarioId,
		hostUser,
	)
	c.users.ChangeUserStatus(playerId, users.UserStatusInRoom)
	roomInfo := c.rooms.GetRoomInfoByPlayerId(playerId)
	response.Data[DataKeyRoom] = roomInfo
	c.broadcastRoomStatus(roomInfo.Uid)
	return response.WithStatus(ResponseStatusOk)
}
