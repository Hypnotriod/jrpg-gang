package engine

import (
	"fmt"
	"sync"
)

type GameEngine struct {
	*sync.RWMutex
	Spot     *Spot      `json:"spot"`
	State    *GameState `json:"state"`
	scenario *GameScenario
}

func (e GameEngine) String() string {
	return fmt.Sprintf(
		"spot: {%v}, state: {%v}",
		e.Spot,
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
		return e.Spot.Battlefield.FindUnitById(uid)
	}
	return nil
}

func (e *GameEngine) prepare() {
	e.Spot = e.scenario.CurrentSpot()
	e.State = NewGameState()
}
