package controller

import "encoding/json"

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
	Data []byte      `json:"-"`
}

type ParsebleRequestData interface {
	*Request | *GameActionRequestData | *CreateGameRoomRequestData | *JoinRequestData |
		*JoinGameRoomRequestData | *ShopActionRequestData | *ConfiguratorActionRequestData
}

func parseRequestData[T ParsebleRequestData](data T, requestRaw []byte) T {
	err := json.Unmarshal(requestRaw, data)
	if err == nil {
		return data
	}
	return nil
}

func parseRequest(r []byte) *Request {
	var typeStr string
	var idStr string
	if len(r) < 10 || r[0] != '{' || r[1] != '"' || r[2] != 't' || r[3] != 'y' || r[4] != 'p' || r[5] != 'e' || r[6] != '"' || r[7] != ':' || r[8] != '"' {
		return nil
	}
	typeBytes := [32]byte{}
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
	idBytes := [32]byte{}
	for i, c := range r {
		if c == '"' {
			r = r[i+1:]
			idStr = string(idBytes[:i])
			break
		} else if !(c >= 'a' && c <= 'h' || c >= '0' && c <= '9') || i == len(idBytes) || i == len(r)-1 {
			return nil
		} else {
			idBytes[i] = c
		}
	}
	if len(r) < 10 || r[0] != ',' || r[1] != '"' || r[2] != 'd' || r[3] != 'a' || r[4] != 't' || r[5] != 'a' || r[6] != '"' || r[7] != ':' || r[8] != '{' {
		return &Request{
			Type: RequestType(typeStr),
			Id:   idStr,
		}
	}
	r = r[8:]
	n := 0
	for i, c := range r {
		if c == '{' {
			n++
		} else if c == '}' {
			n--
		}
		if n == 0 {
			r = r[:i+1]
			break
		}
		if i == len(r)-1 {
			return nil
		}
	}
	return &Request{
		Type: RequestType(typeStr),
		Id:   idStr,
		Data: r,
	}
}
