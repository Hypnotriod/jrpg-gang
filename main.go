package main

import (
	"encoding/json"
	"fmt"
	"jrpg-gang/domain"
)

func main() {
	printArmor()
}

func printArmor() {
	armor := domain.Armor{}
	armorJson := `{
		"id": "1234",
		"type": "armor",
		"name": "Gloves",
		"condition": 100,
		"enhancement": [
			{
				"cutting": 10,
				"crushing": 2,
				"stabbing": 5,
				"curse": 5
			}
		]
	}`
	json.Unmarshal([]byte(armorJson), &armor)
	fmt.Println(armor)
}
