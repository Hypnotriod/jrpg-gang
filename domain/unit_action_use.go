package domain

func (u *Unit) UseInventoryItemOnTarget(target *Unit, uid uint) *ActionResult {
	result := NewActionResult()
	item := u.Inventory.Find(uid)
	if item == nil {
		return result.WithResult(ResultNotFound)
	}
	switch v := item.(type) {
	case *Weapon:
		return u.useWeaponOnTarget(result, target, v)
	case *Disposable:
		return u.useDisposableOnTarget(result, target, v)
	case *Magic:
		return u.useMagicOnTarget(result, target, v)
	}
	return result.WithResult(ResultNotAccomplished)
}

func (u *Unit) useWeaponOnTarget(action *ActionResult, target *Unit, weapon *Weapon) *ActionResult {
	if !weapon.Equipped {
		return action.WithResult(ResultNotEquipped)
	}
	if !u.CheckUseCost(weapon.UseCost) {
		return action.WithResult(ResultCantUse)
	}
	if !weapon.Range.Check(u.Position, target.Position) {
		return action.WithResult(ResultNotReachable)
	}
	var damage []DamageImpact = weapon.Damage
	if weapon.RequiresAmmunition() {
		ammunition := u.Inventory.FindEquippedAmmunition()
		if ammunition == nil {
			return action.WithResult(ResultNoAmmunition)
		}
		if ammunition.Quantity == 0 {
			return action.WithResult(ResultZeroQuantity)
		}
		if ammunition.Kind != weapon.AmmunitionKind {
			return action.WithResult(ResultNotCompatible)
		}
		ammunition.Quantity--
		u.Inventory.FilterAmmunition()
		damage = ammunition.EnchanceDamageImpact(damage)
	}
	u.State.Reduce(weapon.UseCost)
	instDmg, tmpImp := u.Attack(target, damage)
	if len(instDmg) != 0 || len(tmpImp) != 0 {
		weapon.IncreaseWearout()
		u.Inventory.UpdateEquipmentByWeareout()
	}
	action.InstantDamage = append(action.InstantDamage, instDmg...)
	action.TemporalDamage = append(action.TemporalDamage, tmpImp...)
	return action.WithResult(ResultAccomplished)
}

func (u *Unit) useMagicOnTarget(action *ActionResult, target *Unit, magic *Magic) *ActionResult {
	if !u.CheckRequirements(magic.Requirements) || !u.CheckUseCost(magic.UseCost) {
		return action.WithResult(ResultCantUse)
	}
	if !magic.Range.Check(u.Position, target.Position) {
		return action.WithResult(ResultNotReachable)
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
	return action.WithResult(ResultAccomplished)
}

func (u *Unit) useDisposableOnTarget(action *ActionResult, target *Unit, disposable *Disposable) *ActionResult {
	if disposable.Quantity == 0 {
		return action.WithResult(ResultZeroQuantity)
	}
	if disposable.IsHarmful() {
		return action.WithResult(ResultNotAllowed)
	}
	disposable.Quantity--
	u.Inventory.FilterDisposable()
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
	return action.WithResult(ResultAccomplished)
}
