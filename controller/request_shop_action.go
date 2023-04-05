package controller

import (
	"jrpg-gang/domain"
	"jrpg-gang/engine"
)

type ShopActionRequestData struct {
	domain.Action
}

func (c *GameController) handleShopActionRequest(playerId engine.PlayerId, request *Request, response *Response) []byte {
	data := parseRequestData(&ShopActionRequestData{}, request.Data)
	if data == nil {
		return response.WithStatus(ResponseStatusMalformed)
	}
	user, ok := c.users.Get(playerId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	actionResult := c.shop.ExecuteAction(data.Action, &user)
	if actionResult.Result == domain.ResultAccomplished {
		response.Data[DataKeyUnit] = user.Unit
	}
	response.Data[DataKeyAction] = data.Action
	response.Data[DataKeyActionResult] = actionResult
	return response.WithStatus(ResponseStatusOk)
}
