package engine

import (
	"jrpg-gang/domain"
)

type CellConfiguration struct {
	FractionId uint `json:"fraction_id"`
}

type Cell struct {
	Unit   *domain.Unit       `json:"unit,omitempty"`
	Config *CellConfiguration `json:"config,omitempty"`
}
