package controller

import "jrpg-gang/engine"

func (c *GameController) handleShopStatusRequest(playerId engine.PlayerId, request *Request, response *Response) string {
	response.Data[DataKeyShop] = c.shop.GetStatus()
	return response.WithStatus(ResponseStatusOk)
}
