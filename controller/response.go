package controller

type ResponseStatus string

const (
	ResponseStatusOk         ResponseStatus = "ok"
	ResponseStatusError      ResponseStatus = "error"
	ResponseStatusMailformed ResponseStatus = "mailformed"
	ResponseStatusNotAllowed ResponseStatus = "notAllowed"
)

type ResponseDataKey string

const (
	DataKeyActionResult ResponseDataKey = "actionResult"
	DataKeyGameState    ResponseDataKey = "gameState"
)

type Response struct {
	Status ResponseStatus                  `json:"status"`
	Data   map[ResponseDataKey]interface{} `json:"data,omitempty"`
}

func NewResponse() *Response {
	response := &Response{}
	response.Data = make(map[ResponseDataKey]interface{})
	return response
}

func (r *Response) WithStatus(status ResponseStatus) *Response {
	r.Status = status
	return r
}
