package domain

type Provision struct {
	Item
	Quantity uint         `json:"quantity,omitzero"`
	Recovery UnitRecovery `json:"recovery"`
}
