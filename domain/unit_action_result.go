package domain

import (
	"fmt"
	"jrpg-gang/util"
)

type ActionResultType uint

const (
	Accomplished    ActionResultType = iota
	NotAccomplished ActionResultType = iota
	NotFound        ActionResultType = iota
	IsNotEquipped   ActionResultType = iota
	CantUse         ActionResultType = iota
	HasNoAmmunition ActionResultType = iota
	IsNotCompatible ActionResultType = iota
	ZeroQuantity    ActionResultType = iota
)

func (r ActionResultType) String() string {
	var result string
	switch r {
	case Accomplished:
		result = "accomplished"
	case NotAccomplished:
		result = "not accomplished"
	case NotFound:
		result = "not found"
	case IsNotEquipped:
		result = "is not equipped"
	case CantUse:
		result = "can't use"
	case HasNoAmmunition:
		result = "has no ammunition"
	case IsNotCompatible:
		result = "is not compatible"
	case ZeroQuantity:
		result = "zero quantity"
	}
	return fmt.Sprintf("%s (%d)", result, r)
}

type UseInventoryItemActionResult struct {
	InstantDamage        []Damage                 `json:"instantDamage,omitempty"`
	TemporalDamage       []DamageImpact           `json:"temporalDamage,omitempty"`
	InstantRecovery      []UnitState              `json:"instantRecovery,omitempty"`
	TemporalModification []UnitModificationImpact `json:"temporalModification,omitempty"`
	ResultType           ActionResultType         `json:"resultType"`
}

func (r UseInventoryItemActionResult) String() string {
	return fmt.Sprintf("instant damage: [%s], temporal damage: [%s], instant recovery: [%s], temporal modification: [%s], result: %v",
		util.AsCommaSeparatedSlice(r.InstantDamage),
		util.AsCommaSeparatedSlice(r.TemporalDamage),
		util.AsCommaSeparatedSlice(r.InstantRecovery),
		util.AsCommaSeparatedSlice(r.TemporalModification),
		r.ResultType)
}

func (r *UseInventoryItemActionResult) WithResultType(resultType ActionResultType) *UseInventoryItemActionResult {
	r.ResultType = resultType
	return r
}
