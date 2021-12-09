package domain

import (
	"fmt"
	"strings"
)

type UseInventoryItemActionResult struct {
	InstantDamage       []Damage                `json:"instantDamage,omitempty"`
	TemporalDamage      []DamageImpact          `json:"temporalDamage,omitempty"`
	InstantRecovery     []UnitState             `json:"instantRecovery,omitempty"`
	TemporalEnhancement []UnitEnhancementImpact `json:"temporalEnhancement,omitempty"`
}

func (r UseInventoryItemActionResult) String() string {
	props := []string{}

	if len(r.InstantDamage) != 0 {
		props = append(props, fmt.Sprintf("instantDamage: %v", r.InstantDamage))
	}
	if len(r.TemporalDamage) != 0 {
		props = append(props, fmt.Sprintf("temporalDamage: %v", r.TemporalDamage))
	}
	if len(r.InstantRecovery) != 0 {
		props = append(props, fmt.Sprintf("instantRecovery: %v", r.InstantRecovery))
	}
	if len(r.TemporalEnhancement) != 0 {
		props = append(props, fmt.Sprintf("temporalEnhancement: %v", r.TemporalEnhancement))
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
	weapon.IncreaseWearOut()
	instDmg, tmpImp := u.Attack(target, weapon.Damage)
	action.InstantDamage = append(action.InstantDamage, instDmg...)
	action.TemporalDamage = append(action.TemporalDamage, tmpImp...)
}

func (u *Unit) useDisposableOnTarget(action *UseInventoryItemActionResult, target *Unit, disposable *Disposable) {
	if len(disposable.Damage) != 0 {
		instDmg, tmpImp := u.Attack(target, disposable.Damage)
		action.InstantDamage = append(action.InstantDamage, instDmg...)
		action.TemporalDamage = append(action.TemporalDamage, tmpImp...)
	}
	if len(disposable.Enhancement) != 0 {
		instRec, tmpEnch := u.Enhance(target, disposable.Enhancement)
		action.InstantRecovery = append(action.InstantRecovery, instRec...)
		action.TemporalEnhancement = append(action.TemporalEnhancement, tmpEnch...)
	}
}

func (u *Unit) useMagicOnTarget(action *UseInventoryItemActionResult, target *Unit, magic *Magic) {
	if len(magic.Damage) != 0 {
		instDmg, tmpImp := u.Attack(target, magic.Damage)
		action.InstantDamage = append(action.InstantDamage, instDmg...)
		action.TemporalDamage = append(action.TemporalDamage, tmpImp...)
	}
	if len(magic.Enhancement) != 0 {
		instRec, tmpEnch := u.Enhance(target, magic.Enhancement)
		action.InstantRecovery = append(action.InstantRecovery, instRec...)
		action.TemporalEnhancement = append(action.TemporalEnhancement, tmpEnch...)
	}
}
