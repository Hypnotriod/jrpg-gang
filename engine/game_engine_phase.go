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
	case GamePhaseMakeMoveOrActionAI, GamePhaseTakeActionAI:
		e.processAI(result)
	case GamePhaseMakeMoveOrAction, GamePhaseTakeAction:
		e.onUnitCompleteAction(nil)
	case GamePhaseRetreatAction:
		e.processRetreatActionAI(result)
	case GamePhaseActionComplete:
		e.processActionComplete(result)
	case GamePhaseBeforeSpotComplete:
		e.processBeforeSpotComplete(result)
	case GamePhaseSpotComplete:
		e.processSpotComplete(result)
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
		e.state.phase == GamePhaseTakeActionAI ||
		e.state.phase == GamePhaseRetreatAction ||
		e.state.phase == GamePhaseActionComplete ||
		e.state.phase == GamePhaseBeforeSpotComplete ||
		e.state.phase == GamePhaseSpotComplete
}

func (e *GameEngine) prepareNextSpot(actors []*GameUnit) {
	e.scenario.PrepareNextSpot(actors)
	e.state.ResetWaitingOrder()
	e.state.MakeUnitsQueue(e.battlefield().Units)
	e.state.IncrementSpotNumber()
}

func (e *GameEngine) processStartRound() {
	e.state.ResetWaitingOrder()
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
	if e.isLastRound() {
		e.state.ChangePhase(GamePhaseBeforeSpotComplete)
	} else {
		e.state.ChangePhase(GamePhaseReadyForStartRound)
	}
}

func (e *GameEngine) processBeforeSpotComplete(event *GameEvent) {
	event.SpotCompleteResult = NewSpotCompleteResult()
	if e.battlefield().FactionUnitsCount(GameUnitFactionLeft) != 0 {
		e.accumulateSpotBooty(event.SpotCompleteResult)
		e.applySpotExperience(event.SpotCompleteResult)
		e.applySpotAchievements(event.SpotCompleteResult)
		e.restoreActorsState()
	}
	if e.scenario.IsLastSpot() {
		e.state.ChangePhase(GamePhaseScenarioComplete)
	} else {
		e.state.ChangePhase(GamePhaseSpotComplete)
	}
}

func (e *GameEngine) processSpotComplete(event *GameEvent) {
	e.prepareNextSpot(e.battlefield().Units)
	e.state.ChangePhase(GamePhasePrepareUnit)
	event.Spot = e.scenario.CurrentSpot()
}

func (e *GameEngine) switchToNextUnit() {
	unit := e.getActiveUnit()
	if unit.State.ActionPoints > 0 {
		if unit.HasPlayerId() {
			e.state.ChangePhase(GamePhaseMakeMoveOrAction)
		} else {
			e.state.ChangePhase(GamePhaseMakeMoveOrActionAI)
		}
		return
	}
	unit.State.ActionPoints = unit.TotalActionPoints()
	unit.State.IsStunned = false
	if unit.CheckRandomChance(unit.CalculateRetreatChance()) {
		e.state.ChangePhase(GamePhaseRetreatAction)
	} else if unit.HasPlayerId() {
		e.state.ChangePhase(GamePhaseMakeMoveOrAction)
	} else {
		e.state.ChangePhase(GamePhaseMakeMoveOrActionAI)
	}
}

func (e *GameEngine) endRound(event *GameEvent) {
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
	e.accumulateDrop(corpses, result.Drop)
	e.applyExperience(corpses, result.Experience)
	e.applyAchievements(corpses, result.Achievements)
	event.EndRoundResult = result
}

func (e *GameEngine) isLastRound() bool {
	return e.battlefield().FactionsCount() <= 1
}

func (e *GameEngine) onUnitMoveAction() {
	unit := e.getActiveUnit()
	unit.ReduceActionPoints(MOVE_ACTION_POINTS)
	if unit.State.ActionPoints < MIN_ACTION_POINTS {
		e.onUnitCompleteAction(nil)
	} else if unit.HasPlayerId() {
		e.state.ChangePhase(GamePhaseMakeMoveOrAction)
	} else {
		e.state.ChangePhase(GamePhaseMakeMoveOrActionAI)
	}
}

func (e *GameEngine) onUnitWaitAction() {
	unit := e.getActiveUnit()
	e.state.AddUnitToWaitingOrder(unit)
	e.state.UpdateUnitsQueue(e.battlefield().Units)
	e.state.ChangePhase(GamePhaseActionComplete)
}

func (e *GameEngine) onUseItemOnTarget(targetUid uint, actionResult *domain.ActionResult) {
	if actionResult != nil && actionResult.WithStun(targetUid) {
		e.state.PopStunnedUnitFromQueue(targetUid)
	}
}

func (e *GameEngine) onUnitCompleteAction(result *domain.ActionResult) {
	unit := e.getActiveUnit()
	corpses := e.battlefield().FilterSurvivors()
	if result != nil {
		e.accumulateDrop(corpses, result.Drop)
		e.applyExperience(corpses, result.Experience)
		e.applyAchievements(corpses, result.Achievements)
	} else {
		e.accumulateDrop(corpses, nil)
		e.applyExperience(corpses, nil)
		e.applyAchievements(corpses, nil)
	}
	e.battlefield().UpdateCellsFactions()
	if unit.IsDead || unit.State.ActionPoints < MIN_ACTION_POINTS {
		unit.ClearWaitingOrder()
		unit.ClearActionPoints()
		e.state.ShiftUnitsQueue()
	}
	e.state.UpdateUnitsQueue(e.battlefield().Units)
	e.state.ChangePhase(GamePhaseActionComplete)
}

func (e *GameEngine) accumulateSpotBooty(spotCompleteResult *SpotCompleteResult) {
	booty := util.RandomPick(e.rndGen, e.scenario.CurrentSpot().Booty)
	booty.W = 0
	e.state.Booty.Accumulate(booty)
	spotCompleteResult.Booty = booty
}

func (e *GameEngine) applySpotExperience(spotCompleteResult *SpotCompleteResult) {
	experience := e.scenario.CurrentSpot().Experience
	leftUnits := e.battlefield().GetUnitsByFaction(GameUnitFactionLeft)
	for _, unit := range leftUnits {
		unit.Stats.Progress.Experience += experience
		spotCompleteResult.Experience[unit.Uid] += experience
	}
}

func (e *GameEngine) applySpotAchievements(spotCompleteResult *SpotCompleteResult) {
	achievements := e.scenario.CurrentSpot().Achievements
	leftUnits := e.battlefield().GetUnitsByFaction(GameUnitFactionLeft)
	for _, unit := range leftUnits {
		unit.Achievements.Accumulate(achievements)
	}
	spotCompleteResult.Achievements.Accumulate(achievements)
}

func (e *GameEngine) accumulateDrop(corpses []*GameUnit, dropDistribution map[uint]domain.UnitBooty) {
	for n := range corpses {
		if len(corpses[n].Drop) == 0 {
			continue
		}
		drop := util.RandomPick(e.rndGen, corpses[n].Drop)
		if drop.IsEmpty() {
			continue
		}
		if dropDistribution != nil {
			drop.W = 0
			dropDistribution[corpses[n].Uid] = drop
		}
		e.state.Booty.Accumulate(drop)
	}
}

func (e *GameEngine) applyExperience(corpses []*GameUnit, expDistribution map[uint]uint) {
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
			expDistribution[unit.Uid] = unitExperience
		}
		if totalExperience > 0 {
			unit.Stats.Progress.Experience += 1
			if expDistribution != nil {
				expDistribution[unit.Uid]++
			}
			totalExperience--
		}
	}
}

func (e *GameEngine) applyAchievements(corpses []*GameUnit, achDistribution domain.UnitAchievements) {
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
	if achDistribution == nil {
		achDistribution = domain.UnitAchievements{}
	}
	for _, corpse := range rightCorpses {
		achDistribution.Merge(corpse.Achievements)
	}
	for _, unit := range leftUnits {
		unit.Achievements.Merge(achDistribution)
	}
}

func (e *GameEngine) restoreActorsState() {
	leftUnits := e.battlefield().GetUnitsByFaction(GameUnitFactionLeft)
	for i := range leftUnits {
		leftUnits[i].ClearImpact()
		leftUnits[i].State.RestoreToHalf(leftUnits[i].Stats.BaseAttributes)
	}
}
