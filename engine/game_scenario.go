package engine

import (
	"jrpg-gang/domain"
	"jrpg-gang/util"
)

type GameScenarioId string
type GameSpotId string

type GamePath struct {
	W      int        `json:"weight"`
	SpotId GameSpotId `json:"spotId"`
}

func (p GamePath) Weight() int {
	return p.W
}

type GameScenario struct {
	Spots        map[GameSpotId]*Spot    `json:"spots"`
	Path         [][]GamePath            `json:"path"`
	Requirements domain.UnitRequirements `json:"requirements"`
	rndGen       *util.RndGen
	spot         *Spot
	pathIndex    int
}

func (s *GameScenario) Clone() *GameScenario {
	r := &GameScenario{}
	r.Path = s.Path
	r.Spots = make(map[GameSpotId]*Spot)
	for id, spot := range s.Spots {
		r.Spots[id] = spot.Clone()
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
	path := util.RandomPick(s.rndGen, spots)
	s.spot = s.Spots[path.SpotId]
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
	unit.State.RestoreDefault(unit.Stats.BaseAttributes)
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
