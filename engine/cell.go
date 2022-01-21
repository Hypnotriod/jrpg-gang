package engine

import (
	"jrpg-gang/domain"
)

type CellConfiguration struct {
	Position       Position `json:"position"`
	UnitFractionId uint     `json:"unitFractionId"`
}

type Cell struct {
	CellConfiguration
	Unit *domain.Unit `json:"unit,omitempty"`
}
