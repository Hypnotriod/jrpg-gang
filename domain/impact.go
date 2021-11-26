package domain

import "fmt"

type Impact struct {
	Duration int     `json:"duration,omitempty"`
	Chance   float32 `json:"chance"`
}

type DamageImpact struct {
	Impact
	Damage
}

func (d DamageImpact) String() string {
	if d.Chance != 0 {
		return fmt.Sprintf("{%s, chance: %g, duration: %d}", d.Damage.String(), d.Impact.Chance, d.Impact.Duration)
	} else {
		return fmt.Sprintf("{%s, duration: %d}", d.Damage.String(), d.Impact.Duration)
	}
}

type UnitEnhancementImpact struct {
	Impact
	UnitEnhancement
}
