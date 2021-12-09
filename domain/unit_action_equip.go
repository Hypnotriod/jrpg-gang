package domain

func (u *Unit) Equip(uid uint) bool {
	equipment := u.Inventory.FindEquipment(uid)
	if equipment == nil || equipment.SlotsNumber > u.Slots[equipment.Slot] || equipment.IsBroken() {
		return false
	}
	u.UnequipBySlot(equipment.Slot, equipment.SlotsNumber)
	equipment.Equipped = true
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
