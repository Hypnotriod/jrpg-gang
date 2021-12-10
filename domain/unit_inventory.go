package domain

import "fmt"

type UnitInventory struct {
	Weapon     []Weapon     `json:"weapon,omitempty"`
	Magic      []Magic      `json:"magic,omitempty"`
	Armor      []Armor      `json:"armor,omitempty"`
	Disposable []Disposable `json:"disposable,omitempty"`
}

func (i UnitInventory) String() string {
	return fmt.Sprintf(
		"weapon: %v, armor: %v, disposable: %v",
		i.Weapon, i.Armor, i.Disposable)
}

func (i *UnitInventory) IncreaseArmorWearout(equipped bool) {
	for n := range i.Armor {
		item := &i.Armor[n]
		if item.Equipped || !equipped {
			item.IncreaseWearout()
		}
	}
}

func (i *UnitInventory) CheckEquipmentWeareout() {
	for n := range i.Armor {
		item := &i.Armor[n].Equipment
		if item.Equipped || item.IsBroken() {
			item.Equipped = false
		}
	}
	for n := range i.Weapon {
		item := &i.Weapon[n].Equipment
		if item.Equipped || item.IsBroken() {
			item.Equipped = false
		}
	}
}

func (i *UnitInventory) GetEquipment(equipped bool) []*Equipment {
	equipment := []*Equipment{}
	for n := range i.Armor {
		item := &i.Armor[n].Equipment
		if item.Equipped || !equipped {
			equipment = append(equipment, item)
		}
	}
	for n := range i.Weapon {
		item := &i.Weapon[n].Equipment
		if item.Equipped || !equipped {
			equipment = append(equipment, item)
		}
	}
	return equipment
}

func (i *UnitInventory) GetEquipmentBySlot(slot EquipmentSlot, equipped bool) []*Equipment {
	equipment := []*Equipment{}
	for n := range i.Armor {
		item := &i.Armor[n].Equipment
		if (item.Equipped || !equipped) && item.Slot == slot {
			equipment = append(equipment, item)
		}
	}
	for n := range i.Weapon {
		item := &i.Weapon[n].Equipment
		if (item.Equipped || !equipped) && item.Slot == slot {
			equipment = append(equipment, item)
		}
	}
	return equipment
}

func (i *UnitInventory) GetEquippedSlotsNumber(slot EquipmentSlot) uint {
	var slotsNumber uint = 0
	equipment := i.GetEquipmentBySlot(slot, true)
	for n := range equipment {
		slotsNumber += equipment[n].SlotsNumber
	}
	return slotsNumber
}

func (i *UnitInventory) Add(item interface{}) bool {
	switch v := item.(type) {
	case Weapon:
		i.Weapon = append(i.Weapon, v)
		return true
	case *Weapon:
		i.Weapon = append(i.Weapon, *v)
		return true
	case Magic:
		i.Magic = append(i.Magic, v)
		return true
	case *Magic:
		i.Magic = append(i.Magic, *v)
		return true
	case Armor:
		i.Armor = append(i.Armor, v)
		return true
	case *Armor:
		i.Armor = append(i.Armor, *v)
		return true
	case Disposable:
		i.Disposable = append(i.Disposable, v)
		return true
	case *Disposable:
		i.Disposable = append(i.Disposable, *v)
		return true
	}
	return false
}

func (i *UnitInventory) Find(uid uint) interface{} {
	for n := range i.Weapon {
		if i.Weapon[n].Uid == uid {
			return &i.Weapon[n]
		}
	}
	for n := range i.Magic {
		if i.Magic[n].Uid == uid {
			return &i.Magic[n]
		}
	}
	for n := range i.Armor {
		if i.Armor[n].Uid == uid {
			return &i.Armor[n]
		}
	}
	for n := range i.Disposable {
		if i.Disposable[n].Uid == uid {
			return &i.Disposable[n]
		}
	}
	return nil
}

func (i *UnitInventory) FindEquipment(uid uint) *Equipment {
	equipment := i.GetEquipment(false)
	for n := range equipment {
		if equipment[n].Uid == uid {
			return equipment[n]
		}
	}
	return nil
}

func (i *UnitInventory) FindWeapon(uid uint) *Weapon {
	for n := range i.Weapon {
		if i.Weapon[n].Uid == uid {
			return &i.Weapon[n]
		}
	}
	return nil
}

func (i *UnitInventory) FindMagic(uid uint) *Magic {
	for n := range i.Magic {
		if i.Magic[n].Uid == uid {
			return &i.Magic[n]
		}
	}
	return nil
}

func (i *UnitInventory) FindArmor(uid uint) *Armor {
	for n := range i.Armor {
		if i.Armor[n].Uid == uid {
			return &i.Armor[n]
		}
	}
	return nil
}

func (i *UnitInventory) FindDisposable(uid uint) *Disposable {
	for n := range i.Disposable {
		if i.Disposable[n].Uid == uid {
			return &i.Disposable[n]
		}
	}
	return nil
}

func (i *UnitInventory) FilterDisposable() {
	var filteredDisposable []Disposable
	for _, disp := range i.Disposable {
		if disp.Quantity != 0 {
			filteredDisposable = append(filteredDisposable, disp)
		}
	}
	i.Disposable = filteredDisposable
}
