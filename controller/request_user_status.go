package controller

import "jrpg-gang/engine"

func (c *GameController) handleUserStatusRequest(userId engine.UserId, request *Request, response *Response) string {
	user, ok := c.users.Get(userId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	response.fillUserStatus(&user)
	return response.WithStatus(ResponseStatusOk)
}
