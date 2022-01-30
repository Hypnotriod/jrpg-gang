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
	DataKeyActionResult ResponseDataKey = "actionResult"
	DataKeyGameState    ResponseDataKey = "gameState"
	DataKeyUserId       ResponseDataKey = "userId"
	DataKeyRoom         ResponseDataKey = "room"
	DataKeyRooms        ResponseDataKey = "rooms"
	DataKeyUsersCount   ResponseDataKey = "usersCount"
)

type Response struct {
	Type   RequestType                     `json:"type"`
	Id     string                          `json:"id"`
	Status ResponseStatus                  `json:"status"`
	Data   map[ResponseDataKey]interface{} `json:"data,omitempty"`
}

func NewResponse() *Response {
	response := &Response{}
	response.Data = make(map[ResponseDataKey]interface{})
	return response
}

func (r *Response) WithStatus(status ResponseStatus) string {
	r.Status = status
	return util.ObjectToJson(r)
}