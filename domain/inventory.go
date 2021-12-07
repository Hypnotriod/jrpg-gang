package domain

import "fmt"

type Inventory struct {
	Weapon     []Weapon     `json:"weapon,omitempty"`
	Armor      []Armor      `json:"armor,omitempty"`
	Disposable []Disposable `json:"disposable,omitempty"`
}

func (i Inventory) String() string {
	return fmt.Sprintf(
		"weapon: %v, armor: %v, disposable: %v",
		i.Weapon, i.Armor, i.Disposable)
}

func (i *Inventory) GetEquipment(equipped bool) []*Equipment {
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

func (i *Inventory) GetEquipmentBySlot(slot EquipmentSlot, equipped bool) []*Equipment {
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

func (i *Inventory) Add(item interface{}) bool {
	switch v := item.(type) {
	case Weapon:
		i.Weapon = append(i.Weapon, v)
		return true
	case *Weapon:
		i.Weapon = append(i.Weapon, *v)
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

func (i *Inventory) Get(id string) interface{} {
	for n := range i.Weapon {
		if i.Weapon[n].Id == id {
			return &i.Weapon[n]
		}
	}
	for n := range i.Armor {
		if i.Armor[n].Id == id {
			return &i.Armor[n]
		}
	}
	for n := range i.Disposable {
		if i.Disposable[n].Id == id {
			return &i.Disposable[n]
		}
	}
	return nil
}

func (i *Inventory) GetWeapon(id string) *Weapon {
	for n := range i.Weapon {
		if i.Weapon[n].Id == id {
			return &i.Weapon[n]
		}
	}
	return nil
}

func (i *Inventory) GetArmor(id string) *Armor {
	for n := range i.Armor {
		if i.Armor[n].Id == id {
			return &i.Armor[n]
		}
	}
	return nil
}

func (i *Inventory) GetDisposable(id string) *Disposable {
	for n := range i.Disposable {
		if i.Disposable[n].Id == id {
			return &i.Disposable[n]
		}
	}
	return nil
}
