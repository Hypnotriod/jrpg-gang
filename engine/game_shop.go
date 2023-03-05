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
	s.Items.PopulateUids(s.rndGen)
	s.Items.UnequipAmmunition()
	return s
}

func (s *GameShop) ExecuteAction(action domain.Action, unit *domain.Unit, rndGen *util.RndGen) *domain.ActionResult {
	switch action.Action {
	case domain.ActionBuy:
		return s.buy(action, unit, rndGen)
	case domain.ActionSell:
		return s.sell(action, unit, rndGen)
	case domain.ActionRepair:
		return s.repair(action, unit, rndGen)
	}
	return domain.NewActionResult().WithResult(domain.ResultNotAccomplished)
}

func (s *GameShop) buy(action domain.Action, unit *domain.Unit, rndGen *util.RndGen) *domain.ActionResult {
	if action.Quantity == 0 {
		action.Quantity = 1
	}
	item := s.Items.FindItem(action.ItemUid)
	if item == nil {
		return domain.NewActionResult().WithResult(domain.ResultNotFound)
	}
	if !item.Price.Check(unit.Booty, action.Quantity) {
		return domain.NewActionResult().WithResult(domain.ResultNotEnoughResources)
	}
	if itemRef := s.Items.FindWeapon(action.ItemUid); itemRef != nil {
		itemClone := *itemRef
		itemClone.Uid = rndGen.NextUid()
		unit.Booty.Reduce(item.Price, 1)
		unit.Inventory.Add(&itemClone)
		return domain.NewActionResult().WithResult(domain.ResultAccomplished)
	}
	if itemRef := s.Items.FindMagic(action.ItemUid); itemRef != nil {
		itemClone := *itemRef
		itemClone.Uid = rndGen.NextUid()
		unit.Booty.Reduce(item.Price, 1)
		unit.Inventory.Add(&itemClone)
		return domain.NewActionResult().WithResult(domain.ResultAccomplished)
	}
	if itemRef := s.Items.FindArmor(action.ItemUid); itemRef != nil {
		itemClone := *itemRef
		itemClone.Uid = rndGen.NextUid()
		unit.Booty.Reduce(item.Price, 1)
		unit.Inventory.Add(&itemClone)
		return domain.NewActionResult().WithResult(domain.ResultAccomplished)
	}
	if itemRef := s.Items.FindDisposable(action.ItemUid); itemRef != nil {
		itemClone := *itemRef
		itemClone.Uid = rndGen.NextUid()
		itemClone.Quantity = action.Quantity
		unit.Booty.Reduce(item.Price, action.Quantity)
		unit.Inventory.Add(&itemClone)
		return domain.NewActionResult().WithResult(domain.ResultAccomplished)
	}
	if itemRef := s.Items.FindAmmunition(action.ItemUid); itemRef != nil {
		itemClone := *itemRef
		itemClone.Uid = rndGen.NextUid()
		itemClone.Quantity = action.Quantity
		unit.Booty.Reduce(item.Price, action.Quantity)
		unit.Inventory.Add(&itemClone)
		return domain.NewActionResult().WithResult(domain.ResultAccomplished)
	}
	return domain.NewActionResult().WithResult(domain.ResultNotFound)
}

func (s *GameShop) sell(action domain.Action, unit *domain.Unit, rndGen *util.RndGen) *domain.ActionResult {
	result := domain.NewActionResult()
	var wearoutFactor float32 = 1
	if action.Quantity == 0 {
		action.Quantity = 1
	}
	item := unit.Inventory.FindItem(action.ItemUid)
	if item == nil {
		return result.WithResult(domain.ResultNotFound)
	}
	if !item.CanBeSold {
		return result.WithResult(domain.ResultNotAllowed)
	}
	if item.Type == domain.ItemTypeAmmunition && unit.Inventory.RemoveAmmunition(action.ItemUid, action.Quantity) == nil {
		return result.WithResult(domain.ResultNotEnoughResources)
	} else if item.Type == domain.ItemTypeDisposable && unit.Inventory.RemoveDisposable(action.ItemUid, action.Quantity) == nil {
		return result.WithResult(domain.ResultNotEnoughResources)
	} else if action.Quantity != 1 {
		return result.WithResult(domain.ResultNotAllowed)
	} else if item.Type == domain.ItemTypeMagic && unit.Inventory.RemoveMagic(action.ItemUid) == nil {
		return result.WithResult(domain.ResultNotEnoughResources)
	} else {
		equipment := unit.Inventory.RemoveEquipment(action.ItemUid)
		if equipment == nil {
			return result.WithResult(domain.ResultNotEnoughResources)
		}
		if equipment.Durability <= 0 {
			wearoutFactor = 0
		} else {
			wearoutFactor = 1 - equipment.Wearout/equipment.Durability
		}
	}
	price := item.Price
	price.MultiplyAll(SELL_ITEM_PRICE_FACTOR * float32(action.Quantity) * wearoutFactor)
	unit.Booty.Accumulate(price)
	result.Booty = &price
	return result.WithResult(domain.ResultAccomplished)
}

func (s *GameShop) repair(action domain.Action, unit *domain.Unit, rndGen *util.RndGen) *domain.ActionResult {
	result := domain.NewActionResult()
	item := unit.Inventory.FindItem(action.ItemUid)
	if item == nil {
		return result.WithResult(domain.ResultNotFound)
	}
	equipment := unit.Inventory.FindEquipment(action.ItemUid)
	if equipment == nil || equipment.Wearout == 0 || equipment.Durability == 0 {
		return result.WithResult(domain.ResultNotAllowed)
	}
	price := equipment.Price
	price.MultiplyAll(equipment.Wearout / equipment.Durability)
	if !item.Price.Check(unit.Booty, 1) {
		return domain.NewActionResult().WithResult(domain.ResultNotEnoughResources)
	}
	unit.Booty.Reduce(price, 1)
	equipment.Wearout = 0
	return result.WithResult(domain.ResultAccomplished)
}
