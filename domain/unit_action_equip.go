package domain

func (u *Unit) Equip(uid uint) bool {
	var equipment *Equipment
	item := u.Inventory.Get(uid)
	if item == nil {
		return false
	}
	switch v := item.(type) {
	case *Weapon:
		equipment = &v.Equipment
	case *Armor:
		equipment = &v.Equipment
	default:
		return false
	}
	if equipment.SlotsNumber > u.Slots[equipment.Slot] || equipment.IsBroken() {
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
