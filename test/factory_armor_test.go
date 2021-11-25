package test

import (
	"jrpg-gang/domain"
	"jrpg-gang/util"
	"testing"
)

func getBaseGloves(t *testing.T) *domain.Weapon {
	weapon, ok := util.JsonToObject(
		&domain.Weapon{},
		`{
			"id": "2222",
			"type": "armor",
			"name": "Gloves",
			"description": "Base gloves",
			"condition": 100,
			"requirements": {
				"strength": 10,
				"endurance": 30
			},
			"enhancement": [
				{
					"resistance": {
						"cutting": 10,
						"crushing": 2,
						"stabbing": 5,
						"curse": 5
					}
				}
			]
		}`)
	if !ok {
		t.Fatal()
	}
	return weapon.(*domain.Weapon)
}
