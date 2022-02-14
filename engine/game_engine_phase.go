package engine

func (e *GameEngine) UpdatePhase() {
	if e.State.Phase == GamePhaseNone {
		e.State.ChangePhase(GamePhasePlaceUnit)
	}
	if e.State.Phase == GamePhasePlaceUnit && e.Spot.Battlefield.UnitsReady() == len(e.actors) {
		e.startRound()
	}
}

func (e *GameEngine) startRound() {
	e.State.MakeUnitsQueue(e.Spot.Battlefield.Units)
	e.State.ChangePhase(GamePhaseMakeMoveOrAction)
}

func (e *GameEngine) onUnitMove() {
	e.State.ChangePhase(GamePhaseMakeAction)
}

func (e *GameEngine) onUnitAction() {
	e.State.ShiftUnitsQueue()
	e.State.UpdateUnitsQueue(e.Spot.Battlefield.Units)
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
