package engine

import (
	"jrpg-gang/domain"
	"jrpg-gang/util"
	"math/rand"
)

func (e *GameEngine) processAI(event *GameEvent) {
	unit := e.getActiveUnit()
	targets := e.battlefield().FindReachableTargets(unit, true, false)
	if len(targets) != 0 && e.aiTryToAttack(event, unit, targets) {
		return
	}
	if e.state.phase == GamePhaseMakeMoveOrActionAI &&
		len(targets) == 0 && e.aiTryToApproachTheEnemy(event, unit) {
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

func (e *GameEngine) aiTryToApproachTheEnemy(event *GameEvent, unit *GameUnit) bool {
	xShift := 0
	xShiftMin := 1
	yShift := []int{0, -1, 1, -2, 2}
	position := domain.Position{}
	for i := range unit.Inventory.Weapon {
		weapon := &unit.Inventory.Weapon[i]
		if !unit.CanUseWeapon(weapon, true) {
			continue
		}
		if xShift < weapon.Range.MaximumX {
			xShift = weapon.Range.MaximumX
			if weapon.Range.MinimumX > 1 {
				xShiftMin = weapon.Range.MinimumX - 1
			} else {
				xShiftMin = 1
			}
		}
	}
	targets := util.Shuffle(e.rndGen, util.Clone(e.battlefield().Units))
	for _, target := range targets {
		if target.Faction == unit.Faction {
			continue
		}
		x := xShift
		for x >= xShiftMin {
			if target.Faction == GameUnitFactionLeft {
				position.X = target.Position.X + x
			} else {
				position.X = target.Position.X - x
			}
			for _, y := range yShift {
				position.Y = target.Position.Y + y
				if e.battlefield().CanMoveUnitTo(unit, position) {
					e.aiMove(event, unit, position)
					return true
				}
			}
			x--
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

func (e *GameEngine) aiTryToAttack(event *GameEvent, unit *GameUnit, targets []ReachableTarget) bool {
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
	unit.Inventory.UpdateItemsState()
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
