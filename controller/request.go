package controller

import (
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
)

type Request struct {
	Type   RequestType `json:"type"`
	Id     string      `json:"id"`
	UserId UserId      `json:"userId,omitempty"`
}

func parseRequest(requestRaw string) *Request {
	if r, ok := util.JsonToObject(&Request{}, requestRaw); ok {
		return r.(*Request)
	}
	return nil
}
