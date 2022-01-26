package engine

import "jrpg-gang/domain"

func (e *GameEngine) ExecuteAction(action GameAction) domain.ActionResult {
	defer e.Unlock()
	e.Lock()
	switch action.Action {
	case GameAtionUse:
		return e.executeUseAction(action)
	case GameAtionEquip:
		return e.executeEquipAction(action)
	case GameAtionUnequip:
		return e.executeUnequipAction(action)
	case GameAtionMove:
		return e.executeMoveAction(action)
	}
	return *domain.NewActionResult(domain.ResultNotAccomplished)
}

func (e *GameEngine) executeUseAction(action GameAction) domain.ActionResult {
	unit := e.Battlefield.FindUnitById(action.Uid)
	target := e.Battlefield.FindUnitById(action.TargetUid)
	if unit == nil || target == nil {
		return *domain.NewActionResult(domain.ResultNotFound)
	}
	itemType := unit.Inventory.GetItemType(action.ItemUid)
	if unit.FractionId == target.FractionId &&
		(itemType == domain.ItemTypeWeapon || itemType == domain.ItemTypeMagic) {
		return *domain.NewActionResult(domain.ResultNotAllowed)
	}
	if unit.FractionId != target.FractionId && itemType == domain.ItemTypeDisposable {
		return *domain.NewActionResult(domain.ResultNotAllowed)
	}
	return unit.UseInventoryItemOnTarget(&target.Unit, action.ItemUid)
}

func (e *GameEngine) executeEquipAction(action GameAction) domain.ActionResult {
	unit := e.Battlefield.FindUnitById(action.Uid)
	if unit == nil {
		return *domain.NewActionResult(domain.ResultNotFound)
	}
	return unit.Equip(action.ItemUid)
}

func (e *GameEngine) executeUnequipAction(action GameAction) domain.ActionResult {
	unit := e.Battlefield.FindUnitById(action.Uid)
	if unit == nil {
		return *domain.NewActionResult(domain.ResultNotFound)
	}
	return unit.Unequip(action.ItemUid)
}

func (e *GameEngine) executeMoveAction(action GameAction) domain.ActionResult {
	return e.Battlefield.MoveUnit(action.Uid, action.Position)
}
