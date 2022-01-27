package engine

import (
	"fmt"
	"sync"
)

type GameEngine struct {
	*sync.RWMutex
	Battlefield *Battlefield `json:"battlefield"`
	State       GameState    `json:"state"`
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
	defer e.RUnlock()
	e.RLock()
	if uid, ok := e.State.GetActiveUnitUid(); ok {
		return e.Battlefield.FindUnitById(uid)
	}
	return nil
}
