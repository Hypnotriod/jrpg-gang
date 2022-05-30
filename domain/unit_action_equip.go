package domain

func (u *Unit) Equip(uid uint) *ActionResult {
	result := NewActionResult()
	if u.Inventory.FindAmmunition(uid) != nil {
		u.Inventory.EqipAmmunition(uid)
		return result.WithResult(ResultAccomplished)
	}
	equipment := u.Inventory.FindEquipment(uid)
	if equipment == nil {
		return result.WithResult(ResultNotFound)
	}
	if equipment.SlotsNumber > u.Slots[equipment.Slot] {
		return result.WithResult(ResultNotEnoughSlots)
	}
	if equipment.IsBroken() {
		return result.WithResult(ResultIsBroken)
	}
	if !u.CheckRequirements(equipment.Requirements) {
		return result.WithResult(ResultCantUse)
	}
	freeSlots := u.GetFreeSlotsNumber(equipment.Slot)
	if freeSlots < equipment.SlotsNumber {
		u.UnequipBySlot(equipment.Slot, equipment.SlotsNumber-freeSlots)
	}
	equipment.Equipped = true
	return result.WithResult(ResultAccomplished)
}

func (u *Unit) Unequip(uid uint) *ActionResult {
	result := NewActionResult()
	if u.Inventory.FindAmmunition(uid) != nil {
		u.Inventory.UnequipAmmunition()
		return result.WithResult(ResultAccomplished)
	}
	equipment := u.Inventory.GetEquipment(true)
	for i := range equipment {
		if equipment[i].Uid == uid {
			equipment[i].Equipped = false
			return result.WithResult(ResultAccomplished)
		}
	}
	return result.WithResult(ResultNotFound)
}

func (u *Unit) UnequipBySlot(slot EquipmentSlot, slotsToRemove uint) {
	equipment := u.Inventory.GetEquipmentBySlot(slot, true)
	for i := range equipment {
		if slotsToRemove == 0 {
			break
		}
		equipment[i].Equipped = false
		slotsToRemove--
	}
}

func (u *Unit) GetFreeSlotsNumber(slot EquipmentSlot) uint {
	return u.Slots[slot] - u.Inventory.GetEquippedSlotsNumber(slot)
}
