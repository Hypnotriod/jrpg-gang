package domain

import (
	"jrpg-gang/util"
)

type ActionResultType string

const (
	ResultAccomplished       ActionResultType = "accomplished"
	ResultNotAccomplished    ActionResultType = "notAccomplished"
	ResultNotAllowed         ActionResultType = "notAllowed"
	ResultNotFound           ActionResultType = "notFound"
	ResultNotEmpty           ActionResultType = "notEmpty"
	ResultNotEquipped        ActionResultType = "notEquipped"
	ResultNotReachable       ActionResultType = "notReachable"
	ResultCantUse            ActionResultType = "cantUse"
	ResultOutOfBounds        ActionResultType = "outOfBounds"
	ResultNoAmmunition       ActionResultType = "noAmmunition"
	ResultNotCompatible      ActionResultType = "notCompatible"
	ResultZeroQuantity       ActionResultType = "zeroQuantity"
	ResultNotEnoughSlots     ActionResultType = "notEnoughSlots"
	ResultNotEnoughResources ActionResultType = "notEnoughResources"
	ResultIsBroken           ActionResultType = "isBroken"
)

type ActionResult struct {
	InstantDamage        []Damage                 `json:"instantDamage,omitempty"`
	TemporalDamage       []DamageImpact           `json:"temporalDamage,omitempty"`
	InstantRecovery      []UnitRecovery           `json:"instantRecovery,omitempty"`
	TemporalModification []UnitModificationImpact `json:"temporalModification,omitempty"`
	Experience           map[uint]uint            `json:"experience,omitempty"`
	Drop                 map[uint]UnitBooty       `json:"drop,omitempty"`
	Booty                *UnitBooty               `json:"booty,omitempty"`
	Result               ActionResultType         `json:"result"`
}

func (r *ActionResult) WithStun() bool {
	return util.Any(r.InstantDamage, func(damage Damage) bool {
		return damage.WithStun
	})
}

func (r *ActionResult) IsCriticalMiss() bool {
	return util.Any(r.InstantDamage, func(damage Damage) bool {
		return damage.IsCriticalMiss
	}) || util.Any(r.TemporalDamage, func(damage DamageImpact) bool {
		return damage.IsCriticalMiss
	})
}

func NewActionResult() *ActionResult {
	action := &ActionResult{}
	action.Experience = map[uint]uint{}
	action.Drop = map[uint]UnitBooty{}
	return action
}

func (r *ActionResult) WithResult(result ActionResultType) *ActionResult {
	r.Result = result
	return r
}
