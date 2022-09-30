package controller

import (
	"jrpg-gang/util"
)

type RequestType string

const (
	RequestJoin               RequestType = "join"
	RequestEnterLobby         RequestType = "enterLobby"
	RequestExitLobby          RequestType = "exitLobby"
	RequestCreateGameRoom     RequestType = "createRoom"
	RequestDestroyGameRoom    RequestType = "destroyRoom"
	RequestJoinGameRoom       RequestType = "joinRoom"
	RequestLeaveGameRoom      RequestType = "leaveRoom"
	RequestLobbyStatus        RequestType = "lobbyStatus"
	RequestRoomStatus         RequestType = "roomStatus"
	RequestConfiguratorAction RequestType = "configuratorAction"
	RequestUserStatus         RequestType = "userStatus"
	RequestShopStatus         RequestType = "shopStatus"
	RequestShopAction         RequestType = "shopAction"
	RequestStartGame          RequestType = "startGame"
	RequestLeaveGame          RequestType = "leaveGame"
	RequestGameAction         RequestType = "gameAction"
	RequestNextGamePhase      RequestType = "nextGamePhase"
	RequestGameState          RequestType = "gameState"
	RequestPlayerInfo         RequestType = "playerInfo"
)

type Request struct {
	Type RequestType `json:"type"`
	Id   string      `json:"id"`
}

type ParsebleRequest interface {
	*Request | *GameActionRequest | *CreateGameRoomRequest | *JoinRequest | *DestroyGameRoomRequest |
		*LobbyStatusRequest | *UserStatusRequest | *JoinGameRoomRequest | *LeaveGameRoomRequest |
		*StartGameRequest | *GameStateRequest | *GameNextPhaseRequest | *ShopStatusRequest | *ShopActionRequest |
		*ConfiguratorActionRequest | *PlayerInfoRequest | *GameLeaveRequest | *EnterLobbyRequest | *ExitLobbyRequest
}

func parseRequest[T ParsebleRequest](data T, requestRaw string) T {
	r, err := util.JsonToObject(data, requestRaw)
	if err == nil {
		return r.(T)
	}
	return nil
}

func parseRequestManual(raw string) *Request {
	var typeStr string
	var idStr string
	r := []byte(raw)
	if len(r) < 10 || r[0] != '{' || r[1] != '"' || r[2] != 't' || r[3] != 'y' || r[4] != 'p' || r[5] != 'e' || r[6] != '"' || r[7] != ':' || r[8] != '"' {
		return nil
	}
	typeBytes := [16]byte{}
	r = r[9:]
	for i, c := range r {
		if c == '"' {
			r = r[i+1:]
			typeStr = string(typeBytes[:i])
			break
		} else if !(c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z') || i == len(typeBytes) || i == len(r)-1 {
			return nil
		} else {
			typeBytes[i] = c
		}
	}
	if len(r) < 8 || r[0] != ',' || r[1] != '"' || r[2] != 'i' || r[3] != 'd' || r[4] != '"' || r[5] != ':' || r[6] != '"' {
		return nil
	}
	r = r[7:]
	idBytes := [16]byte{}
	for i, c := range r {
		if c == '"' {
			idStr = string(idBytes[:i])
			break
		} else if !(c >= 'a' && c <= 'h' || c >= '0' && c <= '9') || i == len(idBytes) || i == len(r)-1 {
			return nil
		} else {
			idBytes[i] = c
		}
	}
	return &Request{
		Type: RequestType(typeStr),
		Id:   idStr,
	}
}
