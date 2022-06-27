package controller

import (
	"jrpg-gang/domain"
	"jrpg-gang/util"
)

func NewTestShopItems() *domain.UnitInventory {
	// todo: prepare items configuration
	items, _ := util.JsonToObject(
		&domain.UnitInventory{},
		`{
			"weapon": [
			],
			"magic": [
			],
			"armor": [
			],
			"disposable": [
				{
					"name": "Health Sm",
					"type": "disposable",
					"code": "health-potion-01",
					"description": "Small health potion",
					"modification": [
						{
							"recovery": {
								"health": 35
							}
						},
						{
							"chance": 50,
							"duration": 5,
							"recovery": {
								"health": 2
							}
						}
					],
					"price": {
						"coins": 250
					}
				},
				{
					"name": "Poison",
					"type": "disposable",
					"code": "poison-01",
					"description": "Regular Poison",
					"modification": [
						{
							"duration": 3,
							"damage": {
								"poison": 10
							}
						}
					],
					"price": {
						"coins": 10
					}
				}
			],
			"ammunition": [
				{
					"name": "Arrow",
					"type": "ammunition",
					"code": "arrow-01",
					"description": "Base arrow",
					"kind": "arrow",
					"damage": [
						{
							"stabbing": 10
						},
						{
							"duration": 3,
							"chance": 15,
							"bleeding": 3
						}
					],
					"price": {
						"coins": 100
					}
				}
			]
		}`)
	return items.(*domain.UnitInventory)
}
