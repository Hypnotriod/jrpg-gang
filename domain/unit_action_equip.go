package domain

func (u *Unit) Equip(uid uint) ActionResult {
	action := ActionResult{}
	if u.Inventory.FindAmmunition(uid) != nil {
		u.Inventory.EqipAmmunition(uid)
		return *action.WithResultType(ResultAccomplished)
	}
	equipment := u.Inventory.FindEquipment(uid)
	if equipment == nil {
		return *action.WithResultType(ResultNotFound)
	}
	if equipment.SlotsNumber > u.Slots[equipment.Slot] {
		return *action.WithResultType(ResultNotEnoughSlots)
	}
	if equipment.IsBroken() {
		return *action.WithResultType(ResultIsBroken)
	}
	if !u.CheckRequirements(equipment.Requirements) {
		return *action.WithResultType(ResultCantUse)
	}
	freeSlots := u.GetFreeSlotsNumber(equipment.Slot)
	if freeSlots < equipment.SlotsNumber {
		u.UnequipBySlot(equipment.Slot, equipment.SlotsNumber-freeSlots)
	}
	equipment.Equipped = true
	return *action.WithResultType(ResultAccomplished)
}

func (u *Unit) Unequip(uid uint) ActionResult {
	action := ActionResult{}
	if u.Inventory.FindAmmunition(uid) != nil {
		u.Inventory.UnequipAmmunition()
		return *action.WithResultType(ResultAccomplished)
	}
	equipment := u.Inventory.GetEquipment(true)
	for i := range equipment {
		if equipment[i].Uid == uid {
			equipment[i].Equipped = false
			return *action.WithResultType(ResultAccomplished)
		}
	}
	return *action.WithResultType(ResultNotFound)
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
