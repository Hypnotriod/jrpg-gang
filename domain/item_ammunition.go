package domain

import "fmt"

type Ammunition struct {
	Item
	Selected bool           `json:"selected"`
	Kind     string         `json:"kind"`
	Quantity uint           `json:"quantity"`
	Damage   []DamageImpact `json:"damage"`
}

func (a Ammunition) String() string {
	return fmt.Sprintf(
		"%s, description: %s, kind: %s, quantity: %d, selected: %t, damage: %v",
		a.Name,
		a.Description,
		a.Kind,
		a.Quantity,
		a.Selected,
		a.Damage,
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
