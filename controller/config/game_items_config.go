package config

import (
	"encoding/json"
	"jrpg-gang/domain"
	"os"
	"sync"
)

type GameItemsConfig struct {
	mu         sync.RWMutex
	inventory  *domain.UnitInventory
	codeToItem map[domain.ItemCode]any
}

func NewGameItemsConfig() *GameItemsConfig {
	c := &GameItemsConfig{}
	c.codeToItem = make(map[domain.ItemCode]any)
	return c
}

func (c *GameItemsConfig) Load(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	defer c.mu.Unlock()
	c.mu.Lock()
	inventory := &domain.UnitInventory{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(inventory)
	if err != nil {
		return err
	}
	c.prepare(inventory)
	return nil
}

func (c *GameItemsConfig) prepare(inventory *domain.UnitInventory) {
	c.inventory = inventory
	c.codeToItem = make(map[domain.ItemCode]any)
	c.inventory.Prepare()
	c.inventory.PopulateCodeToItemMap(&c.codeToItem)
}
