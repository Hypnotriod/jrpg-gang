package controller

import (
	"jrpg-gang/controller/factory"
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
)

type StartGameRequest struct {
	Request
}

func (c *GameController) handleStartGameRequest(userId engine.UserId, requestRaw string, response *Response) string {
	request := parseRequest(&StartGameRequest{}, requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMailformed)
	}
	room, ok := c.rooms.PopByHostId(userId)
	if !ok {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	// todo: handle scenario
	scenario := factory.NewTestScenario()
	actors := room.GetActors()
	engine := engine.NewGameEngine(scenario, actors)
	state := engine.NewGameEvent()
	userIds := engine.GetUserIds()
	for _, userId := range userIds {
		c.users.ChangeUserStatus(userId, users.UserStatusInGame)
	}
	c.broadcastGameState(userIds, state)
	c.broadcastRoomStatus(room.Uid)
	c.broadcastUsersStatus(userIds)
	c.engines.Add(userId, engine)
	return response.WithStatus(ResponseStatusOk)
}
