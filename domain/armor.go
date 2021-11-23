package domain

import "fmt"

type Armor struct {
	Equipment
}

func (a Armor) String() string {
	return fmt.Sprintf(
		"Armor: name: %s, description: %s, condition: %g, equipped: %t, requirements: {%v}, enhancement: %v",
		a.Name,
		a.Description,
		a.Condition,
		a.Equipped,
		a.Requirements,
		a.Enhancement,
	)
}
