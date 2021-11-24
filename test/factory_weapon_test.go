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
				"strength": 5
			},
			"impact": [
				{
					"chance": 40,
					"cutting": 30,
					"crushing": 5
				},
				{
					"duration": 3,
					"chance": 15,
					"bleeding": 3
				}
			],
			"enhancement": [
				{
					"cutting": 30,
					"crushing": 5
				}
			]
		}`)
	if !ok {
		t.Fatal()
	}
	return weapon.(*domain.Weapon)
}
