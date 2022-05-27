package shop

import (
	"jrpg-gang/domain"
	"jrpg-gang/util"
)

type GameShop struct {
	Items  *domain.UnitInventory `json:"items"`
	rndGen *util.RndGen
}

func NewGameShop(items *domain.UnitInventory) *GameShop {
	s := &GameShop{}
	s.rndGen = util.NewRndGen()
	s.Items = items
	s.Items.Prepare()
	s.Items.PopulateUids(s.rndGen)
	return s
}

func (s *GameShop) HasItem(uid uint) bool {
	item := s.Items.FindItem(uid)
	return item != nil
}

func (s *GameShop) CheckPrice(unit *domain.Unit, uid uint) bool {
	item := s.Items.FindItem(uid)
	return item != nil && item.Price.Check(unit.Booty)
}

func (s *GameShop) Buy(unit *domain.Unit, uid uint, rndGen *util.RndGen) bool {
	item := s.Items.FindItem(uid)
	if item == nil {
		return false
	}
	if !item.Price.Check(unit.Booty) {
		return false
	}
	unit.Booty.Reduce(item.Price)
	if itemRef := s.Items.FindWeapon(uid); itemRef != nil {
		itemClone := *itemRef
		itemClone.Uid = rndGen.NextUid()
		unit.Inventory.Add(itemClone)
		return true
	}
	if itemRef := s.Items.FindMagic(uid); itemRef != nil {
		itemClone := *itemRef
		itemClone.Uid = rndGen.NextUid()
		unit.Inventory.Add(itemClone)
		return true
	}
	if itemRef := s.Items.FindArmor(uid); itemRef != nil {
		itemClone := *itemRef
		itemClone.Uid = rndGen.NextUid()
		unit.Inventory.Add(itemClone)
		return true
	}
	if itemRef := s.Items.FindDisposable(uid); itemRef != nil {
		itemClone := *itemRef
		itemClone.Uid = rndGen.NextUid()
		unit.Inventory.Add(itemClone)
		return true
	}
	if itemRef := s.Items.FindAmmunition(uid); itemRef != nil {
		itemClone := *itemRef
		itemClone.Uid = rndGen.NextUid()
		unit.Inventory.Add(itemClone)
		return true
	}
	return false
}
