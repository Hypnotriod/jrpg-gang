package test

import (
	"jrpg-gang/engine"
	"jrpg-gang/util"
	"testing"
)

func newBasicScenario(t *testing.T) *engine.GameScenario {
	unit, err := util.JsonToObject(
		&engine.GameScenario{},
		`{
			"spots": [
				{
					"name": "Winter Cave",
					"battlefield": {
						"matrix": [
							[
								{"type": "space", "factions": [0]},
								{"type": "space", "factions": [0]},
								{"type": "space", "factions": [1]},
								{"type": "space", "factions": [1]}
							],
							[
								{"type": "space", "factions": [0]},
								{"type": "obstacle", "factions": []},
								{"type": "space", "factions": [1]},
								{"type": "space", "factions": [1]}
							],
							[
								{"type": "space", "factions": [0]},
								{"type": "space", "factions": [0]},
								{"type": "obstacle", "factions": []},
								{"type": "space", "factions": [1]}
							],
							[
								{"type": "space", "factions": [0]},
								{"type": "space", "factions": [0]},
								{"type": "space", "factions": [1]},
								{"type": "space", "factions": [1]}
							]
						],
						"units": [
							{
								"name": "North Troll",
								"faction": 1,
								"state": {
									"health": 100,
									"stamina": 100
								},
								"stats": {
									"progress": {
										"level": 10,
										"experience": 45550
									},
									"baseAttributes": {
										"health": 100,
										"stamina": 100
									},
									"attributes": {
										"strength": 15,
										"physique": 20,
										"agility": 5,
										"endurance": 30,
										"intelligence": 0,
										"luck": 5
									},
									"resistance": {
										"stabbing": 10,
										"cutting": 10,
										"crushing": 20,
										"fire": 5,
										"cold": 50,
										"lighting": 15,
										"poison": 10,
										"exhaustion": 20,
										"manaDrain": 5,
										"fear": 5,
										"curse": 0
									}
								},
								"slots": {
									"weapon": 1
								},
								"inventory": {
									"armor": [
										{
											"type": "armor",
											"name": "Rusty Helmet",
											"description": "Rusty helmet",
											"wearout": 35,
											"durability": 40,
											"slot": "head",
											"slotsNumber": 1,
											"equipped": true,
											"requirements": {
												"strength": 5
											},
											"modification": [
												{
													"resistance": {
														"exhaustion": 20
													}
												}	
											]
										}
									],
									"weapon": [
										{
											"type": "weapon",
											"name": "Axe",
											"description": "Rusty Axe",
											"durability": 700,
											"slot": "weapon",
											"slotsNumber": 1,
											"requirements": {
												"strength": 5
											},
											"range": {
												"maximumX": 1,
												"maximumY": 1
											},
											"useCost": {
												"stamina": 10
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
											]
										}
									]
								},
								"position": {
									"x": 2,
									"y": 1
								}
							}
						]
					}
				}
			],
			"path": [
				{"0": 1}
			]
		}`)
	if err != nil {
		t.Fatal()
	}
	return unit.(*engine.GameScenario)
}

func newGameUnitTank(t *testing.T) *engine.GameUnit {
	unit, _ := util.JsonToObject(
		&engine.GameUnit{},
		`{
			"name": "Tank",
			"faction": 0,
			"userId": "userId1",
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
			"modification": [
				{
					"recovery": {
						"stamina": 20
					}
				}
			],
			"inventory": {
				"weapon": [
					{
						"type": "weapon",
						"name": "Axe",
						"description": "Rusty Axe",
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
						"useCost": {
							"stamina": 10
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
			}
		}`)
	return unit.(*engine.GameUnit)
}
