package controller

import (
	"jrpg-gang/domain"
	"jrpg-gang/engine"
)

type GameRoomDismissMercenaryRequestData struct {
	Code domain.UnitCode `json:"code"`
}

func (c *GameController) handleGameRoomDismissMercenaryRequest(playerId engine.PlayerId, request *Request, response *Response) []byte {
	data := parseRequestData(&GameRoomDismissMercenaryRequestData{}, request.Data)
	if data == nil {
		return response.WithStatus(ResponseStatusMalformed)
	}
	if !c.rooms.ExistsForHostId(playerId) {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	roomUid, ok := c.rooms.RemoveMercenary(playerId, data.Code)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	user, _ := c.users.Get(playerId)
	c.mercenaries.Refund(data.Code, &user.Unit.Unit)
	c.persistUser(&user)
	response.Data[DataKeyRoom] = c.rooms.GetRoomInfoByUid(roomUid)
	c.broadcastRoomStatus(roomUid)
	return response.WithStatus(ResponseStatusOk)
}
