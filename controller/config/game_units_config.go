package config

import (
	"jrpg-gang/domain"
	"jrpg-gang/engine"
	"jrpg-gang/util"
	"sync"
)

type GameUnitsConfig struct {
	mu         sync.RWMutex
	units      []engine.GameUnit
	codeToUnit map[domain.UnitCode]*engine.GameUnit
}

func NewGameUnitsConfig() *GameUnitsConfig {
	c := &GameUnitsConfig{}
	c.codeToUnit = make(map[domain.UnitCode]*engine.GameUnit)
	return c
}

func (c *GameUnitsConfig) GetByCode(code domain.UnitCode) *engine.GameUnit {
	defer c.mu.RUnlock()
	c.mu.RLock()
	unit, ok := c.codeToUnit[code]
	if !ok {
		return nil
	}
	return unit.Clone()
}

func (c *GameUnitsConfig) LoadUnits(path string, itemsConfig *GameItemsConfig) error {
	units, err := util.ReadJsonFile(&[]engine.GameUnit{}, path)
	if err != nil {
		return err
	}
	defer c.mu.Unlock()
	c.mu.Lock()
	c.update(units, itemsConfig)
	return nil
}

func (c *GameUnitsConfig) update(units *[]engine.GameUnit, itemsConfig *GameItemsConfig) {
	c.units = *units
	c.codeToUnit = make(map[domain.UnitCode]*engine.GameUnit)
	for i := range c.units {
		itemsConfig.PopulateFromDescriptor(&c.units[i].Inventory)
		c.codeToUnit[c.units[i].Code] = &c.units[i]
	}
}
