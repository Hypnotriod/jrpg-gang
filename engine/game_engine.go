package engine

import (
	"fmt"
	"jrpg-gang/util"
	"sync"
)

type GameEngine struct {
	*sync.RWMutex
	rndGen   *util.RndGen
	spot     *Spot
	state    *GameState
	actors   []*GameUnit
	scenario *GameScenario
}

func (e GameEngine) String() string {
	return fmt.Sprintf(
		"spot: {%v}, state: {%v}",
		e.spot,
		e.state,
	)
}

func NewGameEngine(scenario *GameScenario, actors []*GameUnit) *GameEngine {
	e := &GameEngine{}
	e.RWMutex = &sync.RWMutex{}
	e.scenario = scenario
	e.actors = actors
	e.rndGen = util.NewRndGen()
	scenario.Initialize(e.rndGen, actors)
	e.prepare()
	return e
}

func (e *GameEngine) getActiveUnit() *GameUnit {
	if uid, ok := e.state.GetCurrentActiveUnitUid(); ok {
		return e.spot.Battlefield.FindUnitById(uid)
	}
	return nil
}

func (e *GameEngine) prepare() {
	e.spot = e.scenario.CurrentSpot()
	e.state = NewGameState()
}

func (e *GameEngine) findActorByUserId(userId UserId) *GameUnit {
	for i := 0; i < len(e.actors); i++ {
		if e.actors[i].UserId == userId {
			return e.actors[i]
		}
	}
	return nil
}
