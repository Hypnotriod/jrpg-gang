package controller

const USER_NICKNAME_REGEX string = `^[a-zA-Z0-9][a-zA-Z0-9-_ ]{2,18}[a-zA-Z0-9]$`
const GAME_ROOM_MAX_CAPACITY uint = 4
const CHAT_MAX_MESSAGES uint = 50
const CHAT_MAX_MESSAGE_LENGTH uint = 128

const ITEMS_CONFIG_PATH string = "./private/items_config.json"
const UNITS_CONFIG_PATH string = "./private/units_config.json"
const SHOP_CONFIG_PATH string = "./private/shop_config.json"
const SCENARIO_CONFIG_PATH string = "./private/scenarios_config.json"
const JOBS_CONFIG_PATH string = "./private/jobs_config.json"
