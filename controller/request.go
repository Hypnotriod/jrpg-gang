package controller

import "jrpg-gang/engine"

type RequestType string

const (
	RequestCreateGameEngine RequestType = "createEngine"
)

type Request struct {
	UserId string      `json:"userId"`
	Type   RequestType `json:"type"`
	Data   RequestData `json:"data"`
}

type RequestData struct {
	AllowedUsers []string        `json:"allowedUsers,omitempty"`
	Matrix       [][]engine.Cell `json:"matrix,omitempty"`
}
