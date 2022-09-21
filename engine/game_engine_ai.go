package engine

import "jrpg-gang/domain"

func (e *GameEngine) processAI(event *GameEvent) {
	unit := e.getActiveUnit()
	if e.aiTryToAttack(event, unit) {
		return
	}
	if e.state.phase == GamePhaseMakeMoveOrActionAI && e.aiTryToMove(event, unit) {
		return
	}
	e.onUnitCompleteAction()
}

func (e *GameEngine) aiTryToMove(event *GameEvent, unit *GameUnit) bool {
	position := domain.Position{}
	yShift := []int{0, -1, 1}
	for _, target := range e.battlefield().Units {
		if target.Faction == unit.Faction {
			continue
		}
		if target.Faction == GameUnitFactionLeft {
			position.X = target.Position.X + 1
		} else {
			position.X = target.Position.X - 1
		}
		for _, y := range yShift {
			position.Y = target.Position.Y + y
			if e.battlefield().CanMoveUnitTo(unit, position) {
				e.aiMove(event, unit, position)
				return true
			}
		}
	}
	return false
}

func (e *GameEngine) aiMove(event *GameEvent, unit *GameUnit, position domain.Position) {
	unitAction := NewGameUnitActionResult()
	unitAction.Action = domain.Action{
		Action:   domain.ActionMove,
		Uid:      unit.Uid,
		Position: &position,
	}
	unitAction.Result = *e.battlefield().MoveUnit(unit.Uid, position)
	event.UnitActionResult = unitAction
	e.onUnitMoveAction()
}

func (e *GameEngine) aiTryToAttack(event *GameEvent, unit *GameUnit) bool {
	targets := e.battlefield().FindReachableTargets(unit)
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
	unitAction := NewGameUnitActionResult()
	unitAction.Action = domain.Action{
		Action:    domain.ActionUse,
		Uid:       unit.Uid,
		TargetUid: target.Uid,
		ItemUid:   weaponUid,
	}
	result := unit.UseInventoryItemOnTarget(&target.Unit, weaponUid)
	unitAction.Result = *result
	event.UnitActionResult = unitAction
	e.onUnitUseAction(target.Uid, result)
	e.onUnitCompleteAction()
}
