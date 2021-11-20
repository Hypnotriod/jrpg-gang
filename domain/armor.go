package domain

type Armor struct {
	Equipment
	Condition   float32           `json:"condition"`
	Enhancement []UnitEnhancement `json:"enhancement"`
}
