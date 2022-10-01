package config

import (
	"encoding/json"
	"jrpg-gang/domain"
	"jrpg-gang/engine"
	"os"
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

func (c *GameUnitsConfig) Load(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	defer c.mu.Unlock()
	c.mu.Lock()
	units := &[]engine.GameUnit{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(units)
	if err != nil {
		return err
	}
	c.update(units)
	return nil
}

func (c *GameUnitsConfig) update(units *[]engine.GameUnit) {
	c.units = *units
	c.codeToUnit = make(map[domain.UnitCode]*engine.GameUnit)
	for i := range c.units {
		c.codeToUnit[c.units[i].Code] = &c.units[i]
	}
}
