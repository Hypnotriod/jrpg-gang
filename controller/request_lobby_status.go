package controller

import "jrpg-gang/engine"

func (c *GameController) handleLobbyStatusRequest(playerId engine.PlayerId, request *Request, response *Response) []byte {
	response.Data[DataKeyRooms] = c.rooms.GetAllRoomInfosList()
	return response.WithStatus(ResponseStatusOk)
}
