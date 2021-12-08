package domain

import "fmt"

type UnitInventory struct {
	Weapon     []Weapon     `json:"weapon,omitempty"`
	Spell      []Spell      `json:"spell,omitempty"`
	Armor      []Armor      `json:"armor,omitempty"`
	Disposable []Disposable `json:"disposable,omitempty"`
}

func (i UnitInventory) String() string {
	return fmt.Sprintf(
		"weapon: %v, armor: %v, disposable: %v",
		i.Weapon, i.Armor, i.Disposable)
}

func (i *UnitInventory) IncreaseArmorWearOut(equipped bool) {
	for n := range i.Armor {
		item := &i.Armor[n]
		if item.Equipped || !equipped {
			item.IncreaseWearOut()
		}
	}
}

func (i *UnitInventory) CheckEquipmentWeareout() {
	for n := range i.Armor {
		item := &i.Armor[n].Equipment
		if item.Equipped || item.Wearout >= item.Durability {
			item.Equipped = false
		}
	}
	for n := range i.Weapon {
		item := &i.Weapon[n].Equipment
		if item.Equipped || item.Wearout >= item.Durability {
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

func (i *UnitInventory) Add(item interface{}) bool {
	switch v := item.(type) {
	case Weapon:
		i.Weapon = append(i.Weapon, v)
		return true
	case *Weapon:
		i.Weapon = append(i.Weapon, *v)
		return true
	case Spell:
		i.Spell = append(i.Spell, v)
		return true
	case *Spell:
		i.Spell = append(i.Spell, *v)
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

func (i *UnitInventory) Get(uid uint) interface{} {
	for n := range i.Weapon {
		if i.Weapon[n].Uid == uid {
			return &i.Weapon[n]
		}
	}
	for n := range i.Spell {
		if i.Spell[n].Uid == uid {
			return &i.Spell[n]
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

func (i *UnitInventory) GetWeapon(uid uint) *Weapon {
	for n := range i.Weapon {
		if i.Weapon[n].Uid == uid {
			return &i.Weapon[n]
		}
	}
	return nil
}

func (i *UnitInventory) GetSpell(uid uint) *Spell {
	for n := range i.Spell {
		if i.Spell[n].Uid == uid {
			return &i.Spell[n]
		}
	}
	return nil
}

func (i *UnitInventory) GetArmor(uid uint) *Armor {
	for n := range i.Armor {
		if i.Armor[n].Uid == uid {
			return &i.Armor[n]
		}
	}
	return nil
}

func (i *UnitInventory) GetDisposable(uid uint) *Disposable {
	for n := range i.Disposable {
		if i.Disposable[n].Uid == uid {
			return &i.Disposable[n]
		}
	}
	return nil
}
