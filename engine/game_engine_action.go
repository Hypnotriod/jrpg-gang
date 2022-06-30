package engine

import "jrpg-gang/domain"

func (e *GameEngine) ExecuteUserAction(action domain.Action, userId UserId) *GameEvent {
	result := e.NewGameEventWithUnitAction(&action)
	var actionResult *domain.ActionResult
	switch action.Action {
	case domain.ActionPlace:
		actionResult = e.executePlaceAction(action, userId)
	case domain.ActionUse:
		actionResult = e.executeUseAction(action, userId)
	case domain.ActionEquip:
		actionResult = e.executeEquipAction(action, userId)
	case domain.ActionUnequip:
		actionResult = e.executeUnequipAction(action, userId)
	case domain.ActionMove:
		actionResult = e.executeMoveAction(action, userId)
	default:
		actionResult = domain.NewActionResult().WithResult(domain.ResultNotAccomplished)
	}
	result.UnitActionResult.Result = *actionResult
	result.NextPhase = e.state.phase
	return result
}

func (e *GameEngine) executePlaceAction(action domain.Action, userId UserId) *domain.ActionResult {
	if e.state.phase != GamePhasePlaceUnit {
		return domain.NewActionResult().WithResult(domain.ResultNotAllowed)
	}
	unit := e.FindActorByUserId(userId)
	if unit == nil {
		return domain.NewActionResult().WithResult(domain.ResultNotAllowed)
	}
	result := e.battlefield().PlaceUnit(unit, *action.Position)
	if result.Result == domain.ResultAccomplished {
		e.onUnitPlaced()
	}
	return result
}

func (e *GameEngine) executeUseAction(action domain.Action, userId UserId) *domain.ActionResult {
	if e.state.phase != GamePhaseMakeAction && e.state.phase != GamePhaseMakeMoveOrAction {
		return domain.NewActionResult().WithResult(domain.ResultNotAllowed)
	}
	unit := e.battlefield().FindUnitById(action.Uid)
	target := e.battlefield().FindUnitById(action.TargetUid)
	if unit == nil || target == nil {
		return domain.NewActionResult().WithResult(domain.ResultNotFound)
	}
	if !e.state.IsCurrentActiveUnit(unit) || unit.userId != userId {
		return domain.NewActionResult().WithResult(domain.ResultNotAllowed)
	}
	itemType := unit.Inventory.GetItemType(action.ItemUid)
	if itemType == domain.ItemTypeNone {
		return domain.NewActionResult().WithResult(domain.ResultNotFound)
	}
	if unit.Faction == target.Faction &&
		(itemType == domain.ItemTypeWeapon || itemType == domain.ItemTypeMagic) {
		return domain.NewActionResult().WithResult(domain.ResultNotAllowed)
	}
	if unit.Faction != target.Faction && itemType == domain.ItemTypeDisposable {
		return domain.NewActionResult().WithResult(domain.ResultNotAllowed)
	}
	result := unit.UseInventoryItemOnTarget(&target.Unit, action.ItemUid)
	if result.Result == domain.ResultAccomplished {
		e.onUnitUseAction()
	}
	return result
}

func (e *GameEngine) executeEquipAction(action domain.Action, userId UserId) *domain.ActionResult {
	if e.state.phase != GamePhaseMakeAction && e.state.phase != GamePhaseMakeMoveOrAction {
		return domain.NewActionResult().WithResult(domain.ResultNotAllowed)
	}
	unit := e.battlefield().FindUnitById(action.Uid)
	if unit == nil {
		return domain.NewActionResult().WithResult(domain.ResultNotFound)
	}
	if !e.state.IsCurrentActiveUnit(unit) || unit.userId != userId {
		return domain.NewActionResult().WithResult(domain.ResultNotAllowed)
	}
	return unit.Equip(action.ItemUid)
}

func (e *GameEngine) executeUnequipAction(action domain.Action, userId UserId) *domain.ActionResult {
	if e.state.phase != GamePhaseMakeAction && e.state.phase != GamePhaseMakeMoveOrAction {
		return domain.NewActionResult().WithResult(domain.ResultNotAllowed)
	}
	unit := e.battlefield().FindUnitById(action.Uid)
	if unit == nil {
		return domain.NewActionResult().WithResult(domain.ResultNotFound)
	}
	if !e.state.IsCurrentActiveUnit(unit) || unit.userId != userId {
		return domain.NewActionResult().WithResult(domain.ResultNotAllowed)
	}
	return unit.Unequip(action.ItemUid)
}

func (e *GameEngine) executeMoveAction(action domain.Action, userId UserId) *domain.ActionResult {
	if e.state.phase != GamePhaseMakeMoveOrAction {
		return domain.NewActionResult().WithResult(domain.ResultNotAllowed)
	}
	unit := e.battlefield().FindUnitById(action.Uid)
	if !e.state.IsCurrentActiveUnit(unit) || unit.userId != userId {
		return domain.NewActionResult().WithResult(domain.ResultNotAllowed)
	}
	result := e.battlefield().MoveUnit(action.Uid, *action.Position)
	if result.Result == domain.ResultAccomplished {
		e.onUnitMoveAction()
	}
	return result
}
