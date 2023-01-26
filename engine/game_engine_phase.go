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
	case GamePhaseMakeMoveOrAction, GamePhaseMakeAction:
		e.onUnitCompleteAction(nil)
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
	result := NewEndRoundResult()
	for _, unit := range e.battlefield().Units {
		recovery := unit.ApplyRecoverylOnNextTurn()
		if recovery.HasEffect() {
			result.Recovery[unit.Uid] = recovery
		}
		damage := unit.ApplyDamageOnNextTurn()
		if damage.HasEffect() {
			result.Damage[unit.Uid] = damage
		}
		unit.ReduceModificationOnNextTurn()
	}
	corpses := e.battlefield().FilterSurvivors()
	e.applyExperience(corpses, &result.Experience)
	event.EndRoundResult = result
	isLastRound = e.battlefield().FactionsCount() <= 1
	if isLastRound && e.battlefield().FactionUnitsCount(GameUnitFactionLeft) != 0 {
		e.accumulateBooty(event)
		e.applySpotExperience(event)
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

func (e *GameEngine) clarifyUseActionTargetuid(unitUit uint, targetUid uint, actionResult *domain.ActionResult) uint {
	if actionResult != nil && actionResult.IsCriticalMiss() {
		return unitUit
	}
	return targetUid
}

func (e *GameEngine) onUnitCompleteAction(expDistribution *map[uint]uint) {
	corpses := e.battlefield().FilterSurvivors()
	e.applyExperience(corpses, expDistribution)
	e.battlefield().UpdateCellsFactions()
	e.state.ShiftUnitsQueue()
	e.state.UpdateUnitsQueue(e.battlefield().Units)
	e.state.ChangePhase(GamePhaseActionComplete)
}

func (e *GameEngine) accumulateBooty(event *GameEvent) {
	booty := util.RandomPick(e.rndGen, e.scenario.CurrentSpot().Booty)
	booty.W = 0
	e.state.Booty.Accumulate(booty)
	event.EndRoundResult.Booty = &booty
}

func (e *GameEngine) applySpotExperience(event *GameEvent) {
	experience := e.scenario.CurrentSpot().Experience
	leftUnits := e.battlefield().GetUnitsByFaction(GameUnitFactionLeft)
	for _, unit := range leftUnits {
		unit.Stats.Progress.Experience += experience
		event.EndRoundResult.Experience[unit.Uid] += experience
	}
}

func (e *GameEngine) applyExperience(corpses []*GameUnit, expDistribution *map[uint]uint) {
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
		if expDistribution != nil {
			(*expDistribution)[unit.Uid] = unitExperience
		}
		if totalExperience > 0 {
			unit.Stats.Progress.Experience += 1
			if expDistribution != nil {
				(*expDistribution)[unit.Uid]++
			}
			totalExperience--
		}
	}
}
