package test

import (
	"jrpg-gang/domain"
	"jrpg-gang/util"
	"testing"
)

func getWeakLuckyHero(t *testing.T) *domain.Unit {
	unit, ok := util.JsonToObject(
		&domain.Unit{},
		`{
			"name": "Patrick",
			"state": {
				"health": 50,
				"stamina": 50
			},
			"stats": {
				"progress": {
					"level": 2,
					"experience": 35000
				},
				"baseAttributes": {
					"health": 50,
					"stamina": 50
				},
				"attributes": {
					"strength": 5,
					"physique": 5,
					"dexterity": 5,
					"endurance": 5,
					"intelligence": 5,
					"luck": 100
				},
				"resistance": {
					"stabbing": 5,
					"cutting": 5,
					"crushing": 5,
					"fire": 5,
					"cold": 5,
					"lighting": 5,
					"poison": 5,
					"stunning": 5,
					"exhaustion": 5,
					"fear": 5,
					"curse": 5
				}
			}
		}`)
	if !ok {
		t.Fatal()
	}
	return unit.(*domain.Unit)
}

func getNorthTroll(t *testing.T) *domain.Unit {
	unit, ok := util.JsonToObject(
		&domain.Unit{},
		`{
			"name": "North Troll",
			"state": {
				"health": 100,
				"stamina": 100
			},
			"stats": {
				"progress": {
					"level": 10,
					"experience": 45550
				},
				"baseAttributes": {
					"health": 100,
					"stamina": 100
				},
				"attributes": {
					"strength": 15,
					"physique": 20,
					"dexterity": 5,
					"endurance": 30,
					"intelligence": 0,
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
	return unit.(*domain.Unit)
}
