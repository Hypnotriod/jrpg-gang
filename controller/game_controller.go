package controller

import (
	"jrpg-gang/controller/chat"
	"jrpg-gang/controller/config"
	"jrpg-gang/controller/employment"
	"jrpg-gang/controller/gameengines"
	"jrpg-gang/controller/quests"
	"jrpg-gang/controller/rooms"
	"jrpg-gang/controller/shop"
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
	"jrpg-gang/persistance"

	log "github.com/sirupsen/logrus"
)

type GameControllerBroadcaster interface {
	BroadcastGameMessageSync(playerIds []engine.PlayerId, message []byte)
	BroadcastGameMessageAsync(playerIds []engine.PlayerId, message []byte)
}

type GameController struct {
	users          *users.Users
	rooms          *rooms.GameRooms
	engines        *gameengines.GameEngines
	shop           *shop.Shop
	quests         *quests.Quests
	employment     *employment.Employment
	configurator   *engine.UnitConfigurator
	itemsConfig    *config.GameItemsConfig
	unitsConfig    *config.GameUnitsConfig
	scenarioConfig *config.GameScenariosConfig
	broadcaster    GameControllerBroadcaster
	persistance    *persistance.Persistance
	lobbyChat      *chat.Chat
}

func NewGameController(persistance *persistance.Persistance) *GameController {
	c := &GameController{}
	c.users = users.NewUsers()
	c.rooms = rooms.NewGameRooms()
	c.engines = gameengines.NewGameEngines()
	c.shop = shop.NewShop()
	c.quests = quests.NewQuests()
	c.employment = employment.NewEmployment()
	c.itemsConfig = config.NewGameItemsConfig()
	c.unitsConfig = config.NewGameUnitsConfig()
	c.scenarioConfig = config.NewGameScenariosConfig()
	c.configurator = engine.NewUnitConfigurator()
	c.broadcaster = c
	c.persistance = persistance
	c.lobbyChat = chat.NewChat(chat.ChatConfig{
		MaxMessages:         CHAT_MAX_MESSAGES,
		MaxMessageLength:    CHAT_MAX_MESSAGE_LENGTH,
		MessageRate:         CHAT_MESSAGE_RATE,
		MessageRateDuration: CHAT_MESSAGE_RATE_DURATION,
	}, c.broadcastLobbyChatMessage, c.broadcastLobbyChatParticipant)
	c.init()
	return c
}

func (c *GameController) init() {
	if err := c.itemsConfig.Load(ITEMS_CONFIG_PATH); err != nil {
		log.Fatal("Unable to load items configuration: ", err)
	}
	if err := c.shop.LoadItems(SHOP_CONFIG_PATH, c.itemsConfig); err != nil {
		log.Fatal("Unable to load shop configuration: ", err)
	}
	if err := c.quests.LoadItems(QUESTS_CONFIG_PATH, c.itemsConfig); err != nil {
		log.Fatal("Unable to load quests configuration: ", err)
	}
	if err := c.unitsConfig.LoadUnits(UNITS_CONFIG_PATH, c.itemsConfig); err != nil {
		log.Fatal("Unable to load units configuration: ", err)
	}
	if err := c.scenarioConfig.LoadScenarios(SCENARIO_CONFIG_PATH, c.unitsConfig); err != nil {
		log.Fatal("Unable to load scenarios configuration: ", err)
	}
	if err := c.employment.Load(JOBS_CONFIG_PATH); err != nil {
		log.Fatal("Unable to load jobs configuration: ", err)
	}
}
