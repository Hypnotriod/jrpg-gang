package controller

import "jrpg-gang/engine"

func (c *GameController) handleShopStatusRequest(playerId engine.PlayerId, request *Request, response *Response) []byte {
	user, ok := c.users.Get(playerId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	response.Data[DataKeyShop] = c.shop.GetStatus(&user.Unit.Unit)
	return response.WithStatus(ResponseStatusOk)
}
