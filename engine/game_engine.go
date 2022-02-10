package engine

import (
	"fmt"
	"sync"
)

type GameEngine struct {
	*sync.RWMutex
	Battlefield *Battlefield `json:"battlefield"`
	State       *GameState   `json:"state"`
	scenario    *GameScenario
}

func (e GameEngine) String() string {
	return fmt.Sprintf(
		"battlefield: {%v}, state: {%v}",
		e.Battlefield,
		e.State,
	)
}

func NewGameEngine(scenario *GameScenario) *GameEngine {
	scenario.Initialize()
	e := &GameEngine{}
	e.RWMutex = &sync.RWMutex{}
	e.scenario = scenario
	e.prepare()
	return e
}

func (e *GameEngine) GetActiveUnit() *GameUnit {
	defer e.RUnlock()
	e.RLock()
	if uid, ok := e.State.GetActiveUnitUid(); ok {
		return e.Battlefield.FindUnitById(uid)
	}
	return nil
}

func (e *GameEngine) prepare() {
	e.Battlefield = e.scenario.CurrentBattlefield()
	e.State = NewGameState()
}
