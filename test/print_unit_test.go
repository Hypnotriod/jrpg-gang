package test

import (
	"fmt"
	"jrpg-gang/domain"
	"jrpg-gang/util"
	"testing"
)

func TestPrintUnit(t *testing.T) {
	unit, ok := util.JsonToObject(
		&domain.Unit{},
		`{
			"name": "North Troll",
			"state": {
				"health": 100,
				"stamina": 100,
				"mana": 100
			},
			"stats": {
				"progress": {
					"level": 10,
					"experience": 45555
				},
				"baseAttributes": {
					"health": 100,
					"stamina": 100,
					"mana": 100
				},
				"attributes": {
					"strength": 15,
					"physique": 5,
					"dexterity": 5,
					"endurance": 5,
					"intelligence": 25,
					"luck": 20
				},
				"resistance": {
					"stabbing": 10,
					"cutting": 10,
					"crushing": 20,
					"fire": 5,
					"cold": 50,
					"lighting": 15,
					"poison": 10,
					"stunning": 20,
					"exhaustion": 5,
					"fear": 5,
					"curse": 0
				}
			}
		}`)
	if !ok {
		t.Fatal()
	}
	fmt.Println(unit)
}
