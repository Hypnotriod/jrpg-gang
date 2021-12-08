package domain

import "fmt"

type Armor struct {
	Equipment
}

func (a Armor) String() string {
	return fmt.Sprintf(
		"%s, description: %s, condition: %g, strength: %g, equipped: %t, requirements: {%v}, enhancement: %v",
		a.Name,
		a.Description,
		a.Strength,
		a.Condition,
		a.Equipped,
		a.Requirements,
		a.Enhancement,
	)
}
