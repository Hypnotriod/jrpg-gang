package engine

import (
	"fmt"
	"jrpg-gang/domain"
	"jrpg-gang/util"
)

type GameScenarioId string

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

func (s *GameScenario) Clone() *GameScenario {
	r := &GameScenario{}
	r.Path = s.Path
	for i := range s.Spots {
		r.Spots = append(r.Spots, *s.Spots[i].Clone())
	}
	return r
}

func (s *GameScenario) Initialize(rndGen *util.RndGen, actors []*GameUnit) {
	s.pathIndex = 0
	s.rndGen = rndGen
	s.prepareActors(actors)
	s.prepareUnits()
}

func (s *GameScenario) PrepareNextSpot(actors []*GameUnit) {
	s.pickSpot()
	s.placeActors(actors)
	s.pathIndex++
}

func (s *GameScenario) Dispose() {
	s.rndGen = nil
	s.spot = nil
}

func (s *GameScenario) IsLastSpot() bool {
	return s.pathIndex >= len(s.Path)
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
	if unit.State.IsEmpty() {
		unit.State.UnitBaseAttributes = unit.Stats.BaseAttributes
	}
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
		actors[i].PlayerInfo.UnitUid = actors[i].Uid
	}
}

func (s *GameScenario) placeActors(actors []*GameUnit) {
	s.spot.Battlefield.placeUnitsDefault(actors)
}
