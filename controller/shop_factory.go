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
			],
			"ammunition": [
				{
					"code": "abc1",
					"name": "Arrow",
					"type": "ammunition",
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
