package domain

import (
	"fmt"
	"strings"
)

type UseInventoryItemActionResult struct {
	InstantDamage        []Damage                 `json:"instantDamage,omitempty"`
	TemporalDamage       []DamageImpact           `json:"temporalDamage,omitempty"`
	InstantRecovery      []UnitState              `json:"instantRecovery,omitempty"`
	TemporalModification []UnitModificationImpact `json:"temporalModification,omitempty"`
}

func (r UseInventoryItemActionResult) String() string {
	props := []string{}

	if len(r.InstantDamage) != 0 {
		props = append(props, fmt.Sprintf("instant damage: %v", r.InstantDamage))
	}
	if len(r.TemporalDamage) != 0 {
		props = append(props, fmt.Sprintf("temporal damage: %v", r.TemporalDamage))
	}
	if len(r.InstantRecovery) != 0 {
		props = append(props, fmt.Sprintf("instant recovery: %v", r.InstantRecovery))
	}
	if len(r.TemporalModification) != 0 {
		props = append(props, fmt.Sprintf("temporal modification: %v", r.TemporalModification))
	}

	return strings.Join(props, ", ")
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
	if !weapon.Equipped {
		return
	}
	weapon.IncreaseWearout()
	instDmg, tmpImp := u.Attack(target, weapon.Damage)
	action.InstantDamage = append(action.InstantDamage, instDmg...)
	action.TemporalDamage = append(action.TemporalDamage, tmpImp...)
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
}

func (u *Unit) useMagicOnTarget(action *UseInventoryItemActionResult, target *Unit, magic *Magic) {
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
}
