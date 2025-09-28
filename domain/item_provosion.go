package domain

type Provision struct {
	Item
	Quantity uint         `json:"quantity,omitempty"`
	Recovery UnitRecovery `json:"recovery"`
}
