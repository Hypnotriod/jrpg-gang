package controller

import (
	"jrpg-gang/domain"
	"jrpg-gang/engine"
)

type QuestsActionRequestData struct {
	domain.Action
}

func (c *GameController) handleQuestsActionRequest(playerId engine.PlayerId, request *Request, response *Response) []byte {
	data := parseRequestData(&QuestsActionRequestData{}, request.Data)
	if data == nil {
		return response.WithStatus(ResponseStatusMalformed)
	}
	user, ok := c.users.Get(playerId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	actionResult := c.quests.ExecuteAction(data.Action, &user)
	if actionResult.Result == domain.ResultAccomplished {
		c.persistUser(&user)
		response.Data[DataKeyUnit] = user.Unit
	}
	response.Data[DataKeyAction] = data.Action
	response.Data[DataKeyActionResult] = actionResult
	return response.WithStatus(ResponseStatusOk)
}
