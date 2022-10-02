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
	for _, v := range c.scenarios {
		for n := range v.Spots {
			for _, desc := range v.Spots[n].Battlefield.UnitDescriptor {
				unit := unitsConfig.GetByCode(desc.Code)
				unit.Position = desc.Position
				unit.Faction = desc.Faction
				v.Spots[n].Battlefield.Units = append(v.Spots[n].Battlefield.Units, unit)
			}
		}
	}
}
