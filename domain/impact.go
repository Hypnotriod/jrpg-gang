package domain

type ImpactType string

const (
	Permanent ImpactType = "permanent"
	Temporal  ImpactType = "temporal"
	Immediate ImpactType = "immediate"
)

type Impact struct {
	Duration float32    `json:"duration,omitempty"`
	Chance   float32    `json:"chance,omitempty"`
	Type     ImpactType `json:"type,omitempty"`
}

type DamageImpact struct {
	Impact
	Damage
}

type UnitEnhancementImpact struct {
	Impact
	UnitEnhancement
}
