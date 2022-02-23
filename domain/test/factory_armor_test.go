package test

import (
	"jrpg-gang/domain"
	"jrpg-gang/util"
	"testing"
)

func newBaseGloves(t *testing.T) *domain.Weapon {
	weapon, err := util.JsonToObject(
		&domain.Weapon{},
		`{
			"uid": 1000,
			"type": "armor",
			"name": "Gloves",
			"description": "Base gloves",
			"durability": 500,
			"requirements": {
				"strength": 10,
				"endurance": 30
			},
			"modification": [
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
	if err != nil {
		t.Fatal()
	}
	return weapon.(*domain.Weapon)
}
