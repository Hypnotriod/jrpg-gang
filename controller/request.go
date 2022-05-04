package controller

import (
	"jrpg-gang/engine"
	"jrpg-gang/util"
)

type RequestType string

const (
	RequestJoin            RequestType = "join"
	RequestCreateGameRoom  RequestType = "createRoom"
	RequestDestroyGameRoom RequestType = "destroyRoom"
	RequestJoinGameRoom    RequestType = "joinRoom"
	RequestLeaveGameRoom   RequestType = "leaveRoom"
	RequestLobbyStatus     RequestType = "lobbyStatus"
	RequestStartGame       RequestType = "startGame"
)

type Request struct {
	Type   RequestType   `json:"type"`
	Id     string        `json:"id"`
	UserId engine.UserId `json:"userId,omitempty"`
}

func parseRequest(requestRaw string) *Request {
	if r, err := util.JsonToObject(&Request{}, requestRaw); err == nil {
		return r.(*Request)
	}
	return nil
}
