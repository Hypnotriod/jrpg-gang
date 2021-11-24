package test

import (
	"jrpg-gang/domain"
	"jrpg-gang/util"
	"testing"
)

func getBaseSword(t *testing.T) *domain.Weapon {
	weapon, ok := util.JsonToObject(
		&domain.Weapon{},
		`{
			"id": "1111",
			"type": "weapon",
			"name": "Sword",
			"description": "Base one hand sword",
			"condition": 100,
			"slot": "weapon",
			"slotsNumber": 1,
			"requirements": {
				"strength": 50
			},
			"impact": [
				{
					"chance": 75,
					"cutting": 30,
					"crushing": 5,
					"stabbing": 0
				},
				{
					"duration": 3,
					"chance": 35,
					"bleeding": 3
				}
			],
			"enhancement": [
				{
					"cutting": 30,
					"crushing": 5,
					"stabbing": 0
				}
			]
		}`)
	if !ok {
		t.Fatal()
	}
	return weapon.(*domain.Weapon)
}
