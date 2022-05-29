package engine

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
	s.Items.UnequipAmmunition()
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

func (s *GameShop) ExecuteAction(action domain.Action, unit *domain.Unit, rndGen *util.RndGen) *domain.ActionResult {
	switch action.Action {
	case domain.ActionBuy:
		return s.buy(action, unit, rndGen)
	}
	return domain.NewActionResult().WithResultType(domain.ResultNotAccomplished)
}

func (s *GameShop) buy(action domain.Action, unit *domain.Unit, rndGen *util.RndGen) *domain.ActionResult {
	item := s.Items.FindItem(action.ItemUid)
	if item == nil {
		return domain.NewActionResult().WithResultType(domain.ResultNotFound)
	}
	if !item.Price.Check(unit.Booty) {
		return domain.NewActionResult().WithResultType(domain.ResultNotEnoughResources)
	}
	unit.Booty.Reduce(item.Price)
	if itemRef := s.Items.FindWeapon(action.ItemUid); itemRef != nil {
		itemClone := *itemRef
		itemClone.Uid = rndGen.NextUid()
		unit.Inventory.Add(itemClone)
		return domain.NewActionResult().WithResultType(domain.ResultAccomplished)
	}
	if itemRef := s.Items.FindMagic(action.ItemUid); itemRef != nil {
		itemClone := *itemRef
		itemClone.Uid = rndGen.NextUid()
		unit.Inventory.Add(itemClone)
		return domain.NewActionResult().WithResultType(domain.ResultAccomplished)
	}
	if itemRef := s.Items.FindArmor(action.ItemUid); itemRef != nil {
		itemClone := *itemRef
		itemClone.Uid = rndGen.NextUid()
		unit.Inventory.Add(itemClone)
		return domain.NewActionResult().WithResultType(domain.ResultAccomplished)
	}
	if itemRef := s.Items.FindDisposable(action.ItemUid); itemRef != nil {
		itemClone := *itemRef
		itemClone.Uid = rndGen.NextUid()
		unit.Inventory.Add(itemClone)
		return domain.NewActionResult().WithResultType(domain.ResultAccomplished)
	}
	if itemRef := s.Items.FindAmmunition(action.ItemUid); itemRef != nil {
		itemClone := *itemRef
		itemClone.Uid = rndGen.NextUid()
		unit.Inventory.Add(itemClone)
		return domain.NewActionResult().WithResultType(domain.ResultAccomplished)
	}
	return domain.NewActionResult().WithResultType(domain.ResultNotFound)
}
