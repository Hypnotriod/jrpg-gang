package controller

import (
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
)

type KickFromGameRoomRequestData struct {
	PlayerId engine.PlayerId `json:"playerId"`
}

func (c *GameController) handleKickFromGameRoomRequest(playerId engine.PlayerId, request *Request, response *Response) []byte {
	data := parseRequestData(&KickFromGameRoomRequestData{}, request.Data)
	if data == nil {
		return response.WithStatus(ResponseStatusMalformed)
	}
	if !c.rooms.ExistsForHostId(playerId) {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	roomUid, ok := c.rooms.GetUidByPlayerId(playerId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	if !c.rooms.KickUser(roomUid, data.PlayerId) {
		return response.WithStatus(ResponseStatusFailed)
	}
	c.users.ChangeUserStatus(data.PlayerId, users.UserStatusInLobby)
	response.Data[DataKeyRoom] = c.rooms.GetRoomInfoByUid(roomUid)
	c.broadcastRoomStatus(roomUid)
	return response.WithStatus(ResponseStatusOk)
}
