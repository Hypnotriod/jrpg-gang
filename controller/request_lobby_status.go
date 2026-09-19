package controller

import "jrpg-gang/engine"

func (c *GameController) handleLobbyStatusRequest(playerId engine.PlayerId, request *Request, response *Response) []byte {
	user, ok := c.users.Get(playerId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	response.Data[DataKeyRooms] = c.rooms.GetAllRoomInfos()
	response.Data[DataKeyScenarios] = c.scenariosConfig.GetAvailableScenarioConfigs(user.Unit)
	return response.WithStatus(ResponseStatusOk)
}
