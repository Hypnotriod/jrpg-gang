package controller

import (
	"jrpg-gang/domain"
	"jrpg-gang/engine"
)

type ConfiguratorActionRequestData struct {
	domain.Action
}

func (c *GameController) handleConfiguratorActionRequest(playerId engine.PlayerId, request *Request, response *Response) []byte {
	data := parseRequestData(&ConfiguratorActionRequestData{}, request.Data)
	if data == nil {
		return response.WithStatus(ResponseStatusMalformed)
	}
	user, ok := c.users.Get(playerId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	actionResult := c.configurator.ExecuteAction(data.Action, &user.Unit.Unit)
	if actionResult.Result == domain.ResultAccomplished {
		response.Data[DataKeyUnit] = user.Unit
	}
	if user.Level != user.Unit.Stats.Progress.Level {
		c.users.UpdateOnLevelUp(playerId, &user.Unit.Unit)
	}
	response.Data[DataKeyAction] = request.Data
	response.Data[DataKeyActionResult] = actionResult
	return response.WithStatus(ResponseStatusOk)
}
