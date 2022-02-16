package engine

func (e *GameEngine) UpdatePhase() {
	if e.State.Phase == GamePhaseNone {
		e.State.ChangePhase(GamePhasePlaceUnit)
	}
	if e.State.Phase == GamePhasePlaceUnit && e.Spot.Battlefield.UnitsReady() == len(e.actors) {
		e.startRound()
	}
	if e.State.Phase == GamePhaseActionComplete {
		e.State.ShiftUnitsQueue()
		e.State.UpdateUnitsQueue(e.Spot.Battlefield.Units)
		// todo: prepare next turn
	}
}

func (e *GameEngine) startRound() {
	e.State.MakeUnitsQueue(e.Spot.Battlefield.Units)
	e.State.ChangePhase(GamePhaseMakeMoveOrAction)
}

func (e *GameEngine) endRound() EndTurnResult {
	result := EndTurnResult{}
	for _, unit := range e.Spot.Battlefield.Units {
		result.Recovery[unit.Uid] = unit.ApplyRecoverylOnNextTurn()
		result.Damage[unit.Uid] = unit.ApplyDamageOnNextTurn()
		unit.ReduceModificationOnNextTurn()
	}
	e.Spot.Battlefield.FilterSurvivors()
	return result
}

func (e *GameEngine) onUnitMoveAction() {
	e.State.ChangePhase(GamePhaseMakeAction)
}

func (e *GameEngine) onUnitUseAction() {
	e.State.ChangePhase(GamePhaseActionComplete)
}
