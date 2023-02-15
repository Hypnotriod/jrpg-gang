package config

import (
	"jrpg-gang/engine"
	"jrpg-gang/util"
	"sync"
)

type GameScenariosConfig struct {
	mu        sync.RWMutex
	scenarios map[engine.GameScenarioId]*engine.GameScenario
}

func NewGameScenariosConfig() *GameScenariosConfig {
	c := &GameScenariosConfig{}
	c.scenarios = make(map[engine.GameScenarioId]*engine.GameScenario)
	return c
}

func (c *GameScenariosConfig) Has(id engine.GameScenarioId) bool {
	defer c.mu.RUnlock()
	c.mu.RLock()
	_, ok := c.scenarios[id]
	return ok
}

func (c *GameScenariosConfig) Get(id engine.GameScenarioId) *engine.GameScenario {
	defer c.mu.RUnlock()
	c.mu.RLock()
	scenario, ok := c.scenarios[id]
	if !ok {
		return nil
	}
	return scenario.Clone()
}

func (c *GameScenariosConfig) LoadScenarios(path string, unitsConfig *GameUnitsConfig) error {
	obj := make(map[engine.GameScenarioId]*engine.GameScenario)
	scenarios, err := util.ReadJsonFile(&obj, path)
	if err != nil {
		return err
	}
	defer c.mu.Unlock()
	c.mu.Lock()
	c.prepare(scenarios, unitsConfig)
	return nil
}

func (c *GameScenariosConfig) prepare(scenarios *map[engine.GameScenarioId]*engine.GameScenario, unitsConfig *GameUnitsConfig) {
	c.scenarios = *scenarios
	for _, scenario := range c.scenarios {
		for _, spot := range scenario.Spots {
			for _, desc := range spot.Battlefield.UnitDescriptor {
				unit := unitsConfig.GetByCode(desc.Code)
				unit.Position = desc.Position
				unit.Faction = desc.Faction
				spot.Battlefield.Units = append(spot.Battlefield.Units, unit)
			}
			spot.Battlefield.UnitDescriptor = nil
		}
	}
}
