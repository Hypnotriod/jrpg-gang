package engine

import (
	"jrpg-gang/domain"
	"jrpg-gang/util"
	"math/rand"
)

func (e *GameEngine) processAI(event *GameEvent) {
	unit := e.getActiveUnit()
	if e.aiTryToAttack(event, unit) {
		return
	}
	if e.state.phase == GamePhaseMakeMoveOrActionAI && e.aiTryToMove(event, unit) {
		return
	}
	e.onUnitCompleteAction(nil, nil)
}

func (e *GameEngine) processRetreatActionAI(event *GameEvent) {
	unit := e.getActiveUnit()
	matrix := &e.battlefield().Matrix
	position := domain.Position{}
	testColumn := func(x int) bool {
		lenY := len((*matrix)[x])
		yOffset := rand.Intn(lenY)
		position.X = x
		for y := 0; y < lenY; y++ {
			position.Y = (yOffset + y) % lenY
			if e.battlefield().CanMoveUnitTo(unit, position) {
				e.aiMove(event, unit, position)
				return true
			}
		}
		return false
	}
	if unit.Faction == GameUnitFactionLeft {
		for x := 0; x < len(*matrix); x++ {
			if testColumn(x) {
				break
			}
		}
	} else {
		for x := len(*matrix) - 1; x >= 0; x-- {
			if testColumn(x) {
				break
			}
		}
	}
	e.onUnitCompleteAction(nil, nil)
}

func (e *GameEngine) aiTryToMove(event *GameEvent, unit *GameUnit) bool {
	position := domain.Position{}
	yShift := []int{0, -1, 1, -2, 2}
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
	targets = util.Shuffle(e.rndGen, targets)
	for _, t := range targets {
		if e.aiAttackWithWeapon(event, unit, t.Target, t.WeaponUid) {
			return true
		}
	}
	return false
}

func (e *GameEngine) aiAttackWithWeapon(event *GameEvent, unit *GameUnit, target *GameUnit, weaponUid uint) bool {
	result := unit.Equip(weaponUid)
	if result.Result != domain.ResultAccomplished {
		return false
	}
	result = unit.UseInventoryItemOnTarget(&target.Unit, weaponUid, domain.NewActionResult())
	if result.Result != domain.ResultAccomplished {
		return false
	}
	e.onUseItemOnTarget(target.Uid, result)
	result = e.manageItemSpread(result, unit, target, weaponUid)
	e.onUseItemOnTarget(unit.Uid, result)
	unit.Inventory.Filter()
	e.onUnitCompleteAction(nil, nil)
	unitAction := NewGameUnitActionResult()
	unitAction.Action = domain.Action{
		Action:    domain.ActionUse,
		Uid:       unit.Uid,
		TargetUid: target.Uid,
		ItemUid:   weaponUid,
	}
	unitAction.Result = *result
	event.UnitActionResult = unitAction
	return true
}
