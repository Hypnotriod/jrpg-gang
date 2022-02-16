package engine

import "jrpg-gang/domain"

func (e *GameEngine) ExecuteAction(action GameAction, userId UserId) *domain.ActionResult {
	defer e.Unlock()
	e.Lock()
	switch action.Action {
	case GameAtionUse:
		return e.executeUseAction(action, userId)
	case GameAtionEquip:
		return e.executeEquipAction(action, userId)
	case GameAtionUnequip:
		return e.executeUnequipAction(action, userId)
	case GameAtionMove:
		return e.executeMoveAction(action, userId)
	case GameAtionPlace:
		return e.executePlaceAction(action, userId)
	}
	return domain.NewActionResult(domain.ResultNotAccomplished)
}

func (e *GameEngine) executeUseAction(action GameAction, userId UserId) *domain.ActionResult {
	if e.State.Phase != GamePhaseMakeAction && e.State.Phase != GamePhaseMakeMoveOrAction {
		return domain.NewActionResult(domain.ResultNotAllowed)
	}
	unit := e.Spot.Battlefield.FindUnitById(action.Uid)
	target := e.Spot.Battlefield.FindUnitById(action.TargetUid)
	if unit == nil || target == nil {
		return domain.NewActionResult(domain.ResultNotFound)
	}
	if !e.State.IsCurrentActiveUnit(unit) || unit.UserId != userId {
		return domain.NewActionResult(domain.ResultNotAllowed)
	}
	itemType := unit.Inventory.GetItemType(action.ItemUid)
	if unit.FractionId == target.FractionId &&
		(itemType == domain.ItemTypeWeapon || itemType == domain.ItemTypeMagic) {
		return domain.NewActionResult(domain.ResultNotAllowed)
	}
	if unit.FractionId != target.FractionId && itemType == domain.ItemTypeDisposable {
		return domain.NewActionResult(domain.ResultNotAllowed)
	}
	result := unit.UseInventoryItemOnTarget(&target.Unit, action.ItemUid)
	if result.ResultType == domain.ResultAccomplished {
		e.onUnitUseAction()
	}
	return result
}

func (e *GameEngine) executeEquipAction(action GameAction, userId UserId) *domain.ActionResult {
	if e.State.Phase != GamePhaseMakeAction && e.State.Phase != GamePhaseMakeMoveOrAction {
		return domain.NewActionResult(domain.ResultNotAllowed)
	}
	unit := e.Spot.Battlefield.FindUnitById(action.Uid)
	if unit == nil {
		return domain.NewActionResult(domain.ResultNotFound)
	}
	if !e.State.IsCurrentActiveUnit(unit) || unit.UserId != userId {
		return domain.NewActionResult(domain.ResultNotAllowed)
	}
	return unit.Equip(action.ItemUid)
}

func (e *GameEngine) executeUnequipAction(action GameAction, userId UserId) *domain.ActionResult {
	if e.State.Phase != GamePhaseMakeAction && e.State.Phase != GamePhaseMakeMoveOrAction {
		return domain.NewActionResult(domain.ResultNotAllowed)
	}
	unit := e.Spot.Battlefield.FindUnitById(action.Uid)
	if unit == nil {
		return domain.NewActionResult(domain.ResultNotFound)
	}
	if !e.State.IsCurrentActiveUnit(unit) || unit.UserId != userId {
		return domain.NewActionResult(domain.ResultNotAllowed)
	}
	return unit.Unequip(action.ItemUid)
}

func (e *GameEngine) executeMoveAction(action GameAction, userId UserId) *domain.ActionResult {
	if e.State.Phase != GamePhaseMakeMoveOrAction {
		return domain.NewActionResult(domain.ResultNotAllowed)
	}
	unit := e.Spot.Battlefield.FindUnitById(action.Uid)
	if !e.State.IsCurrentActiveUnit(unit) || unit.UserId != userId {
		return domain.NewActionResult(domain.ResultNotAllowed)
	}
	result := e.Spot.Battlefield.MoveUnit(action.Uid, action.Position)
	if result.ResultType == domain.ResultAccomplished {
		e.onUnitMoveAction()
	}
	return result
}

func (e *GameEngine) executePlaceAction(action GameAction, userId UserId) *domain.ActionResult {
	if e.State.Phase != GamePhasePlaceUnit {
		return domain.NewActionResult(domain.ResultNotAllowed)
	}
	unit := e.findActorByUserId(userId)
	if unit == nil {
		return domain.NewActionResult(domain.ResultNotAllowed)
	}
	return e.Spot.Battlefield.PlaceUnit(unit, action.Position)
}
