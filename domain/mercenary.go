package domain

type Mercenary struct {
	Unit
	Price        *UnitBooty        `json:"price,omitempty"`
	Requirements *UnitRequirements `json:"requirements,omitempty"`
}
