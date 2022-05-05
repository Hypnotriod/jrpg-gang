package controller

import (
	"jrpg-gang/engine"
	"jrpg-gang/util"
)

type StartGameRequest struct {
	Request
}

func parseStartGameRequest(requestRaw string) *StartGameRequest {
	if r, err := util.JsonToObject(&StartGameRequest{}, requestRaw); err == nil {
		return r.(*StartGameRequest)
	}
	return nil
}

func (c *GameController) handleStartGameRequest(requestRaw string, response *Response) string {
	request := parseStartGameRequest(requestRaw)
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
	c.engines.Add(engine)
	return response.WithStatus(ResponseStatusOk)
}
