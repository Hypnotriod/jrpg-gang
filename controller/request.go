package controller

import (
	"jrpg-gang/engine"
	"jrpg-gang/util"
)

type RequestType string

const (
	RequestJoin             RequestType = "join"
	RequestCreateBattleRoom RequestType = "createBattleRoom"
	RequestPlaceUnit        RequestType = "placeUnit"
)

type Request struct {
	Type   RequestType `json:"type"`
	Id     string      `json:"id"`
	UserId string      `json:"userId,omitempty"`
}

func parseRequest(requestRaw string) *Request {
	if r, ok := util.JsonToObject(&Request{}, requestRaw); ok {
		return r.(*Request)
	}
	return nil
}

type JoinRequest struct {
	Request
	Nickname string `json:"nickname"`
}

func parseJoinRequest(requestRaw string) *JoinRequest {
	if r, ok := util.JsonToObject(&JoinRequest{}, requestRaw); ok {
		return r.(*JoinRequest)
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
