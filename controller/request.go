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
	Type RequestType `json:"type"`
	Id   string      `json:"id"`
}

func parseRequest(requestRaw string) *Request {
	if r, ok := util.JsonToObject(&Request{}, requestRaw); ok {
		return r.(*Request)
	}
	return nil
}

type JoinRequest struct {
	Request
	Data struct {
		Nickname string `json:"nickname"`
	} `json:"data"`
}

func parseJoinRequest(requestRaw string) *JoinRequest {
	if r, ok := util.JsonToObject(&JoinRequest{}, requestRaw); ok {
		return r.(*JoinRequest)
	}
	return nil
}

type CreateBattleRoomRequest struct {
	Request
	Data struct {
		UserId       string          `json:"userId"`
		AllowedUsers []string        `json:"allowedUsers"`
		Matrix       [][]engine.Cell `json:"matrix"`
	} `json:"data"`
}

func parseCreateBattleRoomRequest(requestRaw string) *CreateBattleRoomRequest {
	if r, ok := util.JsonToObject(&CreateBattleRoomRequest{}, requestRaw); ok {
		return r.(*CreateBattleRoomRequest)
	}
	return nil
}

type PlaceUnitRequest struct {
	Request
	Data struct {
		UserId string          `json:"userId"`
		Unit   engine.GameUnit `json:"unit"`
	} `json:"data"`
}

func parsePlaceUnitRequest(requestRaw string) *PlaceUnitRequest {
	if r, ok := util.JsonToObject(&PlaceUnitRequest{}, requestRaw); ok {
		return r.(*PlaceUnitRequest)
	}
	return nil
}
