package main

import (
	"encoding/json"
	"fmt"
	"jrpg-gang/domain"
)

func main() {
	printArmor()
	printWeapon()
	accumulateResistance()
}

func accumulateResistance() {
	unit := domain.Unit{}
	equipment := jsonToItem(
		&domain.Weapon{},
		`{
			"name": "The thing",
			"equipped": true,
			"enhancement": [
				{
					"cutting": 30,
					"crushing": 5,
					"stabbing": 0
				},
				{
					"cold": 5
				}
			]
		}`).(*domain.Weapon)
	unit.Items = append(unit.Items, *equipment)
	resistance := unit.TotalResistance()
	fmt.Print(resistance)
}

func printWeapon() {
	weapon := jsonToItem(
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
			"damage": [
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
	armor := jsonToItem(
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

func jsonToItem(object interface{}, config string) interface{} {
	err := json.Unmarshal([]byte(config), object)
	if err != nil {
		fmt.Printf("Can't parse %T: %v", object, err)
	}
	return object
}
