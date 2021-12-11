package domain

func (u *Unit) Equip(uid uint) bool {
	item := u.Inventory.FindEquipment(uid)
	if item == nil ||
		item.SlotsNumber > u.Slots[item.Slot] ||
		item.IsBroken() ||
		u.CheckRequirements(item.Requirements) {
		return false
	}
	freeSlots := u.GetFreeSlotsNumber(item.Slot)
	if freeSlots < item.SlotsNumber {
		u.UnequipBySlot(item.Slot, item.SlotsNumber-freeSlots)
	}
	item.Equipped = true
	return true
}

func (u *Unit) Unequip(uid uint) bool {
	equipment := u.Inventory.GetEquipment(true)
	for i := range equipment {
		if equipment[i].Uid == uid {
			equipment[i].Equipped = false
			return true
		}
	}
	return false
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
