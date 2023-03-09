package engine

import (
	"jrpg-gang/domain"
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
	GamePhaseRetreatAction      GamePhase = "retreatAction"
	GamePhaseActionComplete     GamePhase = "actionComplete"
	GamePhaseBeforeSpotComplete GamePhase = "beforeSpotComplete"
	GamePhaseSpotComplete       GamePhase = "spotComplete"
	GamePhaseScenarioComplete   GamePhase = "scenarioComplete"
)

type GameState struct {
	ActiveUnitsQueue []uint           `json:"activeUnitsQueue"`
	InactiveUnits    []uint           `json:"inactiveUnits"`
	Booty            domain.UnitBooty `json:"booty"`
	phase            GamePhase
}

func NewGameState() *GameState {
	s := &GameState{}
	s.phase = GamePhasePrepareUnit
	s.ActiveUnitsQueue = make([]uint, 0, 10)
	s.InactiveUnits = make([]uint, 0, 10)
	return s
}

func (s *GameState) MakeUnitsQueue(units []*GameUnit) {
	s.InactiveUnits = make([]uint, 0, len(units))
	s.sortActiveUnitsQueue(units)
	s.ActiveUnitsQueue = util.Map(units, func(unit *GameUnit) uint {
		return unit.Uid
	})
}

func (s *GameState) UpdateUnitsQueue(units []*GameUnit) {
	activeUnits := util.Filter(units, func(u *GameUnit) bool {
		return s.isUnitActive(u.Uid)
	})
	s.sortActiveUnitsQueue(activeUnits)
	s.ActiveUnitsQueue = util.Map(activeUnits, func(unit *GameUnit) uint {
		return unit.Uid
	})
}

func (s *GameState) sortActiveUnitsQueue(units []*GameUnit) {
	sort.SliceStable(units, func(i, j int) bool {
		return units[i].TotalInitiative() > units[j].TotalInitiative()
	})
}

func (s *GameState) PopStunnedUnitFromQueue(unitUid uint) {
	s.ActiveUnitsQueue = util.Filter(s.ActiveUnitsQueue, func(uid uint) bool {
		return unitUid != uid
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
