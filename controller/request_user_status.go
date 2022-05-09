package controller

type UserStatusRequest struct {
	Request
}

func (c *GameController) handleUserStatusRequest(requestRaw string, response *Response) string {
	request := parseRequest(&UserStatusRequest{}, requestRaw)
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
