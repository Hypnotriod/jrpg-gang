package controller

import "jrpg-gang/util"

type ResponseStatus string

const (
	ResponseStatusOk            ResponseStatus = "ok"
	ResponseStatusFailed        ResponseStatus = "failed"
	ResponseStatusError         ResponseStatus = "error"
	ResponseStatusMailformed    ResponseStatus = "mailformed"
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
	DataKeyUserId       ResponseDataKey = "userId"
	PlayerInfo          ResponseDataKey = "playerInfo"
	DataKeyUnit         ResponseDataKey = "unit"
	DataKeyRoom         ResponseDataKey = "room"
	DataKeyRooms        ResponseDataKey = "rooms"
	DataKeyShop         ResponseDataKey = "shop"
	DataKeyUsersCount   ResponseDataKey = "usersCount"
)

type Response struct {
	Type   RequestType                     `json:"type"`
	Id     string                          `json:"id,omitempty"`
	Status ResponseStatus                  `json:"status"`
	Data   map[ResponseDataKey]interface{} `json:"data,omitempty"`
}

func NewResponse() *Response {
	response := &Response{}
	response.Data = make(map[ResponseDataKey]interface{})
	return response
}

func (r *Response) fillUserStatus(user *User) {
	r.Data[PlayerInfo] = user.PlayerInfo
	r.Data[DataKeyUserId] = user.id
	r.Data[DataKeyUnit] = user.unit
}

func (r *Response) WithStatus(status ResponseStatus) string {
	r.Status = status
	return util.ObjectToJson(r)
}
