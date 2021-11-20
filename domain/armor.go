package domain

type Armor struct {
	Equipment
	Condition   float32           `json:"condition,omitempty"`
	Enhancement []UnitEnhancement `json:"enhancement,omitempty"`
}
