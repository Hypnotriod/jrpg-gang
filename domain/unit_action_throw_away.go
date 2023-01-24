package domain

func (u *Unit) ThrowAway(uid uint, quantity uint) *ActionResult {
	result := NewActionResult()
	item := u.Inventory.FindItem(uid)
	if item == nil {
		return result.WithResult(ResultNotFound)
	}
	if !item.CanBeThrownAway {
		return result.WithResult(ResultNotAllowed)
	}
	if item.Type == ItemTypeAmmunition && u.Inventory.RemoveAmmunition(uid, quantity) == nil {
		return result.WithResult(ResultNotEnoughResources)
	} else if item.Type == ItemTypeDisposable && u.Inventory.RemoveDisposable(uid, quantity) == nil {
		return result.WithResult(ResultNotEnoughResources)
	} else if quantity != 1 {
		return result.WithResult(ResultNotAllowed)
	} else if item.Type == ItemTypeArmor && u.Inventory.RemoveArmor(uid) == nil {
		return result.WithResult(ResultNotEnoughResources)
	} else if item.Type == ItemTypeWeapon && u.Inventory.RemoveWeapon(uid) == nil {
		return result.WithResult(ResultNotEnoughResources)
	} else if item.Type == ItemTypeMagic && u.Inventory.RemoveMagic(uid) == nil {
		return result.WithResult(ResultNotEnoughResources)
	}
	return result.WithResult(ResultAccomplished)
}
