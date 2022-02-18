package engine

func (e *GameEngine) processAI(event *GameEvent) {
	unit := e.getActiveUnit()
	if len(unit.UserId) != 0 {
		e.aiProcessUnit(event, unit)
		e.state.ChangePhase(GamePhaseActionComplete)
	} else {
		e.state.ChangePhase(GamePhaseMakeMoveOrAction)
	}
}

func (e *GameEngine) aiProcessUnit(event *GameEvent, unit *GameUnit) {
	if e.aiTryToAttack(event, unit) {
		return
	}
	if e.aiTryToMove(event, unit) {
		e.aiTryToAttack(event, unit)
		return
	}
}

func (e *GameEngine) aiTryToMove(event *GameEvent, unit *GameUnit) bool {
	// todo
	return false
}

func (e *GameEngine) aiTryToAttack(event *GameEvent, unit *GameUnit) bool {
	targets := e.spot.Battlefield.FindReachableTargets(unit)
	if len(targets) == 0 {
		return false
	}
	index := e.rndGen.PickIndex(len(targets))
	cnt := 0
	for weaponUid, target := range targets {
		if cnt == index {
			e.aiAttackWithWeapon(event, unit, target, weaponUid)
			return true
		}
		cnt++
	}
	return true
}

func (e *GameEngine) aiAttackWithWeapon(event *GameEvent, unit *GameUnit, target *GameUnit, weaponUid uint) {
	unit.Equip(weaponUid)
	unitAction := &GameUnitActionResult{}
	unitAction.Action = GameAction{
		Action:    GameAtionUse,
		Uid:       unit.Uid,
		TargetUid: target.Uid,
		ItemUid:   weaponUid,
	}
	unitAction.Result = *unit.UseInventoryItemOnTarget(&target.Unit, weaponUid)
	event.UnitActionResult = unitAction
}
