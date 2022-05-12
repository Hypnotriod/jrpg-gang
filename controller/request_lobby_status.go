package controller

import "jrpg-gang/engine"

type LobbyStatusRequest struct {
	Request
}

func (c *GameController) handleLobbyStatusRequest(userId engine.UserId, requestRaw string, response *Response) string {
	request := parseRequest(&LobbyStatusRequest{}, requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMailformed)
	}
	response.Data[DataKeyRooms] = c.rooms.ResponseList()
	response.Data[DataKeyUsersCount] = c.users.TotalCount()
	return response.WithStatus(ResponseStatusOk)
}
