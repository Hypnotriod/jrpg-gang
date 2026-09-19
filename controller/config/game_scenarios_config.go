package config

import (
	"jrpg-gang/engine"
	"jrpg-gang/util"
	"sync"
)

type GameScenariosConfig struct {
	mu        sync.RWMutex
	scenarios []*engine.GameScenario
}

func NewGameScenariosConfig() *GameScenariosConfig {
	c := &GameScenariosConfig{}
	c.scenarios = []*engine.GameScenario{}
	return c
}

func (c *GameScenariosConfig) GetScenario(id engine.GameScenarioId) *engine.GameScenario {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if scenario := util.FindPtr(c.scenarios, func(value *engine.GameScenario) bool {
		return value.Config.Id == id
	}); scenario != nil {
		return scenario.Clone()
	}
	return nil
}

func (c *GameScenariosConfig) GetScenarioConfig(id engine.GameScenarioId) (engine.GameScenarioConfig, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if scenario := util.FindPtr(c.scenarios, func(value *engine.GameScenario) bool {
		return value.Config.Id == id
	}); scenario != nil {
		return scenario.Config, true
	}
	return engine.GameScenarioConfig{}, false
}

func (c *GameScenariosConfig) GetAllScenarioConfigs() []engine.GameScenarioConfig {
	c.mu.RLock()
	defer c.mu.RUnlock()
	configs := []engine.GameScenarioConfig{}
	for _, scenario := range c.scenarios {
		configs = append(configs, scenario.Config)
	}
	return configs
}

func (c *GameScenariosConfig) GetAvailableScenarioConfigs(unit *engine.GameUnit) []engine.GameScenarioConfig {
	c.mu.RLock()
	defer c.mu.RUnlock()
	configs := []engine.GameScenarioConfig{}
	for _, scenario := range c.scenarios {
		if !unit.Quests.Test(scenario.Config.Requirements.Quests) ||
			!unit.Achievements.Test(scenario.Config.Requirements.Achievements) {
			continue
		}
		configs = append(configs, scenario.Config)
	}
	return configs
}

func (c *GameScenariosConfig) LoadScenarios(path string, unitsConfig *GameUnitsConfig) error {
	scenarios, err := util.ReadJsonFile(&[]*engine.GameScenario{}, path)
	if err != nil {
		return err
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.prepare(*scenarios, unitsConfig)
	return nil
}

func (c *GameScenariosConfig) prepare(scenarios []*engine.GameScenario, unitsConfig *GameUnitsConfig) {
	c.scenarios = scenarios
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
