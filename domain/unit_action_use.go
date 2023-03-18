package domain

func (u *Unit) UseInventoryItemOnTarget(target *Unit, uid uint, result *ActionResult) *ActionResult {
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
	if action.Result == ResultEmpty {
		if !weapon.Equipped {
			return action.WithResult(ResultNotEquipped)
		}
		if !u.CheckUseCost(weapon.UseCost) {
			return action.WithResult(ResultCantUse)
		}
		if !u.CanReach(target, weapon.Range) {
			return action.WithResult(ResultNotReachable)
		}
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
		damage = ammunition.EnchanceDamageImpact(damage)
	}
	instDmg, tmpImp := u.Attack(target, damage)
	if action.Result == ResultEmpty {
		u.State.Reduce(weapon.UseCost)
		if len(instDmg) != 0 || len(tmpImp) != 0 {
			weapon.IncreaseWearout()
			u.Inventory.UpdateEquipmentByWeareout()
		}
	}
	action.InstantDamage[target.Uid] = append(action.InstantDamage[target.Uid], instDmg...)
	action.TemporalDamage[target.Uid] = append(action.TemporalDamage[target.Uid], tmpImp...)
	return action.WithResult(ResultAccomplished)
}

func (u *Unit) useMagicOnTarget(action *ActionResult, target *Unit, magic *Magic) *ActionResult {
	if action.Result == ResultEmpty {
		if !u.CheckRequirements(magic.Requirements) || !u.CheckUseCost(magic.UseCost) {
			return action.WithResult(ResultCantUse)
		}
		if !u.CanReach(target, magic.Range) {
			return action.WithResult(ResultNotReachable)
		}
		u.State.Reduce(magic.UseCost)
	}
	if len(magic.Damage) != 0 {
		instDmg, tmpImp := u.Attack(target, magic.Damage)
		action.InstantDamage[target.Uid] = append(action.InstantDamage[target.Uid], instDmg...)
		action.TemporalDamage[target.Uid] = append(action.TemporalDamage[target.Uid], tmpImp...)
	}
	if len(magic.Modification) != 0 {
		instRec, tmpEnch := u.Modify(target, magic.Modification)
		action.InstantRecovery[target.Uid] = append(action.InstantRecovery[target.Uid], instRec...)
		action.TemporalModification[target.Uid] = append(action.TemporalModification[target.Uid], tmpEnch...)
	}
	return action.WithResult(ResultAccomplished)
}

func (u *Unit) useDisposableOnTarget(action *ActionResult, target *Unit, disposable *Disposable) *ActionResult {
	if action.Result == ResultEmpty {
		if disposable.Quantity == 0 {
			return action.WithResult(ResultZeroQuantity)
		}
		if !u.CanReach(target, disposable.Range) {
			return action.WithResult(ResultNotReachable)
		}
		disposable.Quantity--
	}
	if len(disposable.Damage) != 0 {
		instDmg, tmpImp := u.Attack(target, disposable.Damage)
		action.InstantDamage[target.Uid] = append(action.InstantDamage[target.Uid], instDmg...)
		action.TemporalDamage[target.Uid] = append(action.TemporalDamage[target.Uid], tmpImp...)
	}
	if len(disposable.Modification) != 0 {
		instRec, tmpEnch := u.Modify(target, disposable.Modification)
		action.InstantRecovery[target.Uid] = append(action.InstantRecovery[target.Uid], instRec...)
		action.TemporalModification[target.Uid] = append(action.TemporalModification[target.Uid], tmpEnch...)
	}
	return action.WithResult(ResultAccomplished)
}
