package controller

import "jrpg-gang/engine"

func (c *GameController) handleLobbyStatusRequest(userId engine.UserId, request *Request, response *Response) string {
	response.Data[DataKeyRooms] = c.rooms.GetAllRoomInfosList()
	response.Data[DataKeyUsersCount] = c.users.TotalCount()
	return response.WithStatus(ResponseStatusOk)
}
