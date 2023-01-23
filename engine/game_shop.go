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
	var wearoutFactor float32 = 1
	if action.Quantity == 0 {
		action.Quantity = 1
	}
	item := unit.Inventory.FindItem(action.ItemUid)
	if item == nil {
		return domain.NewActionResult().WithResult(domain.ResultNotFound)
	}
	if !item.CanBeSold {
		return domain.NewActionResult().WithResult(domain.ResultNotAllowed)
	}
	if amm := unit.Inventory.FindAmmunition(action.ItemUid); amm != nil {
		if amm.Quantity < action.Quantity {
			return domain.NewActionResult().WithResult(domain.ResultNotEnoughResources)
		}
		amm.Quantity -= action.Quantity
		unit.Inventory.FilterAmmunition()
	} else if disp := unit.Inventory.FindDisposable(action.ItemUid); disp != nil {
		if disp.Quantity < action.Quantity {
			return domain.NewActionResult().WithResult(domain.ResultNotEnoughResources)
		}
		disp.Quantity -= action.Quantity
		unit.Inventory.FilterDisposable()
	} else if action.Quantity != 1 {
		return domain.NewActionResult().WithResult(domain.ResultNotAllowed)
	} else if equip := unit.Inventory.FindEquipment(action.ItemUid); equip != nil {
		if equip.Durability <= 0 {
			wearoutFactor = 0
		} else {
			wearoutFactor = 1 - equip.Wearout/equip.Durability
		}
		unit.Inventory.RemoveArmor(item.Uid)
		unit.Inventory.RemoveWeapon(item.Uid)
	} else {
		unit.Inventory.RemoveMagic(item.Uid)
	}
	price := item.Price
	price.MultiplyAll(SELL_ITEM_PRICE_FACTOR * float32(action.Quantity) * wearoutFactor)
	unit.Booty.Accumulate(price)
	return domain.NewActionResult().WithResult(domain.ResultAccomplished)
}
