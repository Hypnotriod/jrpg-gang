package domain

import (
	"fmt"
	"jrpg-gang/util"
)

type Armor struct {
	Equipment
}

func (a Armor) String() string {
	return fmt.Sprintf(
		"%s, description: %s, wearout: %g, durability: %g, equipped: %t, requirements: {%v}, modification: [%s]",
		a.Name,
		a.Description,
		a.Wearout,
		a.Durability,
		a.Equipped,
		a.Requirements,
		util.AsCommaSeparatedSlice(a.Modification),
	)
}
