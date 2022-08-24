package controller

import "jrpg-gang/engine"

type CreateGameRoomRequest struct {
	Request
	Data struct {
		Capacity   uint               `json:"capacity"`
		ScenarioId GameRoomScenarioId `json:"scenarioId"`
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
	room := NewGameRoom()
	room.Capacity = request.Data.Capacity
	room.ScenarioId = request.Data.ScenarioId
	room.host = hostUser
	c.rooms.Add(room)
	response.Data[DataKeyRoom] = room
	c.users.ChangeUserStatus(userId, UserStatusInRoom)
	c.broadcastLobbyStatus()
	return response.WithStatus(ResponseStatusOk)
}
