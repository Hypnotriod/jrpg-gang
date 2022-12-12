package controller

import (
	"jrpg-gang/controller/gameengines"
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
	wrapper := gameengines.NewGameEngineWrapper(engine.NewGameEngine(scenario, actors), c.broadcastGameAction)
	defer wrapper.Unlock()
	wrapper.Lock()
	state, userIds := c.engines.Register(wrapper)
	for _, userId := range userIds {
		c.users.ChangeUserStatus(userId, users.UserStatusInGame)
	}
	c.broadcastGameAction(userIds, state)
	c.broadcastRoomStatus(room.Uid)
	c.broadcastUsersStatus(userIds)
	return response.WithStatus(ResponseStatusOk)
}
