package engine

func (e *GameEngine) StartRound() {
	e.State.MakeUnitsQueue(e.Battlefield.Units)
}

func (e *GameEngine) EndRound() EndTurnResult {
	result := EndTurnResult{}
	for _, unit := range e.Battlefield.Units {
		result.Recovery[unit.Uid] = unit.ApplyRecoverylOnNextTurn()
		result.Damage[unit.Uid] = unit.ApplyDamageOnNextTurn()
		unit.ReduceModificationOnNextTurn()
	}
	e.Battlefield.FilterSurvivors()
	return result
}

func (e *GameEngine) PrepareNextTurn() {
	e.State.UpdateUnitsQueue(e.Battlefield.Units)
}
