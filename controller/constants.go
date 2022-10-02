package controller

const USER_NICKNAME_REGEX string = `^[a-zA-Z0-9][a-zA-Z0-9-_ ]{2,18}[a-zA-Z0-9]$`
const GAME_ROOM_MAX_CAPACITY uint = 4

const ITEMS_CONFIG_PATH string = "./private/items_config.json"
const UNITS_CONFIG_PATH string = "./private/units_config.json"
const SHOP_CONFIG_PATH string = "./private/shop_config.json"
const SCENARIO_CONFIG_PATH string = "./private/scenario_config.json"
