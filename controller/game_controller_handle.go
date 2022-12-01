package controller

import (
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
)

func (c *GameController) ConnectionStatusChanged(userId engine.UserId, isOffline bool) {
	c.users.ConnectionStatusChanged(userId, isOffline)
	if roomUid, ok := c.rooms.ConnectionStatusChanged(userId, isOffline); ok {
		c.broadcastRoomStatus(roomUid)
	}
	state, broadcastUserIds, unlock, ok := c.engines.ConnectionStatusChanged(userId, isOffline)
	if ok {
		c.broadcastGameAction(broadcastUserIds, state)
	}
	if unlock != nil {
		unlock()
	}
}

func (c *GameController) Leave(userId engine.UserId) {
	c.users.RemoveUser(userId)
	if room, ok := c.rooms.PopByHostId(userId); ok {
		c.broadcastRoomStatus(room.Uid)
	} else if roomUid, ok := c.rooms.RemoveUser(userId); ok {
		c.broadcastRoomStatus(roomUid)
	}
	state, broadcastUserIds, unlock, ok := c.engines.RemoveUser(userId)
	if ok {
		c.broadcastGameAction(broadcastUserIds, state)
	}
	if unlock != nil {
		unlock()
	}
}

func (c *GameController) HandleRequest(userId engine.UserId, requestRaw []byte) (engine.UserId, string) {
	response := NewResponse()
	request := parseRequest(requestRaw)
	if request == nil {
		return engine.UserIdEmpty, response.WithStatus(ResponseStatusMalformed)
	}
	response.Type = request.Type
	response.Id = request.Id
	if request.Type == RequestJoin {
		if userId != engine.UserIdEmpty {
			return engine.UserIdEmpty, response.WithStatus(ResponseStatusNotAllowed)
		}
		return c.handleJoinRequest(request, response)
	}
	return engine.UserIdEmpty, c.serveRequest(userId, request, response)
}

func (c *GameController) serveRequest(userId engine.UserId, request *Request, response *Response) string {
	status := c.users.GetUserStatus(userId)
	if status == users.UserStatusNotFound {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	switch request.Type {
	case RequestGameAction:
		return c.handleGameActionRequest(userId, request, response)
	case RequestNextGamePhase:
		return c.handleGameNextPhaseRequest(userId, request, response)
	case RequestGameState:
		return c.handleGameStateRequest(userId, request, response)
	case RequestPlayerInfo:
		return c.handlePlayerInfoRequest(userId, request, response)
	case RequestLeaveGame:
		return c.handleGameLeaveRequest(userId, request, response)
	}
	if status.Test(users.UserStatusInGame) {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	switch request.Type {
	case RequestEnterLobby:
		return c.handleEnterLobbyRequest(userId, request, response)
	case RequestExitLobby:
		return c.handleExitLobbyRequest(userId, request, response)
	case RequestShopStatus:
		return c.handleShopStatusRequest(userId, request, response)
	case RequestCreateGameRoom:
		return c.handleCreateGameRoomRequest(userId, request, response)
	case RequestDestroyGameRoom:
		return c.handleDestroyGameRoomRequest(userId, request, response)
	case RequestLobbyStatus:
		return c.handleLobbyStatusRequest(userId, request, response)
	case RequestUserStatus:
		return c.handleUserStatusRequest(userId, request, response)
	case RequestJoinGameRoom:
		return c.handleJoinGameRoomRequest(userId, request, response)
	case RequestLeaveGameRoom:
		return c.handleLeaveGameRoomRequest(userId, request, response)
	case RequestStartGame:
		return c.handleStartGameRequest(userId, request, response)
	}
	if status.Test(users.UserStatusInRoom) {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	switch request.Type {
	case RequestConfiguratorAction:
		return c.handleConfiguratorActionRequest(userId, request, response)
	case RequestShopAction:
		return c.handleShopActionRequest(userId, request, response)
	}

	return response.WithStatus(ResponseStatusUnsupported)
}
