package domain

import (
	"fmt"
	"jrpg-gang/util"
)

type Ammunition struct {
	Item
	Equipped bool           `json:"equipped"`
	Kind     string         `json:"kind"`
	Quantity uint           `json:"quantity"`
	Damage   []DamageImpact `json:"damage,omitempty"`
}

func (a Ammunition) String() string {
	return fmt.Sprintf(
		"%s, description: %s, kind: %s, quantity: %d, equipped: %t, damage: [%s], uid: %d",
		a.Name,
		a.Description,
		a.Kind,
		a.Quantity,
		a.Equipped,
		util.AsCommaSeparatedObjectsSlice(a.Damage),
		a.Uid,
	)
}

func (a *Ammunition) EnchanceDamageImpact(damage []DamageImpact) []DamageImpact {
	result := []DamageImpact{}
	instantDamage := DamageImpact{}
	temporalDamage := []DamageImpact{}
	for _, imp := range a.Damage {
		if imp.Duration == 0 {
			instantDamage.Accumulate(imp.Damage)
			instantDamage.Chance += imp.Chance
		} else {
			temporalDamage = append(temporalDamage, imp)
		}
	}
	for _, imp := range damage {
		if imp.Duration == 0 {
			imp.Damage.Accumulate(instantDamage.Damage)
			imp.Chance += instantDamage.Chance
		}
		result = append(result, imp)
	}
	return append(result, temporalDamage...)
}
