package engine

import (
	"jrpg-gang/domain"
	"jrpg-gang/util"
)

type GameEngine struct {
	rndGen   *util.RndGen
	state    *GameState
	actors   []*GameUnit
	scenario *GameScenario
}

func NewGameEngine(scenario *GameScenario, actors []*GameUnit) *GameEngine {
	e := &GameEngine{}
	e.scenario = scenario
	e.actors = actors
	e.rndGen = util.NewRndGen()
	e.state = NewGameState(scenario)
	e.scenario.Initialize(e.rndGen, actors)
	e.resetActorsReady()
	e.prepareActors(actors)
	e.prepareNextSpot(actors)
	return e
}

func (e *GameEngine) prepareActors(actors []*GameUnit) {
	for _, actor := range actors {
		actor.Booty = domain.UnitBooty{}
	}
}

func (e *GameEngine) resetActorsReady() {
	for _, actor := range e.actors {
		actor.PlayerInfo.IsReady = false
	}
}

func (e *GameEngine) UpdateActorReady(playerId PlayerId, value bool) bool {
	unit := util.Findp(e.actors, func(actor *GameUnit) bool {
		return actor.PlayerInfo.Id == playerId
	})
	if unit != nil {
		unit.PlayerInfo.IsReady = value
		return true
	}
	return false
}

func (e *GameEngine) AllActorsReady() bool {
	return util.Every(e.actors, func(actor *GameUnit) bool {
		return actor.PlayerInfo.IsReady
	})
}

func (e *GameEngine) AllActorsDead() bool {
	return util.Every(e.actors, func(actor *GameUnit) bool {
		return actor.IsDead
	})
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

func (e *GameEngine) GetPlayerIds() []PlayerId {
	result := []PlayerId{}
	for _, unit := range e.actors {
		result = append(result, unit.GetPlayerId())
	}
	return result
}

func (e *GameEngine) GetRestPlayerIds(playerId PlayerId) []PlayerId {
	result := []PlayerId{}
	for _, unit := range e.actors {
		if playerId != unit.GetPlayerId() {
			result = append(result, unit.GetPlayerId())
		}
	}
	return result
}

func (e *GameEngine) FindActorByPlayerId(playerId PlayerId) *GameUnit {
	return util.Findp(e.actors, func(u *GameUnit) bool {
		return u.GetPlayerId() == playerId
	})
}

func (e *GameEngine) RemoveActor(playerId PlayerId) bool {
	actor := e.FindActorByPlayerId(playerId)
	if actor == nil {
		return false
	}
	if e.state.IsCurrentActiveUnit(actor) && e.isActionPhase() {
		actor.ClearActionPoints()
		e.onUnitCompleteAction(nil)
	}
	actor.PlayerInfo.IsReady = false
	e.battlefield().MoveToCorpses(actor.Uid)
	e.state.UpdateUnitsQueue(e.battlefield().Units)
	restActors := []*GameUnit{}
	for i := 0; i < len(e.actors); i++ {
		if e.actors[i].GetPlayerId() != playerId {
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

func (e *GameEngine) UpdateUserConnectionStatus(playerId PlayerId, isOffline bool) bool {
	actor := e.FindActorByPlayerId(playerId)
	if actor == nil {
		return false
	}
	actor.PlayerInfo.IsOffline = isOffline
	return true
}

func (e *GameEngine) ClearActiveUnitActionPoints() {
	if unit := e.getActiveUnit(); unit != nil {
		unit.ClearActionPoints()
	}
}

func (e *GameEngine) battlefield() *Battlefield {
	return &e.scenario.CurrentSpot().Battlefield
}

func (e *GameEngine) getActiveUnit() *GameUnit {
	if uid, ok := e.state.GetCurrentActiveUnitUid(); ok {
		return e.battlefield().FindUnit(uid)
	}
	return nil
}
