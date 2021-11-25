package test

import (
	"fmt"
	"jrpg-gang/domain"
	"jrpg-gang/util"
	"testing"
)

func TestAccumulateResistance(t *testing.T) {
	unit := getAgileHero(t)
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
	for _, v := range *equipment.(*[]domain.Weapon) {
		unit.Items = append(unit.Items, v)
	}
	fmt.Printf("Total enhancement: {%v}\n", unit.TotalEnhancement(true))
}
