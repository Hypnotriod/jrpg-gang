package engine

import "jrpg-gang/domain"

func (e *GameEngine) ExecuteUserAction(action domain.Action, playerId PlayerId) *GameEvent {
	result := e.NewGameEventWithUnitAction(&action)
	var actionResult *domain.ActionResult
	switch action.Action {
	case domain.ActionPlace:
		actionResult = e.executePlaceAction(action, playerId)
	case domain.ActionUse:
		actionResult = e.executeUseAction(action, playerId)
	case domain.ActionEquip:
		actionResult = e.executeEquipAction(action, playerId)
	case domain.ActionUnequip:
		actionResult = e.executeUnequipAction(action, playerId)
	case domain.ActionThrowAway:
		actionResult = e.executeThrowAwayAction(action, playerId)
	case domain.ActionMove:
		actionResult = e.executeMoveAction(action, playerId)
	case domain.ActionSkip:
		actionResult = e.executeSkipAction(action, playerId)
	default:
		actionResult = domain.NewActionResult().WithResult(domain.ResultNotAccomplished)
	}
	result.UnitActionResult.Result = *actionResult
	result.PlayersInfo = e.GetPlayersInfo()
	result.NextPhase = e.state.phase
	return result
}

func (e *GameEngine) executePlaceAction(action domain.Action, playerId PlayerId) *domain.ActionResult {
	if e.state.phase != GamePhasePrepareUnit {
		return domain.NewActionResult().WithResult(domain.ResultNotAllowed)
	}
	unit := e.battlefield().FindUnit(action.Uid)
	if unit == nil {
		return domain.NewActionResult().WithResult(domain.ResultNotFound)
	}
	if unit.IsDead || unit.GetPlayerId() != playerId {
		return domain.NewActionResult().WithResult(domain.ResultNotAllowed)
	}
	return e.battlefield().PlaceUnit(unit, *action.Position)
}

func (e *GameEngine) executeUseAction(action domain.Action, playerId PlayerId) *domain.ActionResult {
	if !e.isActionPhase() {
		return domain.NewActionResult().WithResult(domain.ResultNotAllowed)
	}
	unit := e.battlefield().FindUnit(action.Uid)
	target := e.battlefield().FindUnit(action.TargetUid)
	if unit == nil || target == nil {
		return domain.NewActionResult().WithResult(domain.ResultNotFound)
	}
	if unit.IsDead || !e.state.IsCurrentActiveUnit(unit) || unit.GetPlayerId() != playerId {
		return domain.NewActionResult().WithResult(domain.ResultNotAllowed)
	}
	item := unit.Inventory.Find(action.ItemUid)
	if item == nil {
		return domain.NewActionResult().WithResult(domain.ResultNotFound)
	}
	switch i := item.(type) {
	case *domain.Weapon:
		if unit.Faction == target.Faction {
			return domain.NewActionResult().WithResult(domain.ResultNotAllowed)
		}
	case *domain.Magic:
		if unit.Faction == target.Faction && len(i.Damage) > 0 {
			return domain.NewActionResult().WithResult(domain.ResultNotAllowed)
		}
		if unit.Faction != target.Faction && len(i.Modification) > 0 {
			return domain.NewActionResult().WithResult(domain.ResultNotAllowed)
		}
	case *domain.Disposable:
		if unit.Faction == target.Faction && len(i.Damage) > 0 {
			return domain.NewActionResult().WithResult(domain.ResultNotAllowed)
		}
		if unit.Faction != target.Faction && len(i.Modification) > 0 {
			return domain.NewActionResult().WithResult(domain.ResultNotAllowed)
		}
	default:
		return domain.NewActionResult().WithResult(domain.ResultNotAllowed)
	}
	result := unit.UseInventoryItemOnTarget(&target.Unit, action.ItemUid)
	action.TargetUid = e.clarifyUseActionTargetuid(unit.Uid, target.Uid, result)
	if result.Result == domain.ResultAccomplished {
		e.onUnitUseAction(action.TargetUid, result)
		e.onUnitCompleteAction(&result.Experience)
	}
	return result
}

func (e *GameEngine) executeEquipAction(action domain.Action, playerId PlayerId) *domain.ActionResult {
	if !e.isActionPhase() && e.state.phase != GamePhasePrepareUnit {
		return domain.NewActionResult().WithResult(domain.ResultNotAllowed)
	}
	unit := e.battlefield().FindUnit(action.Uid)
	if unit == nil {
		return domain.NewActionResult().WithResult(domain.ResultNotFound)
	}
	if unit.IsDead || unit.GetPlayerId() != playerId ||
		e.state.phase != GamePhasePrepareUnit && !e.state.IsCurrentActiveUnit(unit) {
		return domain.NewActionResult().WithResult(domain.ResultNotAllowed)
	}
	result := unit.Equip(action.ItemUid)
	if e.state.phase == GamePhasePrepareUnit {
		e.state.UpdateUnitsQueue(e.battlefield().Units)
	}
	return result
}

func (e *GameEngine) executeUnequipAction(action domain.Action, playerId PlayerId) *domain.ActionResult {
	if !e.isActionPhase() && e.state.phase != GamePhasePrepareUnit {
		return domain.NewActionResult().WithResult(domain.ResultNotAllowed)
	}
	unit := e.battlefield().FindUnit(action.Uid)
	if unit == nil {
		return domain.NewActionResult().WithResult(domain.ResultNotFound)
	}
	if unit.IsDead || unit.GetPlayerId() != playerId ||
		e.state.phase != GamePhasePrepareUnit && !e.state.IsCurrentActiveUnit(unit) {
		return domain.NewActionResult().WithResult(domain.ResultNotAllowed)
	}
	result := unit.Unequip(action.ItemUid)
	if e.state.phase == GamePhasePrepareUnit {
		e.state.UpdateUnitsQueue(e.battlefield().Units)
	}
	return result
}

func (e *GameEngine) executeThrowAwayAction(action domain.Action, playerId PlayerId) *domain.ActionResult {
	if !e.isActionPhase() && e.state.phase != GamePhasePrepareUnit {
		return domain.NewActionResult().WithResult(domain.ResultNotAllowed)
	}
	unit := e.battlefield().FindUnit(action.Uid)
	if unit == nil {
		return domain.NewActionResult().WithResult(domain.ResultNotFound)
	}
	if unit.IsDead || unit.GetPlayerId() != playerId ||
		e.state.phase != GamePhasePrepareUnit && !e.state.IsCurrentActiveUnit(unit) {
		return domain.NewActionResult().WithResult(domain.ResultNotAllowed)
	}
	if action.Quantity == 0 {
		action.Quantity = 1
	}
	result := unit.ThrowAway(action.ItemUid, action.Quantity)
	if e.state.phase == GamePhasePrepareUnit {
		e.state.UpdateUnitsQueue(e.battlefield().Units)
	}
	return result
}

func (e *GameEngine) executeMoveAction(action domain.Action, playerId PlayerId) *domain.ActionResult {
	if e.state.phase != GamePhaseMakeMoveOrAction {
		return domain.NewActionResult().WithResult(domain.ResultNotAllowed)
	}
	unit := e.battlefield().FindUnit(action.Uid)
	if unit == nil {
		return domain.NewActionResult().WithResult(domain.ResultNotFound)
	}
	if unit.IsDead || unit.GetPlayerId() != playerId || !e.state.IsCurrentActiveUnit(unit) {
		return domain.NewActionResult().WithResult(domain.ResultNotAllowed)
	}
	result := e.battlefield().MoveUnit(unit.Uid, *action.Position)
	if result.Result == domain.ResultAccomplished {
		e.onUnitMoveAction()
	}
	return result
}

func (e *GameEngine) executeSkipAction(action domain.Action, playerId PlayerId) *domain.ActionResult {
	if !e.isActionPhase() {
		return domain.NewActionResult().WithResult(domain.ResultNotAllowed)
	}
	unit := e.battlefield().FindUnit(action.Uid)
	if unit == nil {
		return domain.NewActionResult().WithResult(domain.ResultNotFound)
	}
	if unit.IsDead || unit.GetPlayerId() != playerId || !e.state.IsCurrentActiveUnit(unit) {
		return domain.NewActionResult().WithResult(domain.ResultNotAllowed)
	}
	result := domain.NewActionResult()
	e.onUnitCompleteAction(&result.Experience)
	return result.WithResult(domain.ResultAccomplished)
}

func (e *GameEngine) isActionPhase() bool {
	return e.state.phase == GamePhaseMakeAction || e.state.phase == GamePhaseMakeMoveOrAction
}

func (e *GameEngine) canTakeAShare() bool {
	return e.state.phase == GamePhaseSpotComplete || e.state.phase == GamePhaseScenarioComplete
}
