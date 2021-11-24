package domain

import "fmt"

type Impact struct {
	Duration float32 `json:"duration,omitempty"`
	Chance   float32 `json:"chance,omitempty"`
}

type DamageImpact struct {
	Impact
	Damage
}

func (d DamageImpact) String() string {
	return fmt.Sprintf(
		"{%s, chance: %g, duration: %g}",
		d.Damage.String(),
		d.Impact.Chance,
		d.Impact.Duration,
	)
}

type UnitEnhancementImpact struct {
	Impact
	UnitEnhancement
}
