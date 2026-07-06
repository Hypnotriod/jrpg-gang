package controller

import "jrpg-gang/engine"

func (c *GameController) handleMercenariesStatusRequest(playerId engine.PlayerId, request *Request, response *Response) []byte {
	response.Data[DataKeyMercenaries] = c.mercenaries.GetStatus()
	return response.WithStatus(ResponseStatusOk)
}
