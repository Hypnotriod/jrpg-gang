package engine

import "jrpg-gang/domain"

type EndTurnResult struct {
	Damage   map[uint]domain.Damage       `json:"damage"`
	Recovery map[uint]domain.UnitRecovery `json:"recovery"`
}

type GameUserActionResult struct {
	Action GameAction          `json:"action"`
	Result domain.ActionResult `json:"result"`
}

type GameEvent struct {
	State            *GameState            `json:"state,omitempty"`
	Spot             *Spot                 `json:"spot,omitempty"`
	UserActionResult *GameUserActionResult `json:"userActionResult,omitempty"`
	EndRoundResult   *EndTurnResult        `json:"endRoundResult,omitempty"`
}
