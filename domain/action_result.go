package domain

import "jrpg-gang/util"

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
	InstantDamage        map[uint][]Damage                 `json:"instantDamage,omitempty"`
	TemporalDamage       map[uint][]DamageImpact           `json:"temporalDamage,omitempty"`
	InstantRecovery      map[uint][]UnitRecovery           `json:"instantRecovery,omitempty"`
	TemporalModification map[uint][]UnitModificationImpact `json:"temporalModification,omitempty"`
	Experience           map[uint]uint                     `json:"experience,omitempty"`
	Drop                 map[uint]UnitBooty                `json:"drop,omitempty"`
	Booty                *UnitBooty                        `json:"booty,omitempty"`
	Result               ActionResultType                  `json:"result"`
	WearoutIncreased     bool                              `json:"-"`
	UseCostReduced       bool                              `json:"-"`
}

func (r *ActionResult) WithStun(unitUid uint) bool {
	return util.Any(r.InstantDamage[unitUid], func(damage Damage) bool {
		return damage.WithStun
	})
}

func NewActionResult() *ActionResult {
	action := &ActionResult{}
	action.InstantDamage = map[uint][]Damage{}
	action.TemporalDamage = map[uint][]DamageImpact{}
	action.InstantRecovery = map[uint][]UnitRecovery{}
	action.TemporalModification = map[uint][]UnitModificationImpact{}
	action.Experience = map[uint]uint{}
	action.Drop = map[uint]UnitBooty{}
	return action
}

func (r *ActionResult) WithResult(result ActionResultType) *ActionResult {
	r.Result = result
	return r
}
