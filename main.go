package main

import (
	"encoding/json"
	"fmt"
	"jrpg-gang/domain"
)

func main() {
	printArmor()
	printWeapon()
}

func printWeapon() {
	weapon := domain.Weapon{}
	weaponConfig := `{
		"id": "1111",
		"type": "weapon",
		"name": "Sword",
		"description": "Base one hand sword",
		"condition": 100,
		"hands": 1,
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
	}`
	json.Unmarshal([]byte(weaponConfig), &weapon)
	fmt.Println(weapon)
}

func printArmor() {
	armor := domain.Armor{}
	armorConfig := `{
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
	}`
	json.Unmarshal([]byte(armorConfig), &armor)
	fmt.Println(armor)
}
