package domain

type Spell struct {
	Item
	Requirements UnitAttributes          `json:"requirements"`
	Impact       []DamageImpact          `json:"impact,omitempty"`
	Enhancement  []UnitEnhancementImpact `json:"enhancement,omitempty"`
}
