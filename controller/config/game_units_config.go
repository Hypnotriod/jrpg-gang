package config

import (
	"jrpg-gang/domain"
	"jrpg-gang/engine"
	"jrpg-gang/util"
	"sync"
)

type GameUnitsConfig struct {
	mu          sync.RWMutex
	units       []engine.GameUnit
	codeToUnit  map[domain.UnitCode]*engine.GameUnit
	classToUnit map[domain.UnitClass]*engine.GameUnit
}

func NewGameUnitsConfig() *GameUnitsConfig {
	c := &GameUnitsConfig{}
	c.codeToUnit = make(map[domain.UnitCode]*engine.GameUnit)
	c.classToUnit = make(map[domain.UnitClass]*engine.GameUnit)
	return c
}

func (c *GameUnitsConfig) GetByCode(code domain.UnitCode) *engine.GameUnit {
	c.mu.RLock()
	defer c.mu.RUnlock()
	unit, ok := c.codeToUnit[code]
	if !ok {
		return nil
	}
	return unit.Clone()
}

func (c *GameUnitsConfig) GetByClass(class domain.UnitClass) *engine.GameUnit {
	c.mu.RLock()
	defer c.mu.RUnlock()
	unit, ok := c.classToUnit[class]
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
	c.mu.Lock()
	defer c.mu.Unlock()
	c.update(units, itemsConfig)
	return nil
}

func (c *GameUnitsConfig) update(units *[]engine.GameUnit, itemsConfig *GameItemsConfig) {
	c.units = *units
	c.codeToUnit = make(map[domain.UnitCode]*engine.GameUnit)
	c.classToUnit = make(map[domain.UnitClass]*engine.GameUnit)
	for i := range c.units {
		unit := &c.units[i]
		itemsConfig.PopulateFromDescriptor(&unit.Inventory)
		c.codeToUnit[unit.Code] = unit
		if unit.Class != domain.UnitClassEmpty {
			c.classToUnit[unit.Class] = unit
		}
	}
}
