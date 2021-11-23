package test

import (
	"fmt"
	"jrpg-gang/domain"
	"jrpg-gang/util"
	"testing"
)

func TestPrintArmor(t *testing.T) {
	armor := util.JsonToObject(
		&domain.Armor{},
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
					"cutting": 10,
					"crushing": 2,
					"stabbing": 5,
					"curse": 5
				}
			]
		}`)
	fmt.Println(armor)
}

func TestPrintWeapon(t *testing.T) {
	weapon := util.JsonToObject(
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
					"cutting": 30,
					"crushing": 5,
					"stabbing": 0
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
	fmt.Println(weapon)
}

func TestPrintUnit(t *testing.T) {
	unit := util.JsonToObject(
		&domain.Unit{},
		`{
			"state": {
				"health": 100,
				"stamina": 100,
				"mana": 100
			}
		}`).(*domain.Unit)
	fmt.Println(*unit)
}
