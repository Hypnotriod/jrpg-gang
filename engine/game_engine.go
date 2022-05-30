package engine

import (
	"fmt"
	"jrpg-gang/util"
)

type GameEngine struct {
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
		util.AsCommaSeparatedObjectsSlice(e.actors),
	)
}

func NewGameEngine(scenario *GameScenario, actors []*GameUnit) *GameEngine {
	e := &GameEngine{}
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
	return e.state.phase
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

func (e *GameEngine) GetUserIds() []UserId {
	result := []UserId{}
	for _, unit := range e.actors {
		result = append(result, unit.userId)
	}
	return result
}

func (e *GameEngine) GetRestUserIds(userId UserId) []UserId {
	result := []UserId{}
	for _, unit := range e.actors {
		if userId != unit.userId {
			result = append(result, unit.userId)
		}
	}
	return result
}

func (e *GameEngine) findActorByUserId(userId UserId) *GameUnit {
	for i := 0; i < len(e.actors); i++ {
		if e.actors[i].userId == userId {
			return e.actors[i]
		}
	}
	return nil
}
