package engine

import (
	"fmt"
	"jrpg-gang/util"
	"sync"
)

type GameEngine struct {
	*sync.RWMutex
	rndGen   *util.RndGen
	state    *GameState
	actors   []*GameUnit
	scenario *GameScenario
}

func (e GameEngine) String() string {
	return fmt.Sprintf(
		"state: {%v}, scenario: {%v}, actors: [%s]",
		e.state,
		e.scenario,
		util.AsCommaSeparatedSlice(e.actors),
	)
}

func NewGameEngine(scenario *GameScenario, actors []*GameUnit) *GameEngine {
	e := &GameEngine{}
	e.RWMutex = &sync.RWMutex{}
	e.rndGen = util.NewRndGen()
	e.state = NewGameState()
	e.scenario = scenario
	e.actors = actors
	scenario.Initialize(e.rndGen, actors)
	return e
}

func (e *GameEngine) Dispose() {
	e.scenario.Dispose()
	e.state = nil
	e.actors = nil
	e.scenario = nil
	e.rndGen = nil
}

func (e *GameEngine) GetPhase() GamePhase {
	return e.state.Phase
}

func (e *GameEngine) battlefield() *Battlefield {
	return &e.scenario.CurrentSpot().Battlefield
}

func (e *GameEngine) getActiveUnit() *GameUnit {
	if uid, ok := e.state.GetCurrentActiveUnitUid(); ok {
		return e.battlefield().FindUnitById(uid)
	}
	return nil
}

func (e *GameEngine) findActorByUserId(userId UserId) *GameUnit {
	for i := 0; i < len(e.actors); i++ {
		if e.actors[i].UserId == userId {
			return e.actors[i]
		}
	}
	return nil
}
