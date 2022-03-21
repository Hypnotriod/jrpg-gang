package domain

import (
	"fmt"
	"jrpg-gang/util"
)

type ActionResultType string

const (
	ResultAccomplished    ActionResultType = "accomplished"
	ResultNotAccomplished ActionResultType = "notAccomplished"
	ResultNotAllowed      ActionResultType = "notAllowed"
	ResultNotFound        ActionResultType = "notFound"
	ResultNotEmpty        ActionResultType = "notEmpty"
	ResultNotEquipped     ActionResultType = "notEquipped"
	ResultNotReachable    ActionResultType = "notReachable"
	ResultCantUse         ActionResultType = "cantUse"
	ResultOutOfBounds     ActionResultType = "outOfBounds"
	ResultNoAmmunition    ActionResultType = "noAmmunition"
	ResultNotCompatible   ActionResultType = "notCompatible"
	ResultZeroQuantity    ActionResultType = "zeroQuantity"
	ResultNotEnoughSlots  ActionResultType = "notEnoughSlots"
	ResultIsBroken        ActionResultType = "isBroken"
)

type ActionResult struct {
	InstantDamage        []Damage                 `json:"instantDamage,omitempty"`
	TemporalDamage       []DamageImpact           `json:"temporalDamage,omitempty"`
	InstantRecovery      []UnitRecovery           `json:"instantRecovery,omitempty"`
	TemporalModification []UnitModificationImpact `json:"temporalModification,omitempty"`
	Position             []Position               `json:"position,omitempty"`
	ResultType           ActionResultType         `json:"resultType"`
}

func (r ActionResult) String() string {
	return fmt.Sprintf("instant damage: [%s], temporal damage: [%s], instant recovery: [%s], temporal modification: [%s], position: [%s], result: %s",
		util.AsCommaSeparatedSlice(r.InstantDamage),
		util.AsCommaSeparatedSlice(r.TemporalDamage),
		util.AsCommaSeparatedSlice(r.InstantRecovery),
		util.AsCommaSeparatedSlice(r.TemporalModification),
		util.AsCommaSeparatedSlice(r.Position),
		r.ResultType)
}

func NewActionResult() *ActionResult {
	action := &ActionResult{}
	return action
}

func (r *ActionResult) WithResultType(resultType ActionResultType) *ActionResult {
	r.ResultType = resultType
	return r
}
