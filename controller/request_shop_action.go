package controller

import (
	"jrpg-gang/domain"
	"jrpg-gang/engine"
)

type ShopActionRequestData struct {
	domain.Action
}

func (c *GameController) handleShopActionRequest(userId engine.UserId, request *Request, response *Response) string {
	data := parseRequestData(&ShopActionRequestData{}, request.Data)
	if data == nil {
		return response.WithStatus(ResponseStatusMalformed)
	}
	user, ok := c.users.Get(userId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	actionResult := c.shop.ExecuteAction(data.Action, &user)
	if actionResult.Result == domain.ResultAccomplished {
		response.Data[DataKeyUnit] = user.Unit
	}
	response.Data[DataKeyAction] = request.Data
	response.Data[DataKeyActionResult] = actionResult
	return response.WithStatus(ResponseStatusOk)
}
