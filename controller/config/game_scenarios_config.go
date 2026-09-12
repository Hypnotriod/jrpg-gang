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
	c.scenarios = map[engine.GameScenarioId]*engine.GameScenario{}
	return c
}

func (c *GameScenariosConfig) Scenario(id engine.GameScenarioId) *engine.GameScenario {
	c.mu.RLock()
	defer c.mu.RUnlock()
	scenario, ok := c.scenarios[id]
	if !ok {
		return nil
	}
	return scenario.Clone()
}

func (c *GameScenariosConfig) ScenarioConfig(id engine.GameScenarioId) (engine.GameScenarioConfig, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if scenario, ok := c.scenarios[id]; ok {
		return scenario.Config, ok
	}
	return engine.GameScenarioConfig{}, false
}

func (c *GameScenariosConfig) LoadScenarios(path string, unitsConfig *GameUnitsConfig) error {
	obj := map[engine.GameScenarioId]*engine.GameScenario{}
	scenarios, err := util.ReadJsonFile(&obj, path)
	if err != nil {
		return err
	}
	c.mu.Lock()
	defer c.mu.Unlock()
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
