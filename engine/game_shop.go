package engine

import (
	"jrpg-gang/domain"
	"jrpg-gang/util"
)

type GameShopStatus struct {
	Items    *domain.UnitInventory     `json:"items"`
	Purchase map[uint]domain.UnitBooty `json:"purchase"`
	Repair   map[uint]domain.UnitBooty `json:"repair"`
}

type GameShop struct {
	Items  *domain.UnitInventory
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

func (s *GameShop) GetStatus(unit *domain.Unit) *GameShopStatus {
	r := &GameShopStatus{}
	r.Items = s.Items.Clone()
	r.Purchase = make(map[uint]domain.UnitBooty)
	r.Repair = make(map[uint]domain.UnitBooty)
	for i := range unit.Inventory.Magic {
		if unit.Inventory.Magic[i].CanBeSold {
			r.Purchase[unit.Inventory.Magic[i].Uid] = s.calculatePurchasePrice(&unit.Inventory.Magic[i].Item, 1, 1)
		}
	}
	for i := range unit.Inventory.Disposable {
		if unit.Inventory.Disposable[i].CanBeSold {
			r.Purchase[unit.Inventory.Disposable[i].Uid] = s.calculatePurchasePrice(&unit.Inventory.Disposable[i].Item, 1, 1)
		}
	}
	for i := range unit.Inventory.Ammunition {
		if unit.Inventory.Ammunition[i].CanBeSold {
			r.Purchase[unit.Inventory.Ammunition[i].Uid] = s.calculatePurchasePrice(&unit.Inventory.Ammunition[i].Item, 1, 1)
		}
	}
	for i := range unit.Inventory.Armor {
		if unit.Inventory.Armor[i].CanBeSold {
			r.Purchase[unit.Inventory.Armor[i].Uid] = s.calculatePurchasePrice(
				&unit.Inventory.Armor[i].Item, 1, s.calculateWearoutFactor(&unit.Inventory.Armor[i].Equipment))
		}
		if unit.Inventory.Armor[i].Durability != 0 {
			r.Repair[unit.Inventory.Armor[i].Uid] = s.calculateRepairPrice(&unit.Inventory.Armor[i].Equipment)
		}
	}
	for i := range unit.Inventory.Weapon {
		if unit.Inventory.Weapon[i].CanBeSold {
			r.Purchase[unit.Inventory.Weapon[i].Uid] = s.calculatePurchasePrice(
				&unit.Inventory.Weapon[i].Item, 1, s.calculateWearoutFactor(&unit.Inventory.Weapon[i].Equipment))
		}
		if unit.Inventory.Weapon[i].Durability != 0 {
			r.Repair[unit.Inventory.Weapon[i].Uid] = s.calculateRepairPrice(&unit.Inventory.Weapon[i].Equipment)
		}
	}
	return r
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
	if !item.CanBeSold || item.Type == domain.ItemTypeMagic {
		return result.WithResult(domain.ResultNotAllowed)
	}
	if item.Type == domain.ItemTypeAmmunition {
		if unit.Inventory.RemoveAmmunition(action.ItemUid, action.Quantity) == nil {
			return result.WithResult(domain.ResultNotEnoughResources)
		}
	} else if item.Type == domain.ItemTypeDisposable {
		if unit.Inventory.RemoveDisposable(action.ItemUid, action.Quantity) == nil {
			return result.WithResult(domain.ResultNotEnoughResources)
		}
	} else if action.Quantity != 1 {
		return result.WithResult(domain.ResultNotAllowed)
	} else {
		equipment := unit.Inventory.RemoveEquipment(action.ItemUid)
		if equipment == nil {
			return result.WithResult(domain.ResultNotFound)
		}
		if equipment.Durability <= 0 {
			wearoutFactor = 0
		} else {
			wearoutFactor = 1 - equipment.Wearout/equipment.Durability
		}
	}
	price := s.calculatePurchasePrice(item, action.Quantity, wearoutFactor)
	unit.Booty.Accumulate(price)
	result.Booty = &price
	return result.WithResult(domain.ResultAccomplished)
}

func (s *GameShop) calculateWearoutFactor(equipment *domain.Equipment) float32 {
	if equipment.Durability <= 0 {
		return 0
	}
	return 1 - equipment.Wearout/equipment.Durability
}

func (s *GameShop) calculatePurchasePrice(item *domain.Item, quantity uint, wearoutFactor float32) domain.UnitBooty {
	price := item.Price
	price.MultiplyAll(PURCHASE_ITEM_PRICE_FACTOR * float32(quantity) * wearoutFactor)
	return price
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
	price := s.calculateRepairPrice(equipment)
	if !price.Check(unit.Booty, 1) {
		return domain.NewActionResult().WithResult(domain.ResultNotEnoughResources)
	}
	unit.Booty.Reduce(price, 1)
	equipment.Wearout = 0
	return result.WithResult(domain.ResultAccomplished)
}

func (s *GameShop) calculateRepairPrice(equipment *domain.Equipment) domain.UnitBooty {
	price := equipment.Price
	price.MultiplyAll(REPAIR_ITEM_PRICE_FACTOR * equipment.Wearout / equipment.Durability)
	return price
}
