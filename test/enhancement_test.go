package test

import (
	"fmt"
	"jrpg-gang/domain"
	"jrpg-gang/util"
	"testing"
)

func TestAccumulateResistance(t *testing.T) {
	unit := newAgileHero(t)
	equipment, ok := util.JsonToObject(
		&[]domain.Weapon{},
		`[
			{
			  "name": "The thing",
			  "equipped": true,
			  "enhancement": [
				{
				  "damage": {
					"cutting": 30,
					"crushing": 5,
					"stabbing": 0
				  },
				  "resistance": {
					"cold": 5,
					"cutting": -2
				  }
				}
			  ]
			}
		]`)
	if !ok {
		t.Fatal()
	}
	unit.Inventory.Weapon = append(unit.Inventory.Weapon, *equipment.(*[]domain.Weapon)...)
	fmt.Printf("Total enhancement: {%v}\n", unit.TotalEnhancement())
}
