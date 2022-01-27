package controller

import "jrpg-gang/engine"

type RequestType string

const (
	RequestCreateBattleRoom RequestType = "createBattleRoom"
	RequestplaceUnit        RequestType = "placeUnit"
)

type Request struct {
	UserId string      `json:"userId"`
	Type   RequestType `json:"type"`
	Data   RequestData `json:"data"`
}

type RequestData struct {
	AllowedUsers []string        `json:"allowedUsers,omitempty"`
	Matrix       [][]engine.Cell `json:"matrix,omitempty"`
	Unit         engine.GameUnit `json:"unit,omitempty"`
}
