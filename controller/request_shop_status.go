package controller

import "jrpg-gang/engine"

type ShopStatusRequest struct {
	Request
}

func (c *GameController) handleShopStatusRequest(userId engine.UserId, requestRaw string, response *Response) string {
	request := parseRequest(&ShopStatusRequest{}, requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMalformed)
	}
	response.Data[DataKeyShop] = c.shop.GetStatus()
	return response.WithStatus(ResponseStatusOk)
}
