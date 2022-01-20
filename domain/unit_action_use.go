package domain

func (u *Unit) UseInventoryItemOnTarget(target *Unit, uid uint) UseInventoryItemActionResult {
	action := UseInventoryItemActionResult{}
	action.ResultType = Accomplished
	item := u.Inventory.Find(uid)
	if item == nil {
		action.ResultType = NotFound
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
		action.ResultType = IsNotEquipped
		return
	}
	if !u.CheckUseCost(weapon.UseCost) {
		action.ResultType = CantUse
		return
	}
	var damage []DamageImpact = weapon.Damage
	if weapon.RequiresAmmunition() {
		ammunition := u.Inventory.FindSelectedAmmunition()
		if ammunition == nil {
			action.ResultType = HasNoAmmunition
			return
		}
		if ammunition.Quantity == 0 {
			action.ResultType = ZeroQuantity
			return
		}
		if ammunition.Kind != weapon.AmmunitionKind {
			action.ResultType = IsNotCompatible
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
}

func (u *Unit) useMagicOnTarget(action *UseInventoryItemActionResult, target *Unit, magic *Magic) {
	if !u.CheckRequirements(magic.Requirements) || !u.CheckUseCost(magic.UseCost) {
		action.ResultType = CantUse
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
}

func (u *Unit) useDisposableOnTarget(action *UseInventoryItemActionResult, target *Unit, disposable *Disposable) {
	if disposable.Quantity == 0 {
		action.ResultType = ZeroQuantity
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
