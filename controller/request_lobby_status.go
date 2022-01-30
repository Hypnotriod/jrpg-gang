package controller

import "jrpg-gang/util"

type LobbyStatusRequest struct {
	Request
}

func parseLobbyStatusRequest(requestRaw string) *LobbyStatusRequest {
	if r, ok := util.JsonToObject(&LobbyStatusRequest{}, requestRaw); ok {
		return r.(*LobbyStatusRequest)
	}
	return nil
}

func (c *GameController) handleLobbyStatusRequest(requestRaw string, response *Response) string {
	request := parseLobbyStatusRequest(requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMailformed)
	}
	response.Data[DataKeyRooms] = c.rooms.ResponseList()
	response.Data[DataKeyUsersCount] = c.users.TotalCount()
	return response.WithStatus(ResponseStatusOk)
}
