package controller

import (
	"jrpg-gang/domain"
	"jrpg-gang/engine"
)

type GameRoomHireMercenaryRequestData struct {
	Code domain.UnitCode `json:"code"`
}

func (c *GameController) handleGameRoomHireMercenaryRequest(playerId engine.PlayerId, request *Request, response *Response) []byte {
	data := parseRequestData(&GameRoomHireMercenaryRequestData{}, request.Data)
	if data == nil {
		return response.WithStatus(ResponseStatusMalformed)
	}
	if !c.rooms.ExistsForHostId(playerId) {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	user, _ := c.users.Get(playerId)
	mercenary := c.mercenaries.Hire(data.Code, &user.Unit.Unit)
	if mercenary == nil {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	roomUid, ok := c.rooms.AddMercenary(user.Id, mercenary)
	if !ok {
		return response.WithStatus(ResponseStatusFailed)
	}
	response.Data[DataKeyRoom] = c.rooms.GetRoomInfoByUid(roomUid)
	c.broadcastRoomStatus(roomUid)
	return response.WithStatus(ResponseStatusOk)
}
