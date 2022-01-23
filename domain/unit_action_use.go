package domain

func (u *Unit) UseInventoryItemOnTarget(target *Unit, uid uint) ActionResult {
	action := ActionResult{}
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

func (u *Unit) useWeaponOnTarget(action *ActionResult, target *Unit, weapon *Weapon) *ActionResult {
	if !weapon.Equipped {
		return action.WithResultType(NotEquipped)
	}
	if !u.CheckUseCost(weapon.UseCost) {
		return action.WithResultType(CantUse)
	}
	var damage []DamageImpact = weapon.Damage
	if weapon.RequiresAmmunition() {
		ammunition := u.Inventory.FindEquippedAmmunition()
		if ammunition == nil {
			return action.WithResultType(NoAmmunition)
		}
		if ammunition.Quantity == 0 {
			return action.WithResultType(ZeroQuantity)
		}
		if ammunition.Kind != weapon.AmmunitionKind {
			return action.WithResultType(NotCompatible)
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

func (u *Unit) useMagicOnTarget(action *ActionResult, target *Unit, magic *Magic) *ActionResult {
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

func (u *Unit) useDisposableOnTarget(action *ActionResult, target *Unit, disposable *Disposable) *ActionResult {
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
