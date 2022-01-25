package engine

import "jrpg-gang/domain"

type EndTurnResult struct {
	Damage   map[uint]domain.Damage       `json:"damage"`
	Recovery map[uint]domain.UnitRecovery `json:"recovery"`
}
