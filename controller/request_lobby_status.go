package controller

import "jrpg-gang/engine"

func (c *GameController) handleLobbyStatusRequest(playerId engine.PlayerId, request *Request, response *Response) []byte {
	response.Data[DataKeyRooms] = c.rooms.GetAllRoomInfosList()
	response.Data[DataKeyUsersCount] = c.users.TotalCount()
	return response.WithStatus(ResponseStatusOk)
}
