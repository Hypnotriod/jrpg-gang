package factory

import (
	"jrpg-gang/engine"
	"jrpg-gang/util"
)

func NewGameUnitByClass(class engine.GameUnitClass) *engine.GameUnit {
	// todo: test purpose only
	switch class {
	case engine.UnitClassTank:
		return newGameUnitTank()
	case engine.UnitClassMage:
		return newGameUnitMage()
	case engine.UnitClassRogue:
		return newGameUnitRogue()
	}
	return nil
}

func newGameUnitTank() *engine.GameUnit {
	// todo: prepare unit configuration
	unit, _ := util.JsonToObject(
		&engine.GameUnit{},
		`{
			"name": "Tank",
			"faction": 0,
			"booty": {
				"coins": 5000
			},
			"state": {
				"health": 100,
				"stamina": 100
			},
			"stats": {
				"progress": {
					"level": 1,
					"experience": 0
				},
				"baseAttributes": {
					"health": 100,
					"stamina": 100
				},
				"attributes": {
					"strength": 5,
					"physique": 5,
					"agility": 5,
					"endurance": 5,
					"intelligence": 5,
					"luck": 5
				},
				"resistance": {
					"stabbing": 5,
					"cutting": 5,
					"crushing": 5,
					"fire": 5,
					"cold": 5,
					"lighting": 5,
					"poison": 5,
					"exhaustion": 5,
					"manaDrain": 5,
					"fear": 5,
					"curse": 5
				}
			},
			"inventory": {
				"weapon": [
					{
						"type": "weapon",
						"name": "Base Sword",
						"code": "sword-01",
						"description": "Base one hand sword",
						"durability": 700,
						"slot": "weapon",
						"slotsNumber": 1,
						"equipped": true,
						"requirements": {
							"strength": 5
						},
						"range": {
							"maximumX": 1,
							"maximumY": 1
						},
						"damage": [
							{
								"chance": 50,
								"cutting": 30,
								"crushing": 5
							},
							{
								"duration": 3,
								"chance": 15,
								"bleeding": 3
							}
						],
						"modification": [
							{
								"damage": {
									"cutting": 30
								}
							}	
						],
						"useCost": {
							"stamina": 30
						}
					}
				]
			},
			"slots": {
				"head": 1,
				"neck": 1,
				"body": 1,
				"hand": 2,
				"leg": 2,
				"weapon": 2
			},
			"damage": [],
			"modification": []
		}`)
	return unit.(*engine.GameUnit)
}

func newGameUnitRogue() *engine.GameUnit {
	// todo: prepare unit configuration
	unit, _ := util.JsonToObject(
		&engine.GameUnit{},
		`{
			"name": "Rogue",
			"faction": 0,
			"booty": {
				"coins": 5000
			},
			"state": {
				"health": 100,
				"stamina": 100
			},
			"stats": {
				"progress": {
					"level": 1,
					"experience": 0
				},
				"baseAttributes": {
					"health": 100,
					"stamina": 100
				},
				"attributes": {
					"strength": 5,
					"physique": 5,
					"agility": 5,
					"endurance": 5,
					"intelligence": 5,
					"luck": 5
				},
				"resistance": {
					"stabbing": 5,
					"cutting": 5,
					"crushing": 5,
					"fire": 5,
					"cold": 5,
					"lighting": 5,
					"poison": 5,
					"exhaustion": 5,
					"manaDrain": 5,
					"fear": 5,
					"curse": 5
				}
			},
			"inventory": {
				"weapon": [
					{
						"type": "weapon",
						"name": "Bow",
						"code": "bow-01",
						"description": "Base two hand bow",
						"ammunitionKind": "arrow",
						"durability": 700,
						"slot": "weapon",
						"slotsNumber": 2,
						"equipped": true,
						"requirements": {
							"strength": 5
						},
						"range": {
							"maximumX": 10,
							"maximumY": 10,
							"minimumX": 2
						},
						"damage": [
						],
						"useCost": {
							"stamina": 25
						}
					},
					{
						"type": "weapon",
						"name": "Dagger",
						"code": "dagger-01",
						"description": "Sneaky one hand dagger",
						"durability": 700,
						"slot": "weapon",
						"slotsNumber": 1,
						"equipped": false,
						"requirements": {
							"strength": 5
						},
						"range": {
							"maximumX": 1,
							"maximumY": 1
						},
						"damage": [
							{
								"chance": 40,
								"cutting": 15
							},
							{
								"duration": 3,
								"chance": 10,
								"bleeding": 10
							}
						],
						"useCost": {
							"stamina": 15
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
						"quantity": 50,
						"equipped": true,
						"damage": [
							{
								"stabbing": 10,
								"chance": 50
							},
							{
								"duration": 3,
								"chance": 15,
								"bleeding": 3
							}
						]
					}
				]
			},
			"slots": {
				"head": 1,
				"neck": 1,
				"body": 1,
				"hand": 2,
				"leg": 2,
				"weapon": 2
			},
			"damage": [],
			"modification": []
		}`)
	return unit.(*engine.GameUnit)
}

func newGameUnitMage() *engine.GameUnit {
	// todo: prepare unit configuration
	unit, _ := util.JsonToObject(
		&engine.GameUnit{},
		`{
			"name": "Mage",
			"faction": 0,
			"booty": {
				"coins": 5000
			},
			"state": {
				"health": 100,
				"stamina": 100
			},
			"stats": {
				"progress": {
					"level": 1,
					"experience": 0
				},
				"baseAttributes": {
					"health": 100,
					"stamina": 100
				},
				"attributes": {
					"strength": 5,
					"physique": 5,
					"agility": 5,
					"endurance": 5,
					"intelligence": 5,
					"luck": 5
				},
				"resistance": {
					"stabbing": 5,
					"cutting": 5,
					"crushing": 5,
					"fire": 5,
					"cold": 5,
					"lighting": 5,
					"poison": 5,
					"exhaustion": 5,
					"manaDrain": 5,
					"fear": 5,
					"curse": 5
				}
			},
			"inventory": {
				"weapon": [
					{
						"type": "weapon",
						"name": "Staff",
						"code": "sorcerer-staff-01",
						"description": "Base one hand Sorcerer Staff",
						"ammunitionKind": "magicCharge",
						"durability": 700,
						"slot": "weapon",
						"slotsNumber": 1,
						"equipped": true,
						"requirements": {
							"intelligence": 5
						},
						"damage": [
						],
						"range": {
							"maximumX": 3,
							"maximumY": 3
						}
					}
				],
				"ammunition": [
					{
						"name": "Fire",
						"type": "ammunition",
						"code": "fire-charge-01",
						"description": "Sorcerer staff Fire Charge",
						"kind": "magicCharge",
						"quantity": 25,
						"equipped": true,
						"damage": [
							{
								"fire": 15,
								"chance": 50
							},
							{
								"duration": 3,
								"chance": 15,
								"fire": 5
							}
						]
					},
					{
						"name": "Cold",
						"type": "ammunition",
						"code": "cold-charge-01",
						"description": "Sorcerer staff Cold Charge",
						"kind": "magicCharge",
						"quantity": 25,
						"damage": [
							{
								"cold": 15,
								"chance": 50
							},
							{
								"duration": 3,
								"chance": 15,
								"cold": 5
							}
						]
					},
					{
						"name": "Lighting",
						"type": "ammunition",
						"code": "lighting-charge-01",
						"description": "Sorcerer staff Lighting Charge",
						"kind": "magicCharge",
						"quantity": 25,
						"damage": [
							{
								"lighting": 15,
								"chance": 50
							},
							{
								"duration": 3,
								"chance": 15,
								"lighting": 5
							}
						]
					}
				]
			},
			"slots": {
				"head": 1,
				"neck": 1,
				"body": 1,
				"hand": 2,
				"leg": 2,
				"weapon": 2
			},
			"damage": [],
			"modification": []
		}`)
	return unit.(*engine.GameUnit)
}
