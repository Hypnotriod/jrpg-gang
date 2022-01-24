package engine

import "jrpg-gang/domain"

type GameEngine struct {
	Battlefield *Battlefield
}

func NewGameEngine(battlefield *Battlefield) *GameEngine {
	engine := &GameEngine{}
	engine.Battlefield = battlefield
	return engine
}

func (e *GameEngine) ExecuteAction(action GameAction) domain.ActionResult {
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
