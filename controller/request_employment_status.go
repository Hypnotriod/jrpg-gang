package controller

import "jrpg-gang/engine"

func (c *GameController) handleJobStatusRequest(playerId engine.PlayerId, request *Request, response *Response) []byte {
	user, ok := c.users.Get(playerId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	response.Data[DataKeyEmployment] = c.employment.GetStatus(&user)
	return response.WithStatus(ResponseStatusOk)
}
