package engine

func (e *GameEngine) EndTurn() EndTurnResult {
	result := EndTurnResult{}
	for _, unit := range e.Battlefield.Units {
		result.Recovery[unit.Uid] = unit.ApplyRecoverylOnNextTurn()
		result.Damage[unit.Uid] = unit.ApplyDamageOnNextTurn()
		unit.ReduceModificationOnNextTurn()
	}
	e.Battlefield.FilterSurvivors()
	e.Battlefield.SortUnitsByInitiative()
	return result
}
