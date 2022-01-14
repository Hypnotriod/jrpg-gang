package domain

import (
	"fmt"
	"jrpg-gang/util"
)

type UseInventoryItemActionResult struct {
	InstantDamage        []Damage                 `json:"instantDamage,omitempty"`
	TemporalDamage       []DamageImpact           `json:"temporalDamage,omitempty"`
	InstantRecovery      []UnitState              `json:"instantRecovery,omitempty"`
	TemporalModification []UnitModificationImpact `json:"temporalModification,omitempty"`
	Accomplished         bool                     `json:"accomplished"`
}

func (r UseInventoryItemActionResult) String() string {
	return fmt.Sprintf("instant damage: [%s], temporal damage: [%s], instant recovery: [%s], temporal modification: [%s], accomplished: %t",
		util.AsCommaSeparatedSlice(r.InstantDamage),
		util.AsCommaSeparatedSlice(r.TemporalDamage),
		util.AsCommaSeparatedSlice(r.InstantRecovery),
		util.AsCommaSeparatedSlice(r.TemporalModification),
		r.Accomplished)
}

func (u *Unit) UseInventoryItemOnTarget(target *Unit, uid uint) UseInventoryItemActionResult {
	action := UseInventoryItemActionResult{}
	item := u.Inventory.Find(uid)
	if item == nil {
		return action
	}
	switch v := item.(type) {
	case *Weapon:
		u.useWeaponOnTarget(&action, target, v)
	case *Disposable:
		u.useDisposableOnTarget(&action, target, v)
	case *Magic:
		u.useMagicOnTarget(&action, target, v)
	}
	return action
}

func (u *Unit) useWeaponOnTarget(action *UseInventoryItemActionResult, target *Unit, weapon *Weapon) {
	if !weapon.Equipped || !u.CheckUseCost(weapon.UseCost) {
		return
	}
	var damage []DamageImpact = weapon.Damage
	if weapon.RequiresAmmunition() {
		ammunition := u.Inventory.FindSelectedAmmunition()
		if ammunition == nil || ammunition.Kind != weapon.AmmunitionKind || ammunition.Quantity == 0 {
			return
		}
		ammunition.Quantity--
		damage = ammunition.EnchanceDamageImpact(damage)
	}
	u.State.Reduce(weapon.UseCost)
	instDmg, tmpImp := u.Attack(target, damage)
	if len(instDmg) != 0 || len(tmpImp) != 0 {
		weapon.IncreaseWearout()
	}
	action.InstantDamage = append(action.InstantDamage, instDmg...)
	action.TemporalDamage = append(action.TemporalDamage, tmpImp...)
	action.Accomplished = true
}

func (u *Unit) useMagicOnTarget(action *UseInventoryItemActionResult, target *Unit, magic *Magic) {
	if !u.CheckRequirements(magic.Requirements) || !u.CheckUseCost(magic.UseCost) {
		return
	}
	u.State.Reduce(magic.UseCost)
	if len(magic.Damage) != 0 {
		instDmg, tmpImp := u.Attack(target, magic.Damage)
		action.InstantDamage = append(action.InstantDamage, instDmg...)
		action.TemporalDamage = append(action.TemporalDamage, tmpImp...)
	}
	if len(magic.Modification) != 0 {
		instRec, tmpEnch := u.Modify(target, magic.Modification)
		action.InstantRecovery = append(action.InstantRecovery, instRec...)
		action.TemporalModification = append(action.TemporalModification, tmpEnch...)
	}
	action.Accomplished = true
}

func (u *Unit) useDisposableOnTarget(action *UseInventoryItemActionResult, target *Unit, disposable *Disposable) {
	if disposable.Quantity == 0 {
		return
	}
	disposable.Quantity--
	if len(disposable.Damage) != 0 {
		instDmg, tmpImp := u.Attack(target, disposable.Damage)
		action.InstantDamage = append(action.InstantDamage, instDmg...)
		action.TemporalDamage = append(action.TemporalDamage, tmpImp...)
	}
	if len(disposable.Modification) != 0 {
		instRec, tmpEnch := u.Modify(target, disposable.Modification)
		action.InstantRecovery = append(action.InstantRecovery, instRec...)
		action.TemporalModification = append(action.TemporalModification, tmpEnch...)
	}
	action.Accomplished = true
}
