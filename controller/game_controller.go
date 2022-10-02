package controller

import (
	"jrpg-gang/controller/config"
	"jrpg-gang/controller/gameengines"
	"jrpg-gang/controller/rooms"
	"jrpg-gang/controller/shop"
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
)

type GameControllerBroadcaster interface {
	BroadcastGameMessageSync(userIds []engine.UserId, message string)
	BroadcastGameMessageAsync(userIds []engine.UserId, message string)
}

type GameController struct {
	users          *users.Users
	rooms          *rooms.GameRooms
	engines        *gameengines.GameEngines
	shop           *shop.Shop
	configurator   *engine.UnitConfigurator
	itemsConfig    *config.GameItemsConfig
	unitsConfig    *config.GameUnitsConfig
	scenarioConfig *config.GameScenariosConfig
	broadcaster    GameControllerBroadcaster
}

func NewGameController() *GameController {
	c := &GameController{}
	c.users = users.NewUsers()
	c.rooms = rooms.NewGameRooms()
	c.engines = gameengines.NewGameEngines()
	c.shop = shop.NewShop()
	c.itemsConfig = config.NewGameItemsConfig()
	c.unitsConfig = config.NewGameUnitsConfig()
	c.scenarioConfig = config.NewGameScenariosConfig()
	c.configurator = engine.NewUnitConfigurator()
	c.broadcaster = c
	c.prepare()
	return c
}

func (c *GameController) prepare() {
	if err := c.itemsConfig.Load(ITEMS_CONFIG_PATH); err != nil {
		panic(err)
	}
	if err := c.shop.LoadItems(UNITS_CONFIG_PATH, c.itemsConfig); err != nil {
		panic(err)
	}
	if err := c.unitsConfig.LoadUnits(UNITS_CONFIG_PATH, c.itemsConfig); err != nil {
		panic(err)
	}
	if err := c.scenarioConfig.LoadScenarios(SCENARIO_CONFIG_PATH, c.unitsConfig); err != nil {
		panic(err)
	}
}
