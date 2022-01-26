package engine

import (
	"fmt"
)

type GameEngine struct {
	Battlefield  *Battlefield `json:"battlefield"`
	State        GameState    `json:"state"`
	AllowedUsers []string     `json:"allowedUsers"`
}

func (e GameEngine) String() string {
	return fmt.Sprintf(
		"battlefield: {%v}, state: {%v}",
		e.Battlefield,
		e.State,
	)
}

func NewGameEngine(battlefield *Battlefield) *GameEngine {
	engine := &GameEngine{}
	engine.Battlefield = battlefield
	return engine
}

func (e *GameEngine) GetActiveUnit() *GameUnit {
	if uid, ok := e.State.GetActiveUnitUid(); ok {
		return e.Battlefield.FindUnitById(uid)
	}
	return nil
}
