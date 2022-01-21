package domain

func (u *Unit) UseInventoryItemOnTarget(target *Unit, uid uint) UseInventoryItemActionResult {
	action := UseInventoryItemActionResult{}
	item := u.Inventory.Find(uid)
	if item == nil {
		return *action.WithResultType(NotFound)
	}
	switch v := item.(type) {
	case *Weapon:
		return *u.useWeaponOnTarget(&action, target, v)
	case *Disposable:
		return *u.useDisposableOnTarget(&action, target, v)
	case *Magic:
		return *u.useMagicOnTarget(&action, target, v)
	}
	return *action.WithResultType(NotAccomplished)
}

func (u *Unit) useWeaponOnTarget(action *UseInventoryItemActionResult, target *Unit, weapon *Weapon) *UseInventoryItemActionResult {
	if !weapon.Equipped {
		return action.WithResultType(IsNotEquipped)
	}
	if !u.CheckUseCost(weapon.UseCost) {
		return action.WithResultType(CantUse)
	}
	var damage []DamageImpact = weapon.Damage
	if weapon.RequiresAmmunition() {
		ammunition := u.Inventory.FindSelectedAmmunition()
		if ammunition == nil {
			return action.WithResultType(HasNoAmmunition)
		}
		if ammunition.Quantity == 0 {
			return action.WithResultType(ZeroQuantity)
		}
		if ammunition.Kind != weapon.AmmunitionKind {
			return action.WithResultType(IsNotCompatible)
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
	return action.WithResultType(Accomplished)
}

func (u *Unit) useMagicOnTarget(action *UseInventoryItemActionResult, target *Unit, magic *Magic) *UseInventoryItemActionResult {
	if !u.CheckRequirements(magic.Requirements) || !u.CheckUseCost(magic.UseCost) {
		return action.WithResultType(CantUse)
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
	return action.WithResultType(Accomplished)
}

func (u *Unit) useDisposableOnTarget(action *UseInventoryItemActionResult, target *Unit, disposable *Disposable) *UseInventoryItemActionResult {
	if disposable.Quantity == 0 {
		return action.WithResultType(ZeroQuantity)
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
	return action.WithResultType(Accomplished)
}
