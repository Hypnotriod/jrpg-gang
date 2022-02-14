package engine

import (
	"fmt"
	"sync"
)

type GameEngine struct {
	*sync.RWMutex
	Spot     *Spot      `json:"spot"`
	State    *GameState `json:"state"`
	actors   []*GameUnit
	scenario *GameScenario
}

func (e GameEngine) String() string {
	return fmt.Sprintf(
		"spot: {%v}, state: {%v}",
		e.Spot,
		e.State,
	)
}

func NewGameEngine(scenario *GameScenario, actors []*GameUnit) *GameEngine {
	scenario.Initialize(actors)
	e := &GameEngine{}
	e.RWMutex = &sync.RWMutex{}
	e.scenario = scenario
	e.actors = actors
	e.prepare()
	return e
}

func (e *GameEngine) getActiveUnit() *GameUnit {
	if uid, ok := e.State.GetCurrentActiveUnitUid(); ok {
		return e.Spot.Battlefield.FindUnitById(uid)
	}
	return nil
}

func (e *GameEngine) prepare() {
	e.Spot = e.scenario.CurrentSpot()
	e.State = NewGameState()
}

func (e *GameEngine) findActorByUserId(userId UserId) *GameUnit {
	for i := 0; i < len(e.actors); i++ {
		if e.actors[i].UserId == userId {
			return e.actors[i]
		}
	}
	return nil
}
