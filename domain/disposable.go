package domain

type Disposable struct {
	Item
	Impact      []DamageImpact          `json:"impact,omitempty"`
	Enhancement []UnitEnhancementImpact `json:"enhancement,omitempty"`
}
