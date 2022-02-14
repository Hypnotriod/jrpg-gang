package engine

import (
	"fmt"
	"jrpg-gang/util"
)

type GameScenario struct {
	Spots     []Spot  `json:"spots"`
	Path      [][]int `json:"path"`
	spot      *Spot
	pathIndex int
	rndGen    *util.RndGen
}

func (s GameScenario) String() string {
	return fmt.Sprintf(
		"spots: %v",
		s.Spots,
	)
}

func (s *GameScenario) Initialize(actors []*GameUnit) {
	s.pathIndex = 0
	s.rndGen = util.NewRndGen()
	s.prepareMainActors(actors)
	s.asignUids()
	s.pickSpot()
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
	index := s.rndGen.PickInt(spots)
	s.spot = &s.Spots[index]
}

func (s *GameScenario) asignUids() {
	for i := range s.Spots {
		for _, unit := range s.Spots[i].Battlefield.Units {
			s.prepareUnit(unit)
		}
	}
}

func (s *GameScenario) prepareUnit(unit *GameUnit) {
	unit.Uid = s.rndGen.NextUid()
	for j := range unit.Inventory.Ammunition {
		unit.Inventory.Ammunition[j].Uid = s.rndGen.NextUid()
	}
	for j := range unit.Inventory.Armor {
		unit.Inventory.Armor[j].Uid = s.rndGen.NextUid()
	}
	for j := range unit.Inventory.Disposable {
		unit.Inventory.Disposable[j].Uid = s.rndGen.NextUid()
	}
	for j := range unit.Inventory.Magic {
		unit.Inventory.Magic[j].Uid = s.rndGen.NextUid()
	}
	for j := range unit.Inventory.Weapon {
		unit.Inventory.Weapon[j].Uid = s.rndGen.NextUid()
	}
}

func (s *GameScenario) prepareMainActors(actors []*GameUnit) {
	for i := range actors {
		s.prepareUnit(actors[i])
	}
}
