package engine

import (
	"fmt"
	"jrpg-gang/util"
	"sort"
)

type GamePhase string

const (
	GamePhaseReadyToPlaceUnit   GamePhase = "readyToPlaceUnit"
	GamePhasePlaceUnit          GamePhase = "placeUnit"
	GamePhaseReadyForStartRound GamePhase = "readyForStartRound"
	GamePhaseMakeMoveOrActionAI GamePhase = "makeMoveOrActionAI"
	GamePhaseMakeActionAI       GamePhase = "makeActionAI"
	GamePhaseMakeMoveOrAction   GamePhase = "makeMoveOrAction"
	GamePhaseMakeAction         GamePhase = "makeAction"
	GamePhaseActionComplete     GamePhase = "actionComplete"
	GamePhaseBattleComplete     GamePhase = "battleComplete"
)

type GameState struct {
	Phase            GamePhase `json:"phase"`
	ActiveUnitsQueue []uint    `json:"activeUnitsQueue"`
	InactiveUnits    []uint    `json:"inactiveUnits"`
}

func NewGameState() *GameState {
	s := &GameState{}
	s.Phase = GamePhaseReadyToPlaceUnit
	s.ActiveUnitsQueue = make([]uint, 0, 10)
	s.InactiveUnits = make([]uint, 0, 10)
	return s
}

func (s GameState) String() string {
	return fmt.Sprintf(
		"active units queue: {%v}, inactive units: {%v}, phase: %s",
		util.AsCommaSeparatedSlice(s.ActiveUnitsQueue),
		util.AsCommaSeparatedSlice(s.InactiveUnits),
		s.Phase,
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
	if len(s.ActiveUnitsQueue) == 0 {
		return
	}
	activeUnits := []*GameUnit{}
	for _, unit := range units {
		if s.isUnitActive(unit.Uid) {
			activeUnits = append(activeUnits, unit)
		}
	}
	s.ActiveUnitsQueue = make([]uint, 0, len(activeUnits))
	sort.SliceStable(s.ActiveUnitsQueue, func(i, j int) bool {
		return units[i].TotalInitiative() < units[j].TotalInitiative()
	})
}

func (s *GameState) ShiftUnitsQueue() {
	if len(s.ActiveUnitsQueue) != 0 {
		s.InactiveUnits = append(s.InactiveUnits, s.ActiveUnitsQueue[0])
		s.ActiveUnitsQueue = s.ActiveUnitsQueue[1:]
	}
}

func (s *GameState) GetCurrentActiveUnitUid() (uint, bool) {
	if len(s.ActiveUnitsQueue) == 0 {
		return 0, false
	}
	return s.ActiveUnitsQueue[0], true
}

func (s *GameState) IsCurrentActiveUnit(unit *GameUnit) bool {
	uid, ok := s.GetCurrentActiveUnitUid()
	return ok && unit.Uid == uid
}

func (s *GameState) ChangePhase(phase GamePhase) {
	s.Phase = phase
}

func (s *GameState) HasActiveUnits() bool {
	return len(s.ActiveUnitsQueue) != 0
}

func (s *GameState) isUnitActive(uid uint) bool {
	for _, unitUid := range s.ActiveUnitsQueue {
		if unitUid == uid {
			return true
		}
	}
	return false
}
