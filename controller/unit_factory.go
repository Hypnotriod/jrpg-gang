package controller

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

func newGameUnitRogue() *engine.GameUnit {
	// todo: prepare unit configuration
	unit, _ := util.JsonToObject(
		&engine.GameUnit{},
		`{
			"name": "Rogue",
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

func newGameUnitMage() *engine.GameUnit {
	// todo: prepare unit configuration
	unit, _ := util.JsonToObject(
		&engine.GameUnit{},
		`{
			"name": "Mage",
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
