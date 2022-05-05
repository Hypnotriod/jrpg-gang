package controller

import "jrpg-gang/util"

type UserStatusRequest struct {
	Request
}

func parseUserStatusRequest(requestRaw string) *UserStatusRequest {
	if r, err := util.JsonToObject(&UserStatusRequest{}, requestRaw); err == nil {
		return r.(*UserStatusRequest)
	}
	return nil
}

func (c *GameController) handleUserStatusRequest(requestRaw string, response *Response) string {
	request := parseUserStatusRequest(requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMailformed)
	}
	user, ok := c.users.Get(request.UserId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	response.Data[DataKeyUserId] = user.id
	response.Data[DataKeyUnit] = user.unit
	return response.WithStatus(ResponseStatusOk)
}
