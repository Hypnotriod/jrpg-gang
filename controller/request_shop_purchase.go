package controller

import "jrpg-gang/engine"

type ShopPurchaseRequest struct {
	Request
	Data struct {
		ItemUid uint `json:"itemUid"`
	} `json:"data"`
}

func (c *GameController) handleShopPurchaseRequest(userId engine.UserId, requestRaw string, response *Response) string {
	request := parseRequest(&ShopPurchaseRequest{}, requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMailformed)
	}
	user, ok := c.users.Get(userId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	if !c.shop.Buy(user, request.Data.ItemUid) {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	response.Data[DataKeyUnit] = user.unit
	return response.WithStatus(ResponseStatusOk)
}
