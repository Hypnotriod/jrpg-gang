package domain

type UseInventoryItemAction struct {
	InstantDamage       []Damage                `json:"instant_damage,omitempty"`
	TemporalImpact      []DamageImpact          `json:"temporal_impact,omitempty"`
	InstantRecovery     []UnitState             `json:"instant_recovery,omitempty"`
	TemporalEnhancement []UnitEnhancementImpact `json:"temporal_enhancement,omitempty"`
}

func (u *Unit) UseInventoryItemOnTarget(target *Unit, uid uint) UseInventoryItemAction {
	action := UseInventoryItemAction{}
	item := u.Inventory.Get(uid)
	if item == nil {
		return action
	}
	switch v := item.(type) {
	case *Weapon:
		u.useWeaponOnTarget(&action, target, v)
	case *Disposable:
		u.useDisposableOnTarget(&action, target, v)
	case *Spell:
		u.useSpellOnTarget(&action, target, v)
	}
	return action
}

func (u *Unit) useWeaponOnTarget(action *UseInventoryItemAction, target *Unit, weapon *Weapon) {
	if !weapon.Equipped {
		return
	}
	instDmg, tmpImp := u.Attack(target, weapon.Impact)
	action.InstantDamage = append(action.InstantDamage, instDmg...)
	action.TemporalImpact = append(action.TemporalImpact, tmpImp...)
}

func (u *Unit) useDisposableOnTarget(action *UseInventoryItemAction, target *Unit, disposable *Disposable) {
	instDmg, tmpImp := u.Attack(target, disposable.Impact)
	action.InstantDamage = append(action.InstantDamage, instDmg...)
	action.TemporalImpact = append(action.TemporalImpact, tmpImp...)
	instRec, tmpEnch := u.Enhance(target, disposable.Enhancement)
	action.InstantRecovery = append(action.InstantRecovery, instRec...)
	action.TemporalEnhancement = append(action.TemporalEnhancement, tmpEnch...)
}

func (u *Unit) useSpellOnTarget(action *UseInventoryItemAction, target *Unit, spell *Spell) {
	instDmg, tmpImp := u.Attack(target, spell.Impact)
	action.InstantDamage = append(action.InstantDamage, instDmg...)
	action.TemporalImpact = append(action.TemporalImpact, tmpImp...)
	instRec, tmpEnch := u.Enhance(target, spell.Enhancement)
	action.InstantRecovery = append(action.InstantRecovery, instRec...)
	action.TemporalEnhancement = append(action.TemporalEnhancement, tmpEnch...)
}
