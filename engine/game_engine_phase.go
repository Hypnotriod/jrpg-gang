package engine

func (e *GameEngine) NextPhase(event *GameEvent) {
	if e.state.Phase == GamePhaseNone {
		e.state.ChangePhase(GamePhasePlaceUnit)
	} else if e.state.Phase == GamePhasePlaceUnit &&
		e.spot.Battlefield.ContainsUnits(e.actors) == len(e.actors) {
		e.startRound()
	} else if e.state.Phase == GamePhaseProcessAI {
		e.processAI(event)
	} else if e.state.Phase == GamePhaseActionComplete {
		e.processActionComplete(event)
	}
}

func (e *GameEngine) processActionComplete(event *GameEvent) {
	e.state.ShiftUnitsQueue()
	e.state.UpdateUnitsQueue(e.spot.Battlefield.Units)
	if !e.state.HasActiveUnits() {
		e.processRoundComplete(event)
	} else {
		e.processNextTurn()
	}
}

func (e *GameEngine) processRoundComplete(event *GameEvent) {
	e.endRound(event)
	if e.spot.Battlefield.FactionsCount() <= 1 {
		e.processBattleComplete()
	} else {
		e.startRound()
	}
}

func (e *GameEngine) processBattleComplete() {
	e.state.ChangePhase(GamePhaseBattleComplete)
}

func (e *GameEngine) processNextTurn() {
	e.state.ChangePhase(GamePhaseProcessAI)
}

func (e *GameEngine) startRound() {
	e.state.MakeUnitsQueue(e.spot.Battlefield.Units)
	e.state.ChangePhase(GamePhaseProcessAI)
}

func (e *GameEngine) endRound(event *GameEvent) {
	result := &EndTurnResult{}
	for _, unit := range e.spot.Battlefield.Units {
		result.Recovery[unit.Uid] = unit.ApplyRecoverylOnNextTurn()
		result.Damage[unit.Uid] = unit.ApplyDamageOnNextTurn()
		unit.ReduceModificationOnNextTurn()
	}
	e.spot.Battlefield.FilterSurvivors()
	event.EndRoundResult = result
}

func (e *GameEngine) onUnitMoveAction() {
	e.state.ChangePhase(GamePhaseMakeAction)
}

func (e *GameEngine) onUnitUseAction() {
	e.state.ChangePhase(GamePhaseActionComplete)
}
