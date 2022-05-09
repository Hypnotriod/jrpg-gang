package controller

type JoinGameRoomRequest struct {
	Request
	Data struct {
		RoomUid uint `json:"roomUid"`
	} `json:"data"`
}

func (c *GameController) handleJoinGameRoomRequest(requestRaw string, response *Response) string {
	request := parseRequest(&JoinGameRoomRequest{}, requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMailformed)
	}
	if c.rooms.ExistsForUserId(request.UserId) {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	if !c.rooms.Has(request.Data.RoomUid) {
		return response.WithStatus(ResponseStatusNotFound)
	}
	user, _ := c.users.Get(request.UserId)
	if !c.rooms.AddUser(request.Data.RoomUid, user) {
		return response.WithStatus(ResponseStatusFailed)
	}
	response.Data[DataKeyRoom], _ = c.rooms.GetByUserId(user.id)
	return response.WithStatus(ResponseStatusOk)
}
