package engine

import "jrpg-gang/domain"

type UnitConfigurator struct {
}

func NewUnitConfigurator() *UnitConfigurator {
	c := &UnitConfigurator{}
	return c
}

func (c *UnitConfigurator) ExecuteAction(action domain.Action, unit *domain.Unit) *domain.ActionResult {
	switch action.Action {
	case domain.ActionEquip:
		return unit.Equip(action.ItemUid)
	case domain.ActionUnequip:
		return unit.Unequip(action.ItemUid)
	case domain.ActionThrowAway:
		//todo
	case domain.ActionLevelUp:
		return unit.LevelUp()
	case domain.ActionSkillUp:
		return unit.SkillUp(action.Property)
	}
	return domain.NewActionResult().WithResult(domain.ResultNotAccomplished)
}
