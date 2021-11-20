package domain

type ImpactType string

const (
	Permanent ImpactType = "permanent"
	Temporal  ImpactType = "temporal"
	Immediate ImpactType = "immediate"
)

type Impact struct {
	Duration float32    `json:"duration,omitempty"`
	Type     ImpactType `json:"type"`
}

type DamageImpact struct {
	Impact
	Damage
}

type UnitEnhancementImpact struct {
	Impact
	UnitEnhancement
}
