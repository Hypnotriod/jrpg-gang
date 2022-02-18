package engine

func (e *GameEngine) processAI(event *GameEvent) {
	unit := e.getActiveUnit()
	if len(unit.UserId) != 0 {
		e.processUnitAI(event, unit)
		e.state.ChangePhase(GamePhaseActionComplete)
	} else {
		e.state.ChangePhase(GamePhaseMakeMoveOrAction)
	}
}

func (e *GameEngine) processUnitAI(event *GameEvent, unit *GameUnit) {
	// todo
}
