package domain

import (
	"fmt"
	"jrpg-gang/util"
)

type Disposable struct {
	Item
	Quantity     uint                     `json:"quantity,omitempty"`
	Range        ActionRange              `json:"range"`
	Damage       []DamageImpact           `json:"damage,omitempty"`
	Modification []UnitModificationImpact `json:"modification,omitempty"`
}

func (d Disposable) String() string {
	return fmt.Sprintf(
		"%s, description: %s, quantity: %d, range: {%v}, modification: [%s], damage: [%s], uid: %d",
		d.Name,
		d.Description,
		d.Quantity,
		d.Range,
		util.AsCommaSeparatedObjectsSlice(d.Modification),
		util.AsCommaSeparatedObjectsSlice(d.Damage),
		d.Uid,
	)
}
