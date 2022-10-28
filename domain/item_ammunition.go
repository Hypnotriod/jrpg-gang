package domain

type Ammunition struct {
	Item
	Equipped bool           `json:"equipped,omitempty"`
	Kind     string         `json:"kind"`
	Quantity uint           `json:"quantity,omitempty"`
	Damage   []DamageImpact `json:"damage,omitempty"`
}

func (a *Ammunition) EnchanceDamageImpact(damage []DamageImpact) []DamageImpact {
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
			instantDamage.Damage.Accumulate(imp.Damage)
			instantDamage.Chance += imp.Chance
		} else {
			temporalDamage = append(temporalDamage, imp)
		}
	}
	result := []DamageImpact{instantDamage}
	return append(result, temporalDamage...)
}
