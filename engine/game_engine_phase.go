package engine

import (
	"jrpg-gang/domain"
	"jrpg-gang/util"
)

func (e *GameEngine) NextPhase() *GameEvent {
	result := e.NewGameEvent()
	switch e.state.phase {
	case GamePhaseReadyForStartRound, GamePhasePrepareUnit:
		e.processStartRound()
	case GamePhaseMakeMoveOrActionAI, GamePhaseMakeActionAI:
		e.processAI(result)
	case GamePhaseRetreatAction:
		e.processRetreatActionAI(result)
	case GamePhaseActionComplete:
		e.processActionComplete(result)
	case GamePhaseBattleComplete:
		e.processBattleComplete(result)
	}
	e.resetActorsReady()
	result.PlayersInfo = e.GetPlayersInfo()
	result.NextPhase = e.state.phase
	return result
}

func (e *GameEngine) NextPhaseRequired() bool {
	return e.state.phase == GamePhaseReadyForStartRound ||
		e.state.phase == GamePhasePrepareUnit ||
		e.state.phase == GamePhaseMakeMoveOrActionAI ||
		e.state.phase == GamePhaseMakeActionAI ||
		e.state.phase == GamePhaseRetreatAction ||
		e.state.phase == GamePhaseActionComplete ||
		e.state.phase == GamePhaseBattleComplete
}

func (e *GameEngine) prepareNextSpot(actors []*GameUnit) {
	e.scenario.PrepareNextSpot(actors)
	e.state.MakeUnitsQueue(e.battlefield().Units)
}

func (e *GameEngine) processStartRound() {
	e.state.MakeUnitsQueue(e.battlefield().Units)
	e.battlefield().UpdateCellsFactions()
	e.switchToNextUnit()
}

func (e *GameEngine) processActionComplete(event *GameEvent) {
	if !e.state.HasActiveUnits() {
		e.processRoundComplete(event)
	} else {
		e.switchToNextUnit()
	}
}

func (e *GameEngine) processRoundComplete(event *GameEvent) {
	if e.endRound(event) {
		if e.scenario.IsLastSpot() {
			e.state.ChangePhase(GamePhaseDungeonComplete)
		} else {
			e.state.ChangePhase(GamePhaseBattleComplete)
		}
	} else {
		e.state.ChangePhase(GamePhaseReadyForStartRound)
	}
}

func (e *GameEngine) processBattleComplete(event *GameEvent) {
	e.prepareNextSpot(e.battlefield().Units)
	e.state.ChangePhase(GamePhasePrepareUnit)
	event.Spot = e.scenario.CurrentSpot()
}

func (e *GameEngine) switchToNextUnit() {
	unit := e.getActiveUnit()
	unit.State.IsStunned = false
	if unit.CheckRandomChance(unit.CalculateRetreatChance()) {
		e.state.ChangePhase(GamePhaseRetreatAction)
	} else if unit.HasUserId() {
		e.state.ChangePhase(GamePhaseMakeMoveOrAction)
	} else {
		e.state.ChangePhase(GamePhaseMakeMoveOrActionAI)
	}
}

func (e *GameEngine) endRound(event *GameEvent) (isLastRound bool) {
	result := NewEndTurnResult()
	for _, unit := range e.battlefield().Units {
		result.Recovery[unit.Uid] = unit.ApplyRecoverylOnNextTurn()
		result.Damage[unit.Uid] = unit.ApplyDamageOnNextTurn()
		unit.ReduceModificationOnNextTurn()
	}
	corpses := e.battlefield().FilterSurvivors()
	e.applyExperience(corpses)
	event.EndRoundResult = result
	isLastRound = e.battlefield().FactionsCount() <= 1
	if isLastRound && e.battlefield().FactionUnitsCount(GameUnitFactionLeft) != 0 {
		e.accumulateBooty(event)
	}
	return
}

func (e *GameEngine) onUnitMoveAction() {
	unit := e.getActiveUnit()
	if unit.HasUserId() {
		e.state.ChangePhase(GamePhaseMakeAction)
	} else {
		e.state.ChangePhase(GamePhaseMakeActionAI)
	}
}

func (e *GameEngine) onUnitUseAction(targetUid uint, actionResult *domain.ActionResult) {
	if actionResult != nil && actionResult.WithStun() {
		e.state.PopStunnedUnitFromQueue(targetUid)
	}
}

func (e *GameEngine) onUnitCompleteAction() {
	corpses := e.battlefield().FilterSurvivors()
	e.applyExperience(corpses)
	e.battlefield().UpdateCellsFactions()
	e.state.ShiftUnitsQueue()
	e.state.UpdateUnitsQueue(e.battlefield().Units)
	e.state.ChangePhase(GamePhaseActionComplete)
}

func (e *GameEngine) accumulateBooty(event *GameEvent) {
	index := e.rndGen.PickIndex(len(e.scenario.CurrentSpot().Booty))
	booty := e.scenario.CurrentSpot().Booty[index]
	e.state.Booty.Accumulate(booty)
	event.EndRoundResult.Booty = booty
}

func (e *GameEngine) applyExperience(corpses []*GameUnit) {
	if len(corpses) == 0 {
		return
	}
	leftUnits := e.battlefield().GetUnitsByFaction(GameUnitFactionLeft)
	if len(leftUnits) == 0 {
		return
	}
	rightCorpses := util.Filter(corpses, func(corpse *GameUnit) bool {
		return corpse.Faction == GameUnitFactionRight
	})
	totalExperience := util.Reduce(rightCorpses, 0, func(acc uint, corpse *GameUnit) uint {
		return acc + corpse.Stats.Progress.Experience
	})
	unitExperience := totalExperience / uint(len(leftUnits))
	totalExperience -= unitExperience * uint(len(leftUnits))
	for _, unit := range leftUnits {
		unit.Stats.Progress.Experience += unitExperience
		if totalExperience > 0 {
			unit.Stats.Progress.Experience += 1
			totalExperience--
		}
	}
}
