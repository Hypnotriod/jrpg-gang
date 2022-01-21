package engine

import "jrpg-gang/domain"

type GameUnit struct {
	domain.Unit
	FractionId uint `json:"fractionId"`
}
