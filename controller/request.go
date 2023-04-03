package controller

import (
	"bytes"

	jsoniter "github.com/json-iterator/go"
)

var json jsoniter.API = jsoniter.ConfigCompatibleWithStandardLibrary

type RequestType string

const (
	RequestJoin               RequestType = "join"
	RequestLeave              RequestType = "leave"
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
	RequestApplyForAJob       RequestType = "applyForAJob"
	RequestQuitJob            RequestType = "quitJob"
	RequestCompleteJob        RequestType = "completeJob"
	RequestJobsStatus         RequestType = "jobsStatus"
)

type Request struct {
	Type RequestType `json:"type"`
	Id   string      `json:"id"`
	Data []byte      `json:"-"`
}

type ParsebleRequestData interface {
	*Request | *GameActionRequestData | *CreateGameRoomRequestData | *JoinRequestData |
		*JoinGameRoomRequestData | *ShopActionRequestData | *ConfiguratorActionRequestData |
		*GameNextPhaseRequestData | *ApplyForAJobRequestData
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
	if len(r) < 10 || !bytes.Equal(r[:9], []byte(`{"type":"`)) {
		return nil
	}
	r = r[9:]
	for i, c := range r {
		if i == len(r)-1 {
			return nil
		} else if c == '"' {
			typeStr = string(r[:i])
			r = r[i+1:]
			break
		}
	}
	if len(r) < 8 || !bytes.Equal(r[:7], []byte(`,"id":"`)) {
		return nil
	}
	r = r[7:]
	for i, c := range r {
		if i == len(r)-1 {
			return nil
		} else if c == '"' {
			idStr = string(r[:i])
			r = r[i+1:]
			break
		}
	}
	if len(r) < 10 || !bytes.Equal(r[:9], []byte(`,"data":{`)) {
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
