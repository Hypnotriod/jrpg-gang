package controller

import (
	"jrpg-gang/domain"
	"jrpg-gang/engine"
)

type ConfiguratorActionRequest struct {
	Request
	Data domain.Action `json:"data"`
}

func (c *GameController) handleConfiguratorActionRequest(userId engine.UserId, requestRaw string, response *Response) string {
	request := parseRequest(&ConfiguratorActionRequest{}, requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMailformed)
	}
	user, ok := c.users.Get(userId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	actionResult := c.configurator.ExecuteAction(request.Data, &user.Unit.Unit)
	if actionResult.Result == domain.ResultAccomplished {
		c.users.UpdateUnit(user.Id, &user.Unit)
		response.Data[DataKeyUnit] = user.Unit
	}
	response.Data[DataKeyAction] = request.Data
	response.Data[DataKeyActionResult] = actionResult
	return response.WithStatus(ResponseStatusOk)
}
