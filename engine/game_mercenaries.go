package engine

import (
	"jrpg-gang/domain"
	"jrpg-gang/util"
)

type MercenariesStatus struct {
	Mercenaries *[]domain.Mercenary `json:"mercenaries"`
}

type GameMercenaries struct {
	mercenaries *[]domain.Mercenary
}

func NewGameMercenaries(mercenaries *[]domain.Mercenary, populateFromDescriptor func(inventory *domain.UnitInventory)) *GameMercenaries {
	m := &GameMercenaries{
		mercenaries: mercenaries,
	}
	for i := range *mercenaries {
		inventory := &(*mercenaries)[i].Inventory
		populateFromDescriptor(inventory)
	}
	return m
}

func (s *GameMercenaries) GetStatus() *MercenariesStatus {
	r := &MercenariesStatus{}
	r.Mercenaries = s.mercenaries
	return r
}

func (s *GameMercenaries) Hire(code domain.UnitCode, unit *domain.Unit) *GameUnit {
	mercenary := util.Find(*s.mercenaries, func(mercenary domain.Mercenary) bool {
		return mercenary.Code == code
	})
	if mercenary == nil {
		return nil
	}
	if mercenary.Requirements != nil && !unit.CheckRequirements(*mercenary.Requirements) {
		return nil
	}
	if mercenary.Price != nil {
		if !mercenary.Price.Check(unit.Booty, 1) {
			return nil
		}
		unit.Booty.Reduce(*mercenary.Price, 1)
	}
	return NewGameUnit(&mercenary.Unit)
}
