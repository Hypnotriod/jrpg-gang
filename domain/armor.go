package domain

import "fmt"

type Armor struct {
	Equipment
	Condition   float32           `json:"condition"`
	Enhancement []UnitEnhancement `json:"enhancement"`
}

func (a Armor) String() string {
	return fmt.Sprintf(
		"Armor: name: %s, condition: %g, equipped: %t, enhancement: %v",
		a.Name,
		a.Condition,
		a.Equipped,
		a.Enhancement)
}
