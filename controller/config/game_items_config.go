package config

import (
	"jrpg-gang/domain"
	"jrpg-gang/util"
	"sync"
)

type GameItemsConfig struct {
	mu         sync.RWMutex
	items      *domain.UnitInventory
	codeToItem map[domain.ItemCode]any
}

func NewGameItemsConfig() *GameItemsConfig {
	c := &GameItemsConfig{}
	c.codeToItem = make(map[domain.ItemCode]any)
	return c
}

func (c *GameItemsConfig) Load(path string) error {
	items, err := util.ReadJsonFile(&domain.UnitInventory{}, path)
	if err != nil {
		return err
	}
	defer c.mu.Unlock()
	c.mu.Lock()
	c.prepare(items)
	return nil
}

func (c *GameItemsConfig) prepare(items *domain.UnitInventory) {
	c.items = items
	c.codeToItem = make(map[domain.ItemCode]any)
	c.items.UnequipAmmunition()
	c.items.PopulateCodeToItemMap(&c.codeToItem)
}

func (c *GameItemsConfig) PopulateFromDescriptor(inventory *domain.UnitInventory) {
	defer c.mu.RUnlock()
	c.mu.RLock()
	for n := range inventory.Descriptor {
		desc := inventory.Descriptor[n]
		if item, ok := c.codeToItem[desc.Code]; ok {
			switch v := item.(type) {
			case *domain.Weapon:
				itemClone := *v
				itemClone.Equipped = desc.Equipped
				inventory.Add(&itemClone)
			case *domain.Magic:
				itemClone := *v
				inventory.Add(&itemClone)
			case *domain.Armor:
				itemClone := *v
				itemClone.Equipped = desc.Equipped
				inventory.Add(&itemClone)
			case *domain.Disposable:
				itemClone := *v
				itemClone.Quantity = desc.Quantity
				inventory.Add(&itemClone)
			case *domain.Ammunition:
				itemClone := *v
				itemClone.Quantity = desc.Quantity
				itemClone.Equipped = desc.Equipped
				inventory.Add(&itemClone)
			}
		}
	}
	inventory.Descriptor = []domain.UnitInventoryDescriptor{}
}
