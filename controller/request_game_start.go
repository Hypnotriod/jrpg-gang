package controller

import (
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
	room, present := c.rooms.GetByUserId(request.UserId)
	if !present || room.Host.id != request.UserId {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	// todo: handle scenario
	// scenario := NewTestScenario()
	// actors := room.GetActors()
	// room.engine = engine.NewGameEngine(scenario, actors)
	return response.WithStatus(ResponseStatusOk)
}
