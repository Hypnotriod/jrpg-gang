package controller

import (
	"jrpg-gang/controller/users"
)

type ResponseStatus string

const (
	ResponseStatusOk            ResponseStatus = "ok"
	ResponseStatusFailed        ResponseStatus = "failed"
	ResponseStatusError         ResponseStatus = "error"
	ResponseStatusMalformed     ResponseStatus = "malformed"
	ResponseStatusUnsupported   ResponseStatus = "unsupported"
	ResponseStatusNotAllowed    ResponseStatus = "notAllowed"
	ResponseStatusNotFound      ResponseStatus = "notFound"
	ResponseStatusAlreadyExists ResponseStatus = "alreadyExists"
)

type ResponseDataKey string

const (
	DataKeyAction       ResponseDataKey = "action"
	DataKeyActionResult ResponseDataKey = "actionResult"
	DataKeyGameState    ResponseDataKey = "gameState"
	DataKeySessionId    ResponseDataKey = "sessionId"
	DataKeyPlayerInfo   ResponseDataKey = "playerInfo"
	DataKeyPlayersInfo  ResponseDataKey = "playersInfo"
	DataKeyStatus       ResponseDataKey = "status"
	DataKeyUnit         ResponseDataKey = "unit"
	DataKeyRoom         ResponseDataKey = "room"
	DataKeyRooms        ResponseDataKey = "rooms"
	DataKeyShop         ResponseDataKey = "shop"
	DataKeyUsersCount   ResponseDataKey = "usersCount"
	DataKeyEmployment   ResponseDataKey = "employment"
	DataKeyReward       ResponseDataKey = "reward"
)

type Response struct {
	Type   RequestType                     `json:"type,omitempty"`
	Id     string                          `json:"id,omitempty"`
	Status ResponseStatus                  `json:"status"`
	Data   map[ResponseDataKey]interface{} `json:"data,omitempty"`
}

func NewResponse() *Response {
	response := &Response{}
	response.Data = make(map[ResponseDataKey]interface{})
	return response
}

func (r *Response) fillUserStatus(user *users.User) {
	r.Data[DataKeyPlayerInfo] = user.PlayerInfo
	r.Data[DataKeySessionId] = user.SessionId
	r.Data[DataKeyUnit] = user.Unit
	r.Data[DataKeyStatus] = user.Status.Display()
}

func (r *Response) WithStatus(status ResponseStatus) []byte {
	r.Status = status
	if marshalled, err := json.Marshal(r); err == nil {
		return marshalled
	}
	return []byte{}
}
