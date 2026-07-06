package mercenaries

import (
	"jrpg-gang/controller/config"
	"jrpg-gang/domain"
	"jrpg-gang/engine"
	"jrpg-gang/util"
	"sync"
)

type Mercenaries struct {
	mu          sync.RWMutex
	mercenaries *engine.GameMercenaries
}

func NewMercenaries() *Mercenaries {
	m := &Mercenaries{}
	return m
}

func (m *Mercenaries) LoadItems(path string, itemsConfig *config.GameItemsConfig) error {
	mercenaries, err := util.ReadJsonFile(&[]domain.Mercenary{}, path)
	if err != nil {
		return err
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	m.mercenaries = engine.NewGameMercenaries(mercenaries, itemsConfig.PopulateFromDescriptor)
	return nil
}

func (m *Mercenaries) GetStatus() *engine.MercenariesStatus {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.mercenaries.GetStatus()
}

func (m *Mercenaries) Hire(code domain.UnitCode, unit *domain.Unit) *engine.GameUnit {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.mercenaries.Hire(code, unit)
}
