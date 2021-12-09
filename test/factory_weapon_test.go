package test

import (
	"jrpg-gang/domain"
	"jrpg-gang/util"
	"testing"
)

func newBaseSword(t *testing.T) *domain.Weapon {
	weapon, ok := util.JsonToObject(
		&domain.Weapon{},
		`{
			"uid": 4000,
			"type": "weapon",
			"name": "Sword",
			"description": "Base one hand sword",
			"durability": 700,
			"slot": "weapon",
			"slotsNumber": 1,
			"requirements": {
				"strength": 5
			},
			"damage": [
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
					"damage": {
						"cutting": 30
					}
				}	
			]
		}`)
	if !ok {
		t.Fatal()
	}
	return weapon.(*domain.Weapon)
}
