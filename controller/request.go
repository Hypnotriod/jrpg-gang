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
	RequestUserStatus      RequestType = "userStatus"
	RequestStartGame       RequestType = "startGame"
	RequestGameAction      RequestType = "gameAction"
	RequestGameState       RequestType = "gameState"
)

type Request struct {
	UserId engine.UserId `json:"userId,omitempty"`
	Type   RequestType   `json:"type"`
	Id     string        `json:"id"`
}

type ParsebleRequest interface {
	*Request | *GameActionRequest | *CreateGameRoomRequest | *JoinRequest | *DestroyGameRoomRequest |
		*LobbyStatusRequest | *UserStatusRequest | *JoinGameRoomRequest | *LeaveGameRoomRequest |
		*StartGameRequest | *GameStateRequest
}

func parseRequest[T ParsebleRequest](data T, requestRaw string) T {
	r, err := util.JsonToObject(data, requestRaw)
	if err == nil {
		return r.(T)
	}
	return nil
}
