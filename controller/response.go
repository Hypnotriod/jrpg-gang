package controller

import "jrpg-gang/util"

type ResponseStatus string

const (
	ResponseStatusOk            ResponseStatus = "ok"
	ResponseStatusError         ResponseStatus = "error"
	ResponseStatusMailformed    ResponseStatus = "mailformed"
	ResponseStatusUnsupported   ResponseStatus = "unsupported"
	ResponseStatusNotAllowed    ResponseStatus = "notAllowed"
	ResponseStatusAlreadyExists ResponseStatus = "alreadyExists"
)

type ResponseDataKey string

const (
	DataKeyActionResult ResponseDataKey = "actionResult"
	DataKeyGameState    ResponseDataKey = "gameState"
)

type Response struct {
	Type   RequestType                     `json:"type"`
	Id     string                          `json:"id"`
	UserId string                          `json:"userId,omitempty"`
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
