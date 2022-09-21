package factory

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
					"code": "winter-cave-01",
					"battlefield": {
						"matrix": [
							[
								{"type": "space", "code": "tile-grass-01", "factions": [0]},
								{"type": "space", "code": "tile-grass-01", "factions": [0]},
								{"type": "space", "code": "tile-grass-01", "factions": [0]},
								{"type": "space", "code": "tile-grass-01", "factions": [0]}
							],
							[
								{"type": "space", "code": "tile-grass-01", "factions": [0]},
								{"type": "obstacle", "code": "tile-rock-01", "factions": []},
								{"type": "space", "code": "tile-grass-01", "factions": [0]},
								{"type": "space", "code": "tile-grass-01", "factions": [0]}
							],
							[
								{"type": "space", "code": "tile-grass-01", "factions": [1]},
								{"type": "space", "code": "tile-grass-01", "factions": [1]},
								{"type": "obstacle", "code": "tile-rock-01", "factions": []},
								{"type": "space", "code": "tile-grass-01", "factions": [1]}
							],
							[
								{"type": "space", "code": "tile-grass-01", "factions": [1]},
								{"type": "space", "code": "tile-grass-01", "factions": [1]},
								{"type": "space", "code": "tile-grass-01", "factions": [1]},
								{"type": "space", "code": "tile-grass-01", "factions": [1]}
							]
						],
						"units": [
							{
								"name": "North Troll",
								"code": "troll-01",
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
										"initiative": 0,
										"luck": 5
									},
									"resistance": {
										"stabbing": 0,
										"cutting": 0,
										"crushing": 20,
										"fire": 0,
										"cold": 50,
										"lighting": 0,
										"poison": 0,
										"exhaustion": 5,
										"manaDrain": 0,
										"fear": 0,
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
							},
							{
								"name": "Nasty Goblin",
								"code": "goblin-01",
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
										"strength": 10,
										"physique": 5,
										"agility": 15,
										"endurance": 30,
										"intelligence": 0,
										"initiative": 15,
										"luck": 5
									},
									"resistance": {
										"stabbing": 5,
										"cutting": 5,
										"crushing": 0,
										"fire": 0,
										"cold": 0,
										"lighting": 0,
										"poison": 15,
										"exhaustion": 0,
										"manaDrain": 0,
										"fear": 0,
										"curse": 0
									}
								},
								"slots": {
									"weapon": 1
								},
								"inventory": {
									"armor": [
									],
									"weapon": [
										{
											"type": "weapon",
											"name": "Mace",
											"description": "Rusty Mace",
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
													"cutting": 5,
													"crushing": 30
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
									"y": 0
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
