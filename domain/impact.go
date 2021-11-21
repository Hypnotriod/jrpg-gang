package domain

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

type UnitEnhancementImpact struct {
	Impact
	UnitEnhancement
}
