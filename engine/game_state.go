package engine

import (
	"fmt"
	"jrpg-gang/util"
	"sort"
)

type GamePhase string

const (
	GamePhasePrepareUnit        GamePhase = "prepareUnit"
	GamePhaseReadyForStartRound GamePhase = "readyForStartRound"
	GamePhaseMakeMoveOrActionAI GamePhase = "makeMoveOrActionAI"
	GamePhaseMakeActionAI       GamePhase = "makeActionAI"
	GamePhaseMakeMoveOrAction   GamePhase = "makeMoveOrAction"
	GamePhaseMakeAction         GamePhase = "makeAction"
	GamePhaseActionComplete     GamePhase = "actionComplete"
	GamePhaseBattleComplete     GamePhase = "battleComplete"
)

type GameState struct {
	ActiveUnitsQueue []uint `json:"activeUnitsQueue"`
	InactiveUnits    []uint `json:"inactiveUnits"`
	phase            GamePhase
}

func NewGameState() *GameState {
	s := &GameState{}
	s.phase = GamePhasePrepareUnit
	s.ActiveUnitsQueue = make([]uint, 0, 10)
	s.InactiveUnits = make([]uint, 0, 10)
	return s
}

func (s GameState) String() string {
	return fmt.Sprintf(
		"active units queue: [%v], inactive units: [%v], phase: %s",
		util.AsCommaSeparatedSlice(s.ActiveUnitsQueue),
		util.AsCommaSeparatedSlice(s.InactiveUnits),
		s.phase,
	)
}

func (s *GameState) MakeUnitsQueue(units []*GameUnit) {
	s.InactiveUnits = make([]uint, 0, len(units))
	s.ActiveUnitsQueue = util.Map(units, func(unit *GameUnit) uint {
		return unit.Uid
	})
	s.sortActiveUnitsQueue(units)
}

func (s *GameState) UpdateUnitsQueue(units []*GameUnit) {
	activeUnits := util.Filter(units, func(u *GameUnit) bool {
		return s.isUnitActive(u.Uid)
	})
	s.ActiveUnitsQueue = util.Map(activeUnits, func(unit *GameUnit) uint {
		return unit.Uid
	})
	s.sortActiveUnitsQueue(activeUnits)
}

func (s *GameState) sortActiveUnitsQueue(units []*GameUnit) {
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
	s.phase = phase
}

func (s *GameState) HasActiveUnits() bool {
	return len(s.ActiveUnitsQueue) != 0
}

func (s *GameState) isUnitActive(uid uint) bool {
	return util.Contains(s.ActiveUnitsQueue, uid)
}
