package controller

import (
	"jrpg-gang/engine"
)

type StartGameRequest struct {
	Request
}

func (c *GameController) handleStartGameRequest(requestRaw string, response *Response) string {
	request := parseRequest(&StartGameRequest{}, requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMailformed)
	}
	room, ok := c.rooms.PopByHostId(request.UserId)
	if !ok {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	// todo: handle scenario
	scenario := NewTestScenario()
	actors := room.GetActors()
	engine := engine.NewGameEngine(scenario, actors)
	state := engine.NewGameEvent()
	userIds := engine.GetUserIds()
	c.broadcastGameState(userIds, state)
	c.engines.Add(request.UserId, engine)

	return response.WithStatus(ResponseStatusOk)
}
