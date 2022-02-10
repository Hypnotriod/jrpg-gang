package engine

import (
	"fmt"
	"jrpg-gang/util"
)

type GameScenario struct {
	Spots     []Spot `json:"spots"`
	spotIndex uint
	uidGen    *util.UidGen
}

func (s GameScenario) String() string {
	return fmt.Sprintf(
		"spots: %v",
		s.Spots,
	)
}

func (s *GameScenario) Initialize() {
	s.spotIndex = 0
	s.uidGen = util.NewUidGen()
	s.asignUids()
}

func (s *GameScenario) asignUids() {
	for i := range s.Spots {
		for _, unit := range s.Spots[i].Battlefield.Units {
			unit.Uid = s.uidGen.Next()
			for j := range unit.Inventory.Ammunition {
				unit.Inventory.Ammunition[j].Uid = s.uidGen.Next()
			}
			for j := range unit.Inventory.Armor {
				unit.Inventory.Armor[j].Uid = s.uidGen.Next()
			}
			for j := range unit.Inventory.Disposable {
				unit.Inventory.Disposable[j].Uid = s.uidGen.Next()
			}
			for j := range unit.Inventory.Magic {
				unit.Inventory.Magic[j].Uid = s.uidGen.Next()
			}
			for j := range unit.Inventory.Weapon {
				unit.Inventory.Weapon[j].Uid = s.uidGen.Next()
			}
		}
	}
}

func (s *GameScenario) CurrentSpot() *Spot {
	return &s.Spots[s.spotIndex]
}

func (s *GameScenario) CurrentBattlefield() *Battlefield {
	return &s.Spots[s.spotIndex].Battlefield
}
