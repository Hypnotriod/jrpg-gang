package engine

func (e *GameEngine) NextPhase(event *GameEvent) {
	switch e.state.Phase {
	case GamePhaseReadyForStartRound:
		e.processStartRound()
	case GamePhaseMakeMoveOrActionAI, GamePhaseMakeActionAI:
		e.processAI(event)
	case GamePhaseActionComplete:
		e.processActionComplete(event)
	case GamePhaseBattleComplete:
		e.processBattleComplete(event)
	}
}

func (e *GameEngine) NextPhaseRequired() bool {
	return e.state.Phase == GamePhaseReadyForStartRound ||
		e.state.Phase == GamePhaseMakeMoveOrActionAI ||
		e.state.Phase == GamePhaseMakeActionAI ||
		e.state.Phase == GamePhaseActionComplete
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
	if unit.HasUserId() {
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
	e.battlefield().FilterSurvivors()
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

func (e *GameEngine) onUnitUseAction() {
	e.state.ShiftUnitsQueue()
	e.state.UpdateUnitsQueue(e.battlefield().Units)
	e.state.ChangePhase(GamePhaseActionComplete)
}

func (e *GameEngine) onUnitPlaced() {
	if e.battlefield().ContainsUnits(e.actors) == len(e.actors) {
		e.state.ChangePhase(GamePhaseReadyForStartRound)
	}
}
