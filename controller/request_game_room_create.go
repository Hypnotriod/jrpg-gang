package controller

import (
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
)

type CreateGameRoomRequestData struct {
	ScenarioId engine.GameScenarioId `json:"scenarioId"`
}

func (c *GameController) handleCreateGameRoomRequest(playerId engine.PlayerId, request *Request, response *Response) []byte {
	data := parseRequestData(&CreateGameRoomRequestData{}, request.Data)
	if data == nil {
		return response.WithStatus(ResponseStatusMalformed)
	}
	if c.rooms.ExistsForPlayerId(playerId) {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	config, ok := c.scenariosConfig.GetScenarioConfig(data.ScenarioId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	hostUser, _ := c.users.Get(playerId)
	if !hostUser.Unit.CheckRequirements(config.Requirements) {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	c.rooms.Create(
		config.Capacity,
		data.ScenarioId,
		hostUser,
	)
	c.users.ChangeUserStatus(playerId, users.UserStatusInRoom)
	roomInfo := c.rooms.GetRoomInfoByPlayerId(playerId)
	response.Data[DataKeyRoom] = roomInfo
	c.broadcastRoomStatus(roomInfo.Uid)
	return response.WithStatus(ResponseStatusOk)
}
