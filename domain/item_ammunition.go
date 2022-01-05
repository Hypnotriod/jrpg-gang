package domain

import "fmt"

type Ammunition struct {
	Item
	Selected bool   `json:"selected"`
	Kind     string `json:"kind"`
	Quantity uint   `json:"quantity"`
	Damage   Damage `json:"damage,omitempty"`
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
	for _, imp := range damage {
		imp.Damage.Accumulate(a.Damage)
		result = append(result, imp)
	}
	return result
}
