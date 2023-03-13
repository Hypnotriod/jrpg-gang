package engine

import (
	"jrpg-gang/domain"
)

type EndRoundResult struct {
	Damage     map[uint]domain.Damage       `json:"damage,omitempty"`
	Recovery   map[uint]domain.UnitRecovery `json:"recovery,omitempty"`
	Experience map[uint]uint                `json:"experience,omitempty"`
	Drop       map[uint]domain.UnitBooty    `json:"drop,omitempty"`
}

func NewEndRoundResult() *EndRoundResult {
	result := &EndRoundResult{}
	result.Damage = map[uint]domain.Damage{}
	result.Recovery = map[uint]domain.UnitRecovery{}
	result.Experience = map[uint]uint{}
	result.Drop = map[uint]domain.UnitBooty{}
	return result
}

type SpotCompleteResult struct {
	Experience map[uint]uint    `json:"experience,omitempty"`
	Booty      domain.UnitBooty `json:"booty"`
}

func NewSpotCompleteResult() *SpotCompleteResult {
	result := &SpotCompleteResult{}
	result.Experience = map[uint]uint{}
	return result
}

type GameUnitActionResult struct {
	Action domain.Action       `json:"action"`
	Result domain.ActionResult `json:"result"`
}

func NewGameUnitActionResult() *GameUnitActionResult {
	result := &GameUnitActionResult{}
	return result
}

type GameEvent struct {
	Phase              GamePhase             `json:"phase"`
	NextPhase          GamePhase             `json:"nextPhase"`
	PhaseTimeout       float32               `json:"phaseTimeout,omitempty"`
	State              *GameState            `json:"state"`
	Spot               *Spot                 `json:"spot"`
	PlayersInfo        []PlayerInfo          `json:"players"`
	UnitActionResult   *GameUnitActionResult `json:"unitActionResult,omitempty"`
	EndRoundResult     *EndRoundResult       `json:"endRoundResult,omitempty"`
	SpotCompleteResult *SpotCompleteResult   `json:"spotCompleteResult,omitempty"`
}

func (e *GameEvent) WithPhaseTimeout(timeout float32) *GameEvent {
	e.PhaseTimeout = timeout
	return e
}

func (e *GameEngine) NewGameEvent() *GameEvent {
	event := &GameEvent{}
	event.Phase = e.state.phase
	event.NextPhase = e.state.phase
	event.State = e.state
	event.Spot = e.scenario.CurrentSpot()
	return event
}

func (e *GameEngine) NewGameEventWithPlayersInfo() *GameEvent {
	event := &GameEvent{}
	event.Phase = e.state.phase
	event.NextPhase = e.state.phase
	event.State = e.state
	event.Spot = e.scenario.CurrentSpot()
	event.PlayersInfo = e.GetPlayersInfo()
	return event
}

func (e *GameEngine) NewGameEventWithUnitAction(action *domain.Action) *GameEvent {
	event := &GameEvent{}
	event.Phase = e.state.phase
	event.NextPhase = e.state.phase
	event.UnitActionResult = NewGameUnitActionResult()
	event.UnitActionResult.Action = *action
	event.State = e.state
	event.Spot = e.scenario.CurrentSpot()
	return event
}
