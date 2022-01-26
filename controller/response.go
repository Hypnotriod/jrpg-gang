package controller

type ResponseStatus string

const (
	ResponseStatusOk         ResponseStatus = "ok"
	ResponseStatusMailformed ResponseStatus = "mailformed"
)

type Response struct {
	UserId string      `json:"userId,omitempty"`
	Type   RequestType `json:"type,omitempty"`
	Status ResponseStatus
}
