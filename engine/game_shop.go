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
	items        *domain.UnitInventory
	requirements map[domain.ItemCode]domain.UnitRequirements
	rndGen       *util.RndGen
}

func NewGameShop(items *domain.UnitInventory, populateFromDescriptor func(inventory *domain.UnitInventory)) *GameShop {
	s := &GameShop{}
	s.requirements = map[domain.ItemCode]domain.UnitRequirements{}
	s.rndGen = util.NewRndGen()
	s.items = items
	s.readItemRequirements(s.items)
	populateFromDescriptor(s.items)
	s.items.PopulateUids(s.rndGen)
	s.items.UnequipAmmunition()
	return s
}

func (s *GameShop) GetStatus(unit *domain.Unit) *GameShopStatus {
	r := &GameShopStatus{}
	r.Items = s.filterItemsByRequirements(unit)
	r.Purchase = map[uint]domain.UnitBooty{}
	r.Repair = map[uint]domain.UnitBooty{}
	for i := range unit.Inventory.Magic {
		if unit.Inventory.Magic[i].CanBeSold {
			r.Purchase[unit.Inventory.Magic[i].Uid] = s.calculatePurchasePrice(unit.Inventory.Magic[i].Price, 1, 1)
		}
	}
	for i := range unit.Inventory.Disposable {
		if unit.Inventory.Disposable[i].CanBeSold {
			r.Purchase[unit.Inventory.Disposable[i].Uid] = s.calculatePurchasePrice(unit.Inventory.Disposable[i].Price, 1, 1)
		}
	}
	for i := range unit.Inventory.Ammunition {
		if unit.Inventory.Ammunition[i].CanBeSold {
			r.Purchase[unit.Inventory.Ammunition[i].Uid] = s.calculatePurchasePrice(unit.Inventory.Ammunition[i].Price, 1, 1)
		}
	}
	for i := range unit.Inventory.Provision {
		if unit.Inventory.Provision[i].CanBeSold {
			r.Purchase[unit.Inventory.Provision[i].Uid] = s.calculatePurchasePrice(unit.Inventory.Provision[i].Price, 1, 1)
		}
	}
	for i := range unit.Inventory.EquipmentUpgrade {
		if unit.Inventory.EquipmentUpgrade[i].CanBeSold {
			r.Purchase[unit.Inventory.EquipmentUpgrade[i].Uid] = s.calculatePurchasePrice(unit.Inventory.EquipmentUpgrade[i].Price, 1, 1)
		}
	}
	for i := range unit.Inventory.Armor {
		if unit.Inventory.Armor[i].CanBeSold {
			r.Purchase[unit.Inventory.Armor[i].Uid] = s.calculatePurchasePrice(
				unit.Inventory.Armor[i].TotalPrice(), 1, s.calculateWearoutFactor(&unit.Inventory.Armor[i].Equipment))
		}
		if unit.Inventory.Armor[i].Durability != 0 {
			r.Repair[unit.Inventory.Armor[i].Uid] = s.calculateRepairPrice(&unit.Inventory.Armor[i].Equipment)
		}
	}
	for i := range unit.Inventory.Weapon {
		if unit.Inventory.Weapon[i].CanBeSold {
			r.Purchase[unit.Inventory.Weapon[i].Uid] = s.calculatePurchasePrice(
				unit.Inventory.Weapon[i].TotalPrice(), 1, s.calculateWearoutFactor(&unit.Inventory.Weapon[i].Equipment))
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
		return s.repair(action, unit)
	case domain.ActionRepairAll:
		return s.repairAll(action, unit)
	case domain.ActionUpgrade:
		return s.upgrade(action, unit)
	case domain.ActionBuyAndUpgrade:
		return s.buyAndUpgrade(action, unit, rndGen)
	}
	return domain.NewActionResult().WithResult(domain.ResultNotAccomplished)
}

func (s *GameShop) buy(action domain.Action, unit *domain.Unit, rndGen *util.RndGen) *domain.ActionResult {
	if action.Quantity == 0 {
		action.Quantity = 1
	}
	item := s.items.FindItem(action.ItemUid)
	if item == nil {
		return domain.NewActionResult().WithResult(domain.ResultNotFound)
	}
	if requirements, ok := s.requirements[item.Code]; ok && !unit.CheckRequirements(requirements) {
		return domain.NewActionResult().WithResult(domain.ResultNotAllowed)
	}
	if !item.Price.Check(unit.Booty, action.Quantity) {
		return domain.NewActionResult().WithResult(domain.ResultNotEnoughResources)
	}
	if itemRef := s.items.FindWeapon(action.ItemUid); itemRef != nil {
		totalPrice := itemRef.TotalPrice()
		if !totalPrice.Check(unit.Booty, action.Quantity) {
			return domain.NewActionResult().WithResult(domain.ResultNotEnoughResources)
		}
		itemClone := *itemRef
		itemClone.Uid = rndGen.NextUid()
		unit.Booty.Reduce(totalPrice, 1)
		unit.Inventory.Add(&itemClone)
		return domain.NewActionResult().WithResult(domain.ResultAccomplished)
	}
	if itemRef := s.items.FindMagic(action.ItemUid); itemRef != nil {
		itemClone := *itemRef
		itemClone.Uid = rndGen.NextUid()
		unit.Booty.Reduce(item.Price, 1)
		unit.Inventory.Add(&itemClone)
		return domain.NewActionResult().WithResult(domain.ResultAccomplished)
	}
	if itemRef := s.items.FindArmor(action.ItemUid); itemRef != nil {
		totalPrice := itemRef.TotalPrice()
		if !totalPrice.Check(unit.Booty, action.Quantity) {
			return domain.NewActionResult().WithResult(domain.ResultNotEnoughResources)
		}
		itemClone := *itemRef
		itemClone.Uid = rndGen.NextUid()
		unit.Booty.Reduce(totalPrice, 1)
		unit.Inventory.Add(&itemClone)
		return domain.NewActionResult().WithResult(domain.ResultAccomplished)
	}
	if itemRef := s.items.FindDisposable(action.ItemUid); itemRef != nil {
		itemClone := *itemRef
		itemClone.Uid = rndGen.NextUid()
		itemClone.Quantity = action.Quantity
		unit.Booty.Reduce(item.Price, action.Quantity)
		unit.Inventory.Add(&itemClone)
		return domain.NewActionResult().WithResult(domain.ResultAccomplished)
	}
	if itemRef := s.items.FindAmmunition(action.ItemUid); itemRef != nil {
		itemClone := *itemRef
		itemClone.Uid = rndGen.NextUid()
		itemClone.Quantity = action.Quantity
		unit.Booty.Reduce(item.Price, action.Quantity)
		unit.Inventory.Add(&itemClone)
		return domain.NewActionResult().WithResult(domain.ResultAccomplished)
	}
	if itemRef := s.items.FindProvision(action.ItemUid); itemRef != nil {
		itemClone := *itemRef
		itemClone.Uid = rndGen.NextUid()
		itemClone.Quantity = action.Quantity
		unit.Booty.Reduce(item.Price, action.Quantity)
		unit.Inventory.Add(&itemClone)
		return domain.NewActionResult().WithResult(domain.ResultAccomplished)
	}
	if itemRef := s.items.FindEquipmentUpgrade(action.ItemUid); itemRef != nil {
		itemClone := *itemRef
		itemClone.Uid = rndGen.NextUid()
		unit.Booty.Reduce(item.Price, 1)
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
	itemPrice := item.Price
	if item.Type == domain.ItemTypeAmmunition {
		if unit.Inventory.RemoveAmmunition(action.ItemUid, action.Quantity) == nil {
			return result.WithResult(domain.ResultNotEnoughResources)
		}
	} else if item.Type == domain.ItemTypeDisposable {
		if unit.Inventory.RemoveDisposable(action.ItemUid, action.Quantity) == nil {
			return result.WithResult(domain.ResultNotEnoughResources)
		}
	} else if item.Type == domain.ItemTypeProvision {
		if unit.Inventory.RemoveProvision(action.ItemUid, action.Quantity) == nil {
			return result.WithResult(domain.ResultNotEnoughResources)
		}
	} else if action.Quantity != 1 {
		return result.WithResult(domain.ResultNotAllowed)
	} else if item.Type == domain.ItemTypeEquipmentUpgrade {
		if unit.Inventory.RemoveEquipmentUpgrade(action.ItemUid) == nil {
			return result.WithResult(domain.ResultNotFound)
		}
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
		itemPrice = equipment.TotalPrice()
	}
	price := s.calculatePurchasePrice(itemPrice, action.Quantity, wearoutFactor)
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

func (s *GameShop) calculatePurchasePrice(price domain.UnitBooty, quantity uint, wearoutFactor float32) domain.UnitBooty {
	price.MultiplyAll(PURCHASE_ITEM_PRICE_FACTOR * wearoutFactor)
	price.MultiplyAll((float32(quantity)))
	return price
}

func (s *GameShop) repair(action domain.Action, unit *domain.Unit) *domain.ActionResult {
	result := domain.NewActionResult()
	equipment := unit.Inventory.FindEquipment(action.ItemUid)
	if equipment == nil {
		return result.WithResult(domain.ResultNotFound)
	}
	if equipment.Wearout == 0 || equipment.Durability == 0 {
		return result.WithResult(domain.ResultNotAllowed)
	}
	price := s.calculateRepairPrice(equipment)
	if !price.Check(unit.Booty, 1) {
		return result.WithResult(domain.ResultNotEnoughResources)
	}
	unit.Booty.Reduce(price, 1)
	equipment.ResetWearout()
	return result.WithResult(domain.ResultAccomplished)
}

func (s *GameShop) repairAll(action domain.Action, unit *domain.Unit) *domain.ActionResult {
	var price domain.UnitBooty
	result := domain.NewActionResult()
	equipment := unit.Inventory.GetWornEquipment()
	if len(equipment) == 0 {
		return result.WithResult(domain.ResultNotFound)
	}
	for _, e := range equipment {
		price.Accumulate(s.calculateRepairPrice(e))
	}
	if !price.Check(unit.Booty, 1) {
		return result.WithResult(domain.ResultNotEnoughResources)
	}
	unit.Booty.Reduce(price, 1)
	for _, e := range equipment {
		e.ResetWearout()
	}
	return result.WithResult(domain.ResultAccomplished)
}

func (s *GameShop) upgrade(action domain.Action, unit *domain.Unit) *domain.ActionResult {
	result := domain.NewActionResult()
	equipment := unit.Inventory.FindEquipment(action.ItemUid)
	upgrade := unit.Inventory.FindEquipmentUpgrade(action.Uid)
	if upgrade == nil || equipment == nil {
		return result.WithResult(domain.ResultNotFound)
	}
	if len(equipment.Upgrades) >= int(equipment.UpgradesLimit) {
		return result.WithResult(domain.ResultNotAllowed)
	}
	if (upgrade.Damage != nil || upgrade.UseCost != nil) && equipment.Type != domain.ItemTypeWeapon {
		return result.WithResult(domain.ResultNotAllowed)
	}
	equipment.Upgrades = append(equipment.Upgrades, *upgrade)
	unit.Inventory.RemoveEquipmentUpgrade(action.Uid)
	return result.WithResult(domain.ResultAccomplished)
}

func (s *GameShop) buyAndUpgrade(action domain.Action, unit *domain.Unit, rndGen *util.RndGen) *domain.ActionResult {
	result := domain.NewActionResult()
	equipment := unit.Inventory.FindEquipment(action.ItemUid)
	upgrade := s.items.FindEquipmentUpgrade(action.Uid)
	if upgrade == nil || equipment == nil {
		return result.WithResult(domain.ResultNotFound)
	}
	if requirements, ok := s.requirements[upgrade.Code]; ok && !unit.CheckRequirements(requirements) {
		return domain.NewActionResult().WithResult(domain.ResultNotAllowed)
	}
	if !upgrade.Price.Check(unit.Booty, 1) {
		return result.WithResult(domain.ResultNotEnoughResources)
	}
	if len(equipment.Upgrades) >= int(equipment.UpgradesLimit) {
		return result.WithResult(domain.ResultNotAllowed)
	}
	if (upgrade.Damage != nil || upgrade.UseCost != nil) && equipment.Type != domain.ItemTypeWeapon {
		return result.WithResult(domain.ResultNotAllowed)
	}
	upgradeClone := *upgrade
	upgradeClone.Uid = rndGen.NextUid()
	unit.Booty.Reduce(upgrade.Price, 1)
	equipment.Upgrades = append(equipment.Upgrades, upgradeClone)
	return result.WithResult(domain.ResultAccomplished)
}

func (s *GameShop) calculateRepairPrice(equipment *domain.Equipment) domain.UnitBooty {
	price := equipment.TotalPrice()
	price.MultiplyAll(REPAIR_ITEM_PRICE_FACTOR * equipment.Wearout / equipment.Durability)
	return price
}

func (s *GameShop) readItemRequirements(items *domain.UnitInventory) {
	for _, d := range items.Descriptor {
		if d.Requirements != nil {
			s.requirements[d.Code] = *d.Requirements
		}
	}
}

func (s *GameShop) filterItemsByRequirements(unit *domain.Unit) *domain.UnitInventory {
	r := &domain.UnitInventory{}
	attributes := unit.TotalModification().Attributes
	attributes.Accumulate(unit.Stats.Attributes)
	attributes.Normalize()
	r.Weapon = util.Filter(s.items.Weapon, func(item domain.Weapon) bool {
		requirements, ok := s.requirements[item.Code]
		return !ok || requirements.Check(unit, attributes)
	})
	r.Magic = util.Filter(s.items.Magic, func(item domain.Magic) bool {
		requirements, ok := s.requirements[item.Code]
		return !ok || requirements.Check(unit, attributes)
	})
	r.Armor = util.Filter(s.items.Armor, func(item domain.Armor) bool {
		requirements, ok := s.requirements[item.Code]
		return !ok || requirements.Check(unit, attributes)
	})
	r.Disposable = util.Filter(s.items.Disposable, func(item domain.Disposable) bool {
		requirements, ok := s.requirements[item.Code]
		return !ok || requirements.Check(unit, attributes)
	})
	r.Ammunition = util.Filter(s.items.Ammunition, func(item domain.Ammunition) bool {
		requirements, ok := s.requirements[item.Code]
		return !ok || requirements.Check(unit, attributes)
	})
	r.Provision = util.Filter(s.items.Provision, func(item domain.Provision) bool {
		requirements, ok := s.requirements[item.Code]
		return !ok || requirements.Check(unit, attributes)
	})
	r.EquipmentUpgrade = util.Filter(s.items.EquipmentUpgrade, func(item domain.EquipmentUpgrade) bool {
		requirements, ok := s.requirements[item.Code]
		return !ok || requirements.Check(unit, attributes)
	})
	return r
}
