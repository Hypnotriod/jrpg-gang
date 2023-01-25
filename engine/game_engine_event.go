package engine

import (
	"jrpg-gang/domain"
)

type EndRoundResult struct {
	Damage   map[uint]domain.Damage       `json:"damage"`
	Recovery map[uint]domain.UnitRecovery `json:"recovery"`
	Booty    domain.UnitBooty             `json:"booty"`
}

func NewEndRoundResult() *EndRoundResult {
	result := &EndRoundResult{}
	result.Damage = map[uint]domain.Damage{}
	result.Recovery = map[uint]domain.UnitRecovery{}
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
	Phase            GamePhase             `json:"phase"`
	NextPhase        GamePhase             `json:"nextPhase"`
	PhaseTimeout     float32               `json:"phaseTimeout,omitempty"`
	State            *GameState            `json:"state"`
	Spot             *Spot                 `json:"spot"`
	PlayersInfo      []PlayerInfo          `json:"players"`
	UnitActionResult *GameUnitActionResult `json:"unitActionResult,omitempty"`
	EndRoundResult   *EndRoundResult       `json:"endRoundResult,omitempty"`
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
