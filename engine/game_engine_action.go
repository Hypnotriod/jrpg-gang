package engine

import "jrpg-gang/domain"

func (e *GameEngine) ExecuteAction(action GameAction, userId UserId) *domain.ActionResult {
	switch action.Action {
	case GameAtionPlace:
		return e.executePlaceAction(action, userId)
	case GameAtionUse:
		return e.executeUseAction(action, userId)
	case GameAtionEquip:
		return e.executeEquipAction(action, userId)
	case GameAtionUnequip:
		return e.executeUnequipAction(action, userId)
	case GameAtionMove:
		return e.executeMoveAction(action, userId)
	}
	return domain.NewActionResult(domain.ResultNotAccomplished)
}

func (e *GameEngine) executePlaceAction(action GameAction, userId UserId) *domain.ActionResult {
	if e.state.Phase != GamePhasePlaceUnit {
		return domain.NewActionResult(domain.ResultNotAllowed)
	}
	unit := e.findActorByUserId(userId)
	if unit == nil {
		return domain.NewActionResult(domain.ResultNotAllowed)
	}
	result := e.battlefield().PlaceUnit(unit, action.Position)
	if result.ResultType == domain.ResultAccomplished {
		e.onUnitPlaced()
	}
	return result
}

func (e *GameEngine) executeUseAction(action GameAction, userId UserId) *domain.ActionResult {
	if e.state.Phase != GamePhaseMakeAction && e.state.Phase != GamePhaseMakeMoveOrAction {
		return domain.NewActionResult(domain.ResultNotAllowed)
	}
	unit := e.battlefield().FindUnitById(action.Uid)
	target := e.battlefield().FindUnitById(action.TargetUid)
	if unit == nil || target == nil {
		return domain.NewActionResult(domain.ResultNotFound)
	}
	if !e.state.IsCurrentActiveUnit(unit) || unit.UserId != userId {
		return domain.NewActionResult(domain.ResultNotAllowed)
	}
	itemType := unit.Inventory.GetItemType(action.ItemUid)
	if unit.Faction == target.Faction &&
		(itemType == domain.ItemTypeWeapon || itemType == domain.ItemTypeMagic) {
		return domain.NewActionResult(domain.ResultNotAllowed)
	}
	if unit.Faction != target.Faction && itemType == domain.ItemTypeDisposable {
		return domain.NewActionResult(domain.ResultNotAllowed)
	}
	result := unit.UseInventoryItemOnTarget(&target.Unit, action.ItemUid)
	if result.ResultType == domain.ResultAccomplished {
		e.onUnitUseAction()
	}
	return result
}

func (e *GameEngine) executeEquipAction(action GameAction, userId UserId) *domain.ActionResult {
	if e.state.Phase != GamePhaseMakeAction && e.state.Phase != GamePhaseMakeMoveOrAction {
		return domain.NewActionResult(domain.ResultNotAllowed)
	}
	unit := e.battlefield().FindUnitById(action.Uid)
	if unit == nil {
		return domain.NewActionResult(domain.ResultNotFound)
	}
	if !e.state.IsCurrentActiveUnit(unit) || unit.UserId != userId {
		return domain.NewActionResult(domain.ResultNotAllowed)
	}
	return unit.Equip(action.ItemUid)
}

func (e *GameEngine) executeUnequipAction(action GameAction, userId UserId) *domain.ActionResult {
	if e.state.Phase != GamePhaseMakeAction && e.state.Phase != GamePhaseMakeMoveOrAction {
		return domain.NewActionResult(domain.ResultNotAllowed)
	}
	unit := e.battlefield().FindUnitById(action.Uid)
	if unit == nil {
		return domain.NewActionResult(domain.ResultNotFound)
	}
	if !e.state.IsCurrentActiveUnit(unit) || unit.UserId != userId {
		return domain.NewActionResult(domain.ResultNotAllowed)
	}
	return unit.Unequip(action.ItemUid)
}

func (e *GameEngine) executeMoveAction(action GameAction, userId UserId) *domain.ActionResult {
	if e.state.Phase != GamePhaseMakeMoveOrAction {
		return domain.NewActionResult(domain.ResultNotAllowed)
	}
	unit := e.battlefield().FindUnitById(action.Uid)
	if !e.state.IsCurrentActiveUnit(unit) || unit.UserId != userId {
		return domain.NewActionResult(domain.ResultNotAllowed)
	}
	result := e.battlefield().MoveUnit(action.Uid, action.Position)
	if result.ResultType == domain.ResultAccomplished {
		e.onUnitMoveAction()
	}
	return result
}
