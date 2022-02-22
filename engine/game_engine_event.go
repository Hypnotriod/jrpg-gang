package engine

import "jrpg-gang/domain"

type EndTurnResult struct {
	Damage   map[uint]domain.Damage       `json:"damage"`
	Recovery map[uint]domain.UnitRecovery `json:"recovery"`
}

type GameUnitActionResult struct {
	Action GameAction          `json:"action"`
	Result domain.ActionResult `json:"result"`
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

func (e *GameEngine) NewGameEventWithUnitAction(action *GameAction) *GameEvent {
	event := &GameEvent{}
	event.Phase = e.state.Phase
	event.UnitActionResult = &GameUnitActionResult{
		Action: *action,
	}
	event.State = e.state
	event.Spot = e.scenario.CurrentSpot()
	return event
}
