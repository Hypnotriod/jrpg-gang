package controller

import "jrpg-gang/engine"

func (c *GameController) handleLobbyStatusRequest(playerId engine.PlayerId, request *Request, response *Response) []byte {
	response.Data[DataKeyRooms] = c.rooms.GetAllRoomInfos()
	response.Data[DataKeyScenarios] = c.scenariosConfig.GetAllScenarioConfigs()
	return response.WithStatus(ResponseStatusOk)
}
