package engine

func (e *GameEngine) processAI() {
	for {
		unit := e.getActiveUnit()
		if unit == nil || len(unit.UserId) == 0 {
			return
		}
		e.processUnitAI(unit)
	}
}

func (e *GameEngine) processUnitAI(unit *GameUnit) {
	// todo
}
