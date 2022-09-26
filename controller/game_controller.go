package controller

import (
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
	users        *users.Users
	rooms        *rooms.GameRooms
	engines      *gameengines.GameEngines
	shop         *shop.Shop
	configurator *engine.UnitConfigurator
	broadcaster  GameControllerBroadcaster
}

func NewGameController() *GameController {
	c := &GameController{}
	c.users = users.NewUsers()
	c.rooms = rooms.NewGameRooms()
	c.engines = gameengines.NewGameEngines()
	c.shop = shop.NewShop()
	c.configurator = engine.NewUnitConfigurator()
	c.broadcaster = c
	return c
}

func (c *GameController) ConnectionStatusChanged(userId engine.UserId, isOffline bool) {
	c.users.ConnectionStatusChanged(userId, isOffline)
	if roomUid, ok := c.rooms.ConnectionStatusChanged(userId, isOffline); ok {
		c.broadcastRoomStatus(roomUid)
	}
	if state, broadcastUserIds, ok := c.engines.ConnectionStatusChanged(userId, isOffline); ok {
		c.broadcastGameState(broadcastUserIds, state)
	}
}

func (c *GameController) Leave(userId engine.UserId) {
	c.users.RemoveUser(userId)
	if room, ok := c.rooms.PopByHostId(userId); ok {
		c.broadcastRoomStatus(room.Uid)
	} else if roomUid, ok := c.rooms.RemoveUser(userId); ok {
		c.broadcastRoomStatus(roomUid)
	}
	if state, broadcastUserIds, ok := c.engines.RemoveUser(userId); ok {
		c.broadcastGameState(broadcastUserIds, state)
	}
}

func (c *GameController) HandleRequest(userId engine.UserId, requestRaw string) (engine.UserId, string) {
	response := NewResponse()
	request := parseRequest(&Request{}, requestRaw)
	if request == nil {
		return engine.UserIdEmpty, response.WithStatus(ResponseStatusMailformed)
	}
	response.Type = request.Type
	response.Id = request.Id
	if request.Type == RequestJoin {
		if userId != engine.UserIdEmpty {
			return engine.UserIdEmpty, response.WithStatus(ResponseStatusNotAllowed)
		}
		return c.handleJoinRequest(requestRaw, response)
	}
	return engine.UserIdEmpty, c.serveRequest(userId, request, requestRaw, response)
}

func (c *GameController) serveRequest(userId engine.UserId, request *Request, requestRaw string, response *Response) string {
	status := c.users.GetUserStatus(userId)
	if status == users.UserStatusNotFound {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	switch request.Type {
	case RequestGameAction:
		return c.handleGameActionRequest(userId, requestRaw, response)
	case RequestNextGamePhase:
		return c.handleGameNextPhaseRequest(userId, requestRaw, response)
	case RequestGameState:
		return c.handleGameStateRequest(userId, requestRaw, response)
	case RequestPlayerInfo:
		return c.handlePlayerInfoRequest(userId, requestRaw, response)
	case RequestLeaveGame:
		return c.handleGameLeaveRequest(userId, requestRaw, response)
	}
	if status.Test(users.UserStatusInGame) {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	switch request.Type {
	case RequestShopStatus:
		return c.handleShopStatusRequest(userId, requestRaw, response)
	case RequestCreateGameRoom:
		return c.handleCreateGameRoomRequest(userId, requestRaw, response)
	case RequestDestroyGameRoom:
		return c.handleDestroyGameRoomRequest(userId, requestRaw, response)
	case RequestLobbyStatus:
		return c.handleLobbyStatusRequest(userId, requestRaw, response)
	case RequestUserStatus:
		return c.handleUserStatusRequest(userId, requestRaw, response)
	case RequestJoinGameRoom:
		return c.handleJoinGameRoomRequest(userId, requestRaw, response)
	case RequestLeaveGameRoom:
		return c.handleLeaveGameRoomRequest(userId, requestRaw, response)
	case RequestStartGame:
		return c.handleStartGameRequest(userId, requestRaw, response)
	}
	if status.Test(users.UserStatusInRoom) {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	switch request.Type {
	case RequestConfiguratorAction:
		return c.handleConfiguratorActionRequest(userId, requestRaw, response)
	case RequestShopAction:
		return c.handleShopActionRequest(userId, requestRaw, response)
	}

	return response.WithStatus(ResponseStatusUnsupported)
}
