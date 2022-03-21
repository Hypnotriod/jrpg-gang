package engine

import "jrpg-gang/domain"

func (e *GameEngine) ExecuteAction(action domain.Action, userId UserId) *domain.ActionResult {
	switch action.Action {
	case domain.ActionPlace:
		return e.executePlaceAction(action, userId)
	case domain.ActionUse:
		return e.executeUseAction(action, userId)
	case domain.ActionEquip:
		return e.executeEquipAction(action, userId)
	case domain.ActionUnequip:
		return e.executeUnequipAction(action, userId)
	case domain.ActionMove:
		return e.executeMoveAction(action, userId)
	}
	return domain.NewActionResult().WithResultType(domain.ResultNotAccomplished)
}

func (e *GameEngine) executePlaceAction(action domain.Action, userId UserId) *domain.ActionResult {
	if e.state.Phase != GamePhasePlaceUnit {
		return domain.NewActionResult().WithResultType(domain.ResultNotAllowed)
	}
	unit := e.findActorByUserId(userId)
	if unit == nil {
		return domain.NewActionResult().WithResultType(domain.ResultNotAllowed)
	}
	result := e.battlefield().PlaceUnit(unit, action.Position)
	if result.ResultType == domain.ResultAccomplished {
		e.onUnitPlaced()
	}
	return result
}

func (e *GameEngine) executeUseAction(action domain.Action, userId UserId) *domain.ActionResult {
	if e.state.Phase != GamePhaseMakeAction && e.state.Phase != GamePhaseMakeMoveOrAction {
		return domain.NewActionResult().WithResultType(domain.ResultNotAllowed)
	}
	unit := e.battlefield().FindUnitById(action.Uid)
	target := e.battlefield().FindUnitById(action.TargetUid)
	if unit == nil || target == nil {
		return domain.NewActionResult().WithResultType(domain.ResultNotFound)
	}
	if !e.state.IsCurrentActiveUnit(unit) || unit.UserId != userId {
		return domain.NewActionResult().WithResultType(domain.ResultNotAllowed)
	}
	itemType := unit.Inventory.GetItemType(action.ItemUid)
	if unit.Faction == target.Faction &&
		(itemType == domain.ItemTypeWeapon || itemType == domain.ItemTypeMagic) {
		return domain.NewActionResult().WithResultType(domain.ResultNotAllowed)
	}
	if unit.Faction != target.Faction && itemType == domain.ItemTypeDisposable {
		return domain.NewActionResult().WithResultType(domain.ResultNotAllowed)
	}
	result := unit.UseInventoryItemOnTarget(&target.Unit, action.ItemUid)
	if result.ResultType == domain.ResultAccomplished {
		e.onUnitUseAction()
	}
	return result
}

func (e *GameEngine) executeEquipAction(action domain.Action, userId UserId) *domain.ActionResult {
	if e.state.Phase != GamePhaseMakeAction && e.state.Phase != GamePhaseMakeMoveOrAction {
		return domain.NewActionResult().WithResultType(domain.ResultNotAllowed)
	}
	unit := e.battlefield().FindUnitById(action.Uid)
	if unit == nil {
		return domain.NewActionResult().WithResultType(domain.ResultNotFound)
	}
	if !e.state.IsCurrentActiveUnit(unit) || unit.UserId != userId {
		return domain.NewActionResult().WithResultType(domain.ResultNotAllowed)
	}
	return unit.Equip(action.ItemUid)
}

func (e *GameEngine) executeUnequipAction(action domain.Action, userId UserId) *domain.ActionResult {
	if e.state.Phase != GamePhaseMakeAction && e.state.Phase != GamePhaseMakeMoveOrAction {
		return domain.NewActionResult().WithResultType(domain.ResultNotAllowed)
	}
	unit := e.battlefield().FindUnitById(action.Uid)
	if unit == nil {
		return domain.NewActionResult().WithResultType(domain.ResultNotFound)
	}
	if !e.state.IsCurrentActiveUnit(unit) || unit.UserId != userId {
		return domain.NewActionResult().WithResultType(domain.ResultNotAllowed)
	}
	return unit.Unequip(action.ItemUid)
}

func (e *GameEngine) executeMoveAction(action domain.Action, userId UserId) *domain.ActionResult {
	if e.state.Phase != GamePhaseMakeMoveOrAction {
		return domain.NewActionResult().WithResultType(domain.ResultNotAllowed)
	}
	unit := e.battlefield().FindUnitById(action.Uid)
	if !e.state.IsCurrentActiveUnit(unit) || unit.UserId != userId {
		return domain.NewActionResult().WithResultType(domain.ResultNotAllowed)
	}
	result := e.battlefield().MoveUnit(action.Uid, action.Position)
	if result.ResultType == domain.ResultAccomplished {
		e.onUnitMoveAction()
	}
	return result
}
