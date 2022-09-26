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
	result.NextPhase = e.state.phase
	return result
}

func (e *GameEngine) NextPhaseRequired() bool {
	return e.state.phase == GamePhaseReadyForStartRound ||
		e.state.phase == GamePhasePrepareUnit ||
		e.state.phase == GamePhaseMakeMoveOrActionAI ||
		e.state.phase == GamePhaseMakeActionAI ||
		e.state.phase == GamePhaseRetreatAction ||
		e.state.phase == GamePhaseActionComplete
}

func (e *GameEngine) prepareNextSpot() {
	e.scenario.PrepareNextSpot(e.actors)
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
	e.endRound(event)
	if e.battlefield().FactionsCount() <= 1 {
		e.state.ChangePhase(GamePhaseBattleComplete)
	} else {
		e.state.ChangePhase(GamePhaseReadyForStartRound)
	}
}

func (e *GameEngine) processBattleComplete(event *GameEvent) {
	// todo
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

func (e *GameEngine) endRound(event *GameEvent) {
	result := NewEndTurnResult()
	for _, unit := range e.battlefield().Units {
		result.Recovery[unit.Uid] = unit.ApplyRecoverylOnNextTurn()
		result.Damage[unit.Uid] = unit.ApplyDamageOnNextTurn()
		unit.ReduceModificationOnNextTurn()
	}
	corpses := e.battlefield().FilterSurvivors()
	e.applyExperience(corpses)
	event.EndRoundResult = result
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

func (e *GameEngine) applyExperience(corpses []*GameUnit) {
	if len(corpses) == 0 {
		return
	}
	leftUnits := e.battlefield().GetUnitsByFraction(GameUnitFactionLeft)
	if len(leftUnits) == 0 {
		return
	}
	rightCorpses := util.Filter(corpses, func(corpse *GameUnit) bool {
		return corpse.Faction == GameUnitFactionRight
	})
	totalExperience := util.Reduce(rightCorpses, 0, func(acc uint, corpse *GameUnit) uint {
		return acc + corpse.Stats.Progress.Experience
	})
	experience := totalExperience / uint(len(leftUnits))
	for _, unit := range leftUnits {
		unit.Stats.Progress.Experience += experience
	}
	leftUnits[0].Stats.Progress.Experience += totalExperience % uint(len(leftUnits))
}
