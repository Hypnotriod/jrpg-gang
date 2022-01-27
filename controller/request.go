package controller

import (
	"jrpg-gang/engine"
	"jrpg-gang/util"
)

type RequestType string

const (
	RequestCreateBattleRoom RequestType = "createBattleRoom"
	RequestPlaceUnit        RequestType = "placeUnit"
)

type Request struct {
	UserId string      `json:"userId"`
	Type   RequestType `json:"type"`
}

func parseRequest(requestRaw string) *Request {
	if r, ok := util.JsonToObject(&Request{}, requestRaw); ok {
		return r.(*Request)
	}
	return nil
}

type CreateBattleRoomRequest struct {
	Request
	AllowedUsers []string        `json:"allowedUsers"`
	Matrix       [][]engine.Cell `json:"matrix"`
}

func parseCreateBattleRoomRequest(requestRaw string) *CreateBattleRoomRequest {
	if r, ok := util.JsonToObject(&CreateBattleRoomRequest{}, requestRaw); ok {
		return r.(*CreateBattleRoomRequest)
	}
	return nil
}

type PlaceUnitRequest struct {
	Request
	Unit engine.GameUnit `json:"unit"`
}

func parsePlaceUnitRequest(requestRaw string) *PlaceUnitRequest {
	if r, ok := util.JsonToObject(&PlaceUnitRequest{}, requestRaw); ok {
		return r.(*PlaceUnitRequest)
	}
	return nil
}
