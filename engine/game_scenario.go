package engine

import (
	"fmt"
	"jrpg-gang/domain"
	"jrpg-gang/util"
)

type GameScenario struct {
	Spots     []Spot        `json:"spots"`
	Path      []map[int]int `json:"path"`
	rndGen    *util.RndGen
	spot      *Spot
	pathIndex int
}

func (s GameScenario) String() string {
	return fmt.Sprintf(
		"spots: %v",
		s.Spots,
	)
}

func (s *GameScenario) Initialize(rndGen *util.RndGen, actors []*GameUnit) {
	s.pathIndex = 0
	s.rndGen = rndGen
	s.prepareActors(actors)
	s.prepareUnits()
	s.pickSpot()
}

func (s *GameScenario) Dispose() {
	s.rndGen = nil
	s.spot = nil
}

func (s *GameScenario) IsLastSpot() bool {
	return s.pathIndex == len(s.Path)-1
}

func (s *GameScenario) NextSpot() {
	s.pathIndex++
	s.pickSpot()
}

func (s *GameScenario) CurrentSpot() *Spot {
	return s.spot
}

func (s *GameScenario) CurrentBattlefield() *Battlefield {
	return &s.spot.Battlefield
}

func (s *GameScenario) pickSpot() {
	spots := s.Path[s.pathIndex]
	index := s.rndGen.PickIntByWeight(spots)
	s.spot = &s.Spots[index]
}

func (s *GameScenario) prepareUnits() {
	for i := range s.Spots {
		for _, unit := range s.Spots[i].Battlefield.Units {
			s.prepareUnit(unit)
		}
	}
}

func (s *GameScenario) prepareUnit(unit *GameUnit) {
	unit.Uid = s.rndGen.NextUid()
	if unit.Damage == nil {
		unit.Damage = []domain.DamageImpact{}
	}
	if unit.Modification == nil {
		unit.Modification = []domain.UnitModificationImpact{}
	}
	unit.Inventory.Prepare()
	unit.Inventory.PopulateUids(s.rndGen)
}

func (s *GameScenario) prepareActors(actors []*GameUnit) {
	for i := range actors {
		s.prepareUnit(actors[i])
	}
}
