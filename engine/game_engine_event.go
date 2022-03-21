package engine

import "jrpg-gang/domain"

type EndTurnResult struct {
	Damage   map[uint]domain.Damage       `json:"damage"`
	Recovery map[uint]domain.UnitRecovery `json:"recovery"`
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

func NewGameUnitActionResult() *GameUnitActionResult {
	result := &GameUnitActionResult{}
	return result
}

type GameEvent struct {
	Phase            GamePhase             `json:"phase"`
	State            *GameState            `json:"state"`
	Spot             *Spot                 `json:"spot"`
	UnitActionResult *GameUnitActionResult `json:"unitActionResult,omitempty"`
	EndRoundResult   *EndTurnResult        `json:"endRoundResult,omitempty"`
}

func (e *GameEngine) NewGameEvent() *GameEvent {
	event := &GameEvent{}
	event.Phase = e.state.Phase
	event.State = e.state
	event.Spot = e.scenario.CurrentSpot()
	return event
}

func (e *GameEngine) NewGameEventWithUnitAction(action *domain.Action) *GameEvent {
	event := &GameEvent{}
	event.Phase = e.state.Phase
	event.UnitActionResult = NewGameUnitActionResult()
	event.UnitActionResult.Action = *action
	event.State = e.state
	event.Spot = e.scenario.CurrentSpot()
	return event
}
