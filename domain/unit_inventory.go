package domain

import (
	"fmt"
	"jrpg-gang/util"
)

type UnitInventory struct {
	Weapon     []Weapon     `json:"weapon,omitempty"`
	Magic      []Magic      `json:"magic,omitempty"`
	Armor      []Armor      `json:"armor,omitempty"`
	Disposable []Disposable `json:"disposable,omitempty"`
	Ammunition []Ammunition `json:"ammunition,omitempty"`
}

func (i UnitInventory) String() string {
	return fmt.Sprintf(
		"weapon: [%s], magic: [%s], armor: [%s], disposable: [%s], ammunition: [%s]",
		util.AsCommaSeparatedObjectsSlice(i.Weapon),
		util.AsCommaSeparatedObjectsSlice(i.Magic),
		util.AsCommaSeparatedObjectsSlice(i.Armor),
		util.AsCommaSeparatedObjectsSlice(i.Disposable),
		util.AsCommaSeparatedObjectsSlice(i.Ammunition))
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
		if s := i.FindDisposableByCode(v.Code); s != nil {
			s.Quantity += v.Quantity
		} else {
			i.Disposable = append(i.Disposable, v)
		}
		return true
	case *Disposable:
		if s := i.FindDisposableByCode(v.Code); s != nil {
			s.Quantity += v.Quantity
		} else {
			i.Disposable = append(i.Disposable, *v)
		}
		return true
	case Ammunition:
		if s := i.FindAmmunitionByCode(v.Code); s != nil {
			s.Quantity += v.Quantity
		} else {
			i.Ammunition = append(i.Ammunition, v)
		}
		return true
	case *Ammunition:
		if s := i.FindAmmunitionByCode(v.Code); s != nil {
			s.Quantity += v.Quantity
		} else {
			i.Ammunition = append(i.Ammunition, *v)
		}
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
	for n := range i.Ammunition {
		if i.Ammunition[n].Uid == uid {
			return &i.Ammunition[n]
		}
	}
	return nil
}

func (i *UnitInventory) FindItem(uid uint) *Item {
	for n := range i.Weapon {
		if i.Weapon[n].Uid == uid {
			return &i.Weapon[n].Item
		}
	}
	for n := range i.Magic {
		if i.Magic[n].Uid == uid {
			return &i.Magic[n].Item
		}
	}
	for n := range i.Armor {
		if i.Armor[n].Uid == uid {
			return &i.Armor[n].Item
		}
	}
	for n := range i.Disposable {
		if i.Disposable[n].Uid == uid {
			return &i.Disposable[n].Item
		}
	}
	for n := range i.Ammunition {
		if i.Ammunition[n].Uid == uid {
			return &i.Ammunition[n].Item
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

func (i *UnitInventory) FindEquipmentByCode(code ItemCode) *Equipment {
	if code == ItemCodeEmpty {
		return nil
	}
	equipment := i.GetEquipment(false)
	for n := range equipment {
		if equipment[n].Code == code {
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

func (i *UnitInventory) FindWeaponByCode(code ItemCode) *Weapon {
	if code == ItemCodeEmpty {
		return nil
	}
	for n := range i.Weapon {
		if i.Weapon[n].Code == code {
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

func (i *UnitInventory) FindMagicByCode(code ItemCode) *Magic {
	if code == ItemCodeEmpty {
		return nil
	}
	for n := range i.Magic {
		if i.Magic[n].Code == code {
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

func (i *UnitInventory) FindArmorByCode(code ItemCode) *Armor {
	if code == ItemCodeEmpty {
		return nil
	}
	for n := range i.Armor {
		if i.Armor[n].Code == code {
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

func (i *UnitInventory) FindDisposableByCode(code ItemCode) *Disposable {
	if code == ItemCodeEmpty {
		return nil
	}
	for n := range i.Disposable {
		if i.Disposable[n].Code == code {
			return &i.Disposable[n]
		}
	}
	return nil
}

func (i *UnitInventory) FilterDisposable() {
	var filtered []Disposable
	for _, disp := range i.Disposable {
		if disp.Quantity != 0 {
			filtered = append(filtered, disp)
		}
	}
	i.Disposable = filtered
}

func (i *UnitInventory) FindAmmunition(uid uint) *Ammunition {
	for n := range i.Ammunition {
		if i.Ammunition[n].Uid == uid {
			return &i.Ammunition[n]
		}
	}
	return nil
}

func (i *UnitInventory) FindAmmunitionByCode(code ItemCode) *Ammunition {
	if code == ItemCodeEmpty {
		return nil
	}
	for n := range i.Ammunition {
		if i.Ammunition[n].Code == code {
			return &i.Ammunition[n]
		}
	}
	return nil
}

func (i *UnitInventory) FindEquippedAmmunition() *Ammunition {
	for n := range i.Ammunition {
		if i.Ammunition[n].Equipped {
			return &i.Ammunition[n]
		}
	}
	return nil
}

func (i *UnitInventory) EqipAmmunition(uid uint) {
	for n := range i.Ammunition {
		i.Ammunition[n].Equipped = i.Ammunition[n].Uid == uid
	}
}

func (i *UnitInventory) UnequipAmmunition() {
	for n := range i.Ammunition {
		i.Ammunition[n].Equipped = false
	}
}

func (i *UnitInventory) FilterAmmunition() {
	var filtered []Ammunition
	for _, amm := range i.Ammunition {
		if amm.Quantity != 0 {
			filtered = append(filtered, amm)
		}
	}
	i.Ammunition = filtered
}

func (i *UnitInventory) GetItemType(uid uint) ItemType {
	item := i.Find(uid)
	if item == nil {
		return ItemTypeNone
	}
	switch item.(type) {
	case *Armor:
		return ItemTypeMagic
	case *Weapon:
		return ItemTypeWeapon
	case *Magic:
		return ItemTypeMagic
	case *Disposable:
		return ItemTypeDisposable
	case *Ammunition:
		return ItemTypeAmmunition
	}
	return ItemTypeNone
}

func (i *UnitInventory) Prepare() {
	if i.Ammunition == nil {
		i.Ammunition = []Ammunition{}
	}
	if i.Armor == nil {
		i.Armor = []Armor{}
	}
	if i.Disposable == nil {
		i.Disposable = []Disposable{}
	}
	if i.Magic == nil {
		i.Magic = []Magic{}
	}
	if i.Weapon == nil {
		i.Weapon = []Weapon{}
	}
}

func (i *UnitInventory) PopulateUids(rndGen *util.RndGen) {
	for j := range i.Ammunition {
		i.Ammunition[j].Uid = rndGen.NextUid()
	}
	for j := range i.Armor {
		i.Armor[j].Uid = rndGen.NextUid()
	}
	for j := range i.Disposable {
		i.Disposable[j].Uid = rndGen.NextUid()
	}
	for j := range i.Magic {
		i.Magic[j].Uid = rndGen.NextUid()
	}
	for j := range i.Weapon {
		i.Weapon[j].Uid = rndGen.NextUid()
	}
}
