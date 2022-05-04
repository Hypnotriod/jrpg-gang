package controller

import (
	"jrpg-gang/engine"
	"jrpg-gang/util"
)

func NewTestScenario() *engine.GameScenario {
	unit, _ := util.JsonToObject(
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
	return unit.(*engine.GameScenario)
}
