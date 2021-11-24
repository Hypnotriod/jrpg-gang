package domain

import "fmt"

type ImpactType string

const (
	ImpactTypePermanent ImpactType = "permanent"
	ImpactTypeTemporal  ImpactType = "temporal"
	ImpactTypeImmediate ImpactType = "immediate"
)

type Impact struct {
	Duration float32    `json:"duration,omitempty"`
	Chance   float32    `json:"chance,omitempty"`
	Type     ImpactType `json:"type"`
}

type DamageImpact struct {
	Impact
	Damage
}

func (d DamageImpact) String() string {
	return fmt.Sprintf(
		"{damage: {%s, duration: %g, chance: %g}}",
		d.Damage.String(),
		d.Impact.Duration,
		d.Impact.Chance,
	)
}

type UnitEnhancementImpact struct {
	Impact
	UnitEnhancement
}
