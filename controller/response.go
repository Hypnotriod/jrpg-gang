package controller

import (
	"jrpg-gang/domain"
	"jrpg-gang/engine"
)

type ResponseStatus string

const (
	ResponseStatusOk         ResponseStatus = "ok"
	ResponseStatusError      ResponseStatus = "error"
	ResponseStatusMailformed ResponseStatus = "mailformed"
	ResponseStatusNotAllowed ResponseStatus = "notAllowed"
)

type Response struct {
	Status       ResponseStatus      `json:"status"`
	ActionResult domain.ActionResult `json:"actionResult,omitempty"`
	GameState    engine.GameEngine   `json:"gameState,omitempty"`
}

func (r *Response) WithStatus(status ResponseStatus) *Response {
	r.Status = status
	return r
}
