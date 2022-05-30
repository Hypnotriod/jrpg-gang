package controller

import "jrpg-gang/engine"

type UserStatusRequest struct {
	Request
}

func (c *GameController) handleUserStatusRequest(userId engine.UserId, requestRaw string, response *Response) string {
	request := parseRequest(&UserStatusRequest{}, requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMailformed)
	}
	user, ok := c.users.Get(userId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	response.Data[DataKeyUserNickname] = user.Nickname
	response.Data[DataKeyUserId] = user.id
	response.Data[DataKeyUnit] = user.unit
	return response.WithStatus(ResponseStatusOk)
}
