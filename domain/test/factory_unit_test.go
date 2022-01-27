package test

import (
	"jrpg-gang/domain"
	"jrpg-gang/util"
	"testing"
)

func newAgileHero(t *testing.T) *domain.Unit {
	unit, ok := util.JsonToObject(
		&domain.Unit{},
		`{
			"name": "Patrick",
			"state": {
				"health": 50,
				"stamina": 50
			},
			"stats": {
				"progress": {
					"level": 2,
					"experience": 35000
				},
				"baseAttributes": {
					"health": 100,
					"stamina": 100
				},
				"attributes": {
					"strength": 5,
					"physique": 5,
					"agility": 35,
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
			"slots": {
				"head": 1,
				"neck": 1,
				"body": 1,
				"hand": 2,
				"leg": 2,
				"weapon": 2
			}
		}`)
	if !ok {
		t.Fatal()
	}
	return unit.(*domain.Unit)
}

func newAgileHeroWithWeapon(t *testing.T) *domain.Unit {
	unit, ok := util.JsonToObject(
		&domain.Unit{},
		`{
			"name": "Patrick",
			"state": {
				"health": 50,
				"stamina": 50
			},
			"stats": {
				"progress": {
					"level": 2,
					"experience": 35000
				},
				"baseAttributes": {
					"health": 100,
					"stamina": 100
				},
				"attributes": {
					"strength": 5,
					"physique": 5,
					"agility": 35,
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
			"slots": {
				"head": 1,
				"neck": 1,
				"body": 1,
				"hand": 2,
				"leg": 2,
				"weapon": 2
			},
			"inventory": {
				"weapon": [
					{
						"uid": 4000,
						"type": "weapon",
						"name": "Sword",
						"description": "Base one hand sword",
						"durability": 700,
						"slot": "weapon",
						"slotsNumber": 1,
						"requirements": {
							"strength": 5
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
					},
					{
						"uid": 4001,
						"type": "weapon",
						"name": "Bow",
						"description": "Base two hand bow",
						"ammunitionKind": "arrow",
						"durability": 700,
						"slot": "weapon",
						"slotsNumber": 2,
						"requirements": {
							"strength": 5
						},
						"damage": [
							{
								"chance": 40,
								"stabbing": 30
							}
						]
					}
				],
				"ammunition": [
					{
						"uid": 5000,
						"name": "Arrow",
						"type": "ammunition",
						"description": "Base arrow",
						"kind": "arrow",
						"quantity": 5,
						"damage": [
							{
								"stabbing": 10
							},
							{
								"duration": 3,
								"chance": 15,
								"bleeding": 3
							}
						]
					}
				]
			}
		}`)
	if !ok {
		t.Fatal()
	}
	return unit.(*domain.Unit)
}

func newMagicianHeroWithMagic(t *testing.T) *domain.Unit {
	unit, ok := util.JsonToObject(
		&domain.Unit{},
		`{
			"name": "Sedrick",
			"state": {
				"health": 50,
				"stamina": 30,
				"mana": 100
			},
			"stats": {
				"progress": {
					"level": 2,
					"experience": 35000
				},
				"baseAttributes": {
					"health": 100,
					"stamina": 30,
					"mana": 100
				},
				"attributes": {
					"strength": 5,
					"physique": 5,
					"agility": 5,
					"endurance": 5,
					"intelligence": 35,
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
			"slots": {
				"head": 1,
				"neck": 1,
				"body": 1,
				"hand": 2,
				"leg": 2,
				"weapon": 2
			},
			"inventory": {
				"magic": [
					{
						"uid": 3000,
						"type": "magic",
						"name": "Fire ball",
						"description": "Base fire ball",
						"requirements": {
							"intelligence": 15
						},
						"damage": [
							{
								"chance": 40,
								"fire": 30
							},
							{
								"chance": 50,
								"duration": 3,
								"fire": 5
							}
						],
						"useCost": {
							"mana": 30
						}
					},
					{
						"uid": 3001,
						"type": "magic",
						"name": "Small healing",
						"description": "Base small healing",
						"requirements": {
							"intelligence": 15
						},
						"modification": [
							{
								"chance": 50,
								"recovery": {
									"health": 30
								}
							},
							{
								"duration": 3,
								"recovery": {
									"health": 3
								}
							}
						],
						"useCost": {
							"mana": 5
						}
					}
				]
			}
		}`)
	if !ok {
		t.Fatal()
	}
	return unit.(*domain.Unit)
}

func newNorthTroll(t *testing.T) *domain.Unit {
	unit, ok := util.JsonToObject(
		&domain.Unit{},
		`{
			"name": "North Troll",
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
			"inventory": {
				"armor": [
					{
						"uid": 2000,
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
				]
			}
		}`)
	if !ok {
		t.Fatal()
	}
	return unit.(*domain.Unit)
}
