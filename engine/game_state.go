package engine

import (
	"fmt"
	"jrpg-gang/util"
	"sort"
)

type GameState struct {
	ActiveUnitsQueue []uint `json:"activeUnitsQueue"`
	InactiveUnits    []uint `json:"inactiveUnits"`
}

func NewGameState() *GameState {
	s := &GameState{}
	s.ActiveUnitsQueue = make([]uint, 0, 10)
	s.InactiveUnits = make([]uint, 0, 10)
	return s
}

func (s GameState) String() string {
	return fmt.Sprintf(
		"active units queue: {%v}, inactive units: {%v}",
		util.AsCommaSeparatedSlice(s.ActiveUnitsQueue),
		util.AsCommaSeparatedSlice(s.InactiveUnits),
	)
}

func (s *GameState) MakeUnitsQueue(units []*GameUnit) {
	s.ActiveUnitsQueue = make([]uint, 0, len(units))
	s.InactiveUnits = make([]uint, 0, len(units))
	for _, unit := range units {
		s.ActiveUnitsQueue = append(s.ActiveUnitsQueue, unit.Uid)
	}
	sort.SliceStable(s.ActiveUnitsQueue, func(i, j int) bool {
		return units[i].TotalInitiative() < units[j].TotalInitiative()
	})
}

func (s *GameState) UpdateUnitsQueue(units []*GameUnit) {
	activeUnits := []*GameUnit{}
	for _, unit := range units {
		if s.IsUnitActive(unit.Uid) {
			activeUnits = append(activeUnits, unit)
		}
	}
	s.ActiveUnitsQueue = make([]uint, 0, len(activeUnits))
	sort.SliceStable(activeUnits, func(i, j int) bool {
		return units[i].TotalInitiative() < units[j].TotalInitiative()
	})
}

func (s *GameState) IsUnitActive(uid uint) bool {
	for _, unitUid := range s.ActiveUnitsQueue {
		if unitUid == uid {
			return true
		}
	}
	return false
}

func (s *GameState) ShiftUnitsQueue() {
	if len(s.ActiveUnitsQueue) != 0 {
		s.InactiveUnits = append(s.InactiveUnits, s.ActiveUnitsQueue[0])
		s.ActiveUnitsQueue = s.ActiveUnitsQueue[1:]
	}
}

func (s *GameState) GetActiveUnitUid() (uint, bool) {
	if len(s.ActiveUnitsQueue) == 0 {
		return 0, false
	}
	return s.ActiveUnitsQueue[0], true
}
