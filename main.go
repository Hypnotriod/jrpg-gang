package main

import (
	"encoding/json"
	"fmt"
	"jrpg-gang/domain"
)

func main() {
	printUnit()
	printArmor()
	printWeapon()
	accumulateResistance()
}

func printUnit() {
	unit := jsonToObject(
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

func accumulateResistance() {
	unit := domain.Unit{}
	equipment := jsonToObject(
		&[]domain.Weapon{},
		`[{
			"name": "The thing",
			"equipped": true,
			"enhancement": [
				{
					"cutting": 30,
					"crushing": 5,
					"stabbing": 0
				},
				{
					"cold": 5,
					"cutting": -2
				}
			]
		}]`).(*[]domain.Weapon)
	for _, v := range *equipment {
		unit.Items = append(unit.Items, v)
	}
	fmt.Printf("Total resistance: {%v}\n", unit.TotalResistance())
}

func printWeapon() {
	weapon := jsonToObject(
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

func printArmor() {
	armor := jsonToObject(
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

func jsonToObject(object interface{}, config string) interface{} {
	err := json.Unmarshal([]byte(config), object)
	if err != nil {
		fmt.Printf("Can't parse %T: %v", object, err)
	}
	return object
}
