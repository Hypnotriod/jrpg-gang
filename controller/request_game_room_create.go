package controller

type CreateGameRoomRequest struct {
	Request
	Data struct {
		Capacity    uint `json:"capacity"`
		ScenarioUid uint `json:"scenarioUid"`
	} `json:"data"`
}

func (c *GameController) handleCreateGameRoomRequest(requestRaw string, response *Response) string {
	request := parseRequest(&CreateGameRoomRequest{}, requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMailformed)
	}
	if request.Data.Capacity == 0 || request.Data.Capacity > GAME_ROOM_MAX_CAPACITY {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	if c.rooms.ExistsForUserId(request.UserId) {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	// todo: check available scenario
	hostUser, _ := c.users.Get(request.UserId)
	room := NewGameRoom()
	room.Capacity = request.Data.Capacity
	room.ScenarioUid = request.Data.ScenarioUid
	room.Host = hostUser
	c.rooms.Add(room)
	response.Data[DataKeyRoom] = room
	return response.WithStatus(ResponseStatusOk)
}
