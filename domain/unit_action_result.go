package domain

import (
	"fmt"
	"jrpg-gang/util"
)

type ActionResultType string

const (
	Accomplished    ActionResultType = "accomplished"
	NotAccomplished ActionResultType = "notAccomplished"
	NotFound        ActionResultType = "notFound"
	NotEquipped     ActionResultType = "notEquipped"
	CantUse         ActionResultType = "cantUse"
	NoAmmunition    ActionResultType = "noAmmunition"
	NotCompatible   ActionResultType = "notCompatible"
	ZeroQuantity    ActionResultType = "zeroQuantity"
	NotEnoughSlots  ActionResultType = "notEnoughSlots"
	IsBroken        ActionResultType = "isBroken"
)

type ActionResult struct {
	InstantDamage        []Damage                 `json:"instantDamage,omitempty"`
	TemporalDamage       []DamageImpact           `json:"temporalDamage,omitempty"`
	InstantRecovery      []UnitState              `json:"instantRecovery,omitempty"`
	TemporalModification []UnitModificationImpact `json:"temporalModification,omitempty"`
	ResultType           ActionResultType         `json:"resultType"`
}

func (r ActionResult) String() string {
	return fmt.Sprintf("instant damage: [%s], temporal damage: [%s], instant recovery: [%s], temporal modification: [%s], result: %s",
		util.AsCommaSeparatedSlice(r.InstantDamage),
		util.AsCommaSeparatedSlice(r.TemporalDamage),
		util.AsCommaSeparatedSlice(r.InstantRecovery),
		util.AsCommaSeparatedSlice(r.TemporalModification),
		r.ResultType)
}

func (r *ActionResult) WithResultType(resultType ActionResultType) *ActionResult {
	r.ResultType = resultType
	return r
}
