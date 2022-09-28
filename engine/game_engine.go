package engine

import (
	"fmt"
	"jrpg-gang/domain"
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
	e.scenario.Initialize(e.rndGen, actors)
	e.prepareActors(actors)
	e.prepareNextSpot(actors)
	return e
}

func (e *GameEngine) prepareActors(actors []*GameUnit) {
	for _, actor := range actors {
		actor.Booty = domain.UnitBooty{}
	}
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

func (e *GameEngine) GetPlayersInfo() []PlayerInfo {
	result := []PlayerInfo{}
	for _, unit := range e.actors {
		result = append(result, *unit.PlayerInfo)
	}
	return result
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

func (e *GameEngine) FindActorByUserId(userId UserId) *GameUnit {
	return util.Findp(e.actors, func(u *GameUnit) bool {
		return u.userId == userId
	})
}

func (e *GameEngine) RemoveActor(userId UserId) bool {
	actor := e.FindActorByUserId(userId)
	if actor == nil {
		return false
	}
	if e.state.IsCurrentActiveUnit(actor) && e.isActionPhase() {
		e.onUnitCompleteAction()
	}
	e.battlefield().MoveToCorpsesById(actor.Uid)
	e.state.UpdateUnitsQueue(e.battlefield().Units)
	restActors := []*GameUnit{}
	for i := 0; i < len(e.actors); i++ {
		if e.actors[i].userId != userId {
			restActors = append(restActors, e.actors[i])
		}
	}
	e.actors = restActors
	return true
}

func (e *GameEngine) TakeAShare() domain.UnitBooty {
	leftUnits := e.battlefield().FactionUnitsCount(GameUnitFactionLeft)
	if !e.canTakeAShare() || leftUnits == 0 {
		return domain.UnitBooty{}
	}
	return e.state.Booty.TakeAShare(leftUnits)
}

func (e *GameEngine) UpdateUserConnectionStatus(userId UserId, isOffline bool) bool {
	actor := e.FindActorByUserId(userId)
	if actor == nil {
		return false
	}
	actor.PlayerInfo.IsOffline = isOffline
	return true
}
