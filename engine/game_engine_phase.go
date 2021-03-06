package engine

func (e *GameEngine) NextPhase() *GameEvent {
	result := e.NewGameEvent()
	switch e.state.phase {
	case GamePhaseReadyForStartRound:
		e.processStartRound()
	case GamePhaseMakeMoveOrActionAI, GamePhaseMakeActionAI:
		e.processAI(result)
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
		e.state.phase == GamePhaseMakeMoveOrActionAI ||
		e.state.phase == GamePhaseMakeActionAI ||
		e.state.phase == GamePhaseActionComplete
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
