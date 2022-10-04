package controller

import (
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
)

func (c *GameController) handleStartGameRequest(userId engine.UserId, request *Request, response *Response) string {
	room, ok := c.rooms.PopByHostId(userId)
	if !ok {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	scenario := c.scenarioConfig.Get(room.ScenarioId)
	if scenario == nil {
		return response.WithStatus(ResponseStatusNotFound)
	}
	actors := room.GetActors()
	engine := engine.NewGameEngine(scenario, actors)
	state := engine.NewGameEventWithPlayerInfo()
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
