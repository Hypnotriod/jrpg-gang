package engine

import (
	"fmt"
	"jrpg-gang/domain"
)

type EndTurnResult struct {
	Damage   map[uint]domain.Damage       `json:"damage"`
	Recovery map[uint]domain.UnitRecovery `json:"recovery"`
}

func (r *EndTurnResult) String() string {
	return fmt.Sprintf(
		"damage: {%v}, recovery: {%v}",
		r.Damage,
		r.Recovery,
	)
}

func NewEndTurnResult() *EndTurnResult {
	result := &EndTurnResult{}
	result.Damage = map[uint]domain.Damage{}
	result.Recovery = map[uint]domain.UnitRecovery{}
	return result
}

type GameUnitActionResult struct {
	Action domain.Action       `json:"action"`
	Result domain.ActionResult `json:"result"`
}

func (r *GameUnitActionResult) String() string {
	return fmt.Sprintf(
		"action: {%v}, result: {%v}",
		r.Action,
		r.Result,
	)
}

func NewGameUnitActionResult() *GameUnitActionResult {
	result := &GameUnitActionResult{}
	return result
}

type GameEvent struct {
	Phase            GamePhase             `json:"phase"`
	NextPhase        GamePhase             `json:"nextPhase"`
	State            *GameState            `json:"state"`
	Spot             *Spot                 `json:"spot"`
	UnitActionResult *GameUnitActionResult `json:"unitActionResult,omitempty"`
	EndRoundResult   *EndTurnResult        `json:"endRoundResult,omitempty"`
}

func (e *GameEvent) String() string {
	return fmt.Sprintf(
		"phase: %s, next phase: %s, state: {%v}, spot: {%v}, unitActionResult: {%v}, endRoundResult: {%v}",
		e.Phase,
		e.NextPhase,
		e.State,
		e.Spot,
		e.UnitActionResult,
		e.EndRoundResult,
	)
}

func (e *GameEngine) NewGameEvent() *GameEvent {
	event := &GameEvent{}
	event.Phase = e.state.phase
	event.NextPhase = e.state.phase
	event.State = e.state
	event.Spot = e.scenario.CurrentSpot()
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
