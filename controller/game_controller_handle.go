package controller

import (
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
)

func (c *GameController) ConnectionStatusChanged(playerId engine.PlayerId, isOffline bool) {
	c.users.ConnectionStatusChanged(playerId, isOffline)
	if roomUid, ok := c.rooms.ConnectionStatusChanged(playerId, isOffline); ok {
		c.broadcastRoomStatus(roomUid)
	}
	if wrapper, ok := c.engines.Find(playerId); ok {
		defer wrapper.Unlock()
		wrapper.Lock()
		state, broadcastPlayerIds, ok := wrapper.ConnectionStatusChanged(playerId, isOffline)
		if ok {
			c.broadcastGameAction(broadcastPlayerIds, state)
		}
	}
}

func (c *GameController) Leave(playerId engine.PlayerId) {
	c.users.RemoveUser(playerId)
	if room, ok := c.rooms.PopByHostId(playerId); ok {
		c.broadcastRoomStatus(room.Uid)
	} else if roomUid, ok := c.rooms.RemoveUser(playerId); ok {
		c.broadcastRoomStatus(roomUid)
	}
	if wrapper, ok := c.engines.Unregister(playerId); ok {
		defer wrapper.Unlock()
		wrapper.Lock()
		state, broadcastPlayerIds, ok := wrapper.RemoveUser(playerId)
		if ok {
			c.broadcastGameAction(broadcastPlayerIds, state)
		}
	}
}

func (c *GameController) HandleRequest(playerId engine.PlayerId, requestRaw []byte) (engine.PlayerId, string) {
	response := NewResponse()
	request := parseRequest(requestRaw)
	if request == nil {
		return engine.PlayerIdEmpty, response.WithStatus(ResponseStatusMalformed)
	}
	response.Type = request.Type
	response.Id = request.Id
	if request.Type == RequestJoin {
		if playerId != engine.PlayerIdEmpty {
			return engine.PlayerIdEmpty, response.WithStatus(ResponseStatusNotAllowed)
		}
		return c.handleJoinRequest(request, response)
	}
	return engine.PlayerIdEmpty, c.serveRequest(playerId, request, response)
}

func (c *GameController) serveRequest(playerId engine.PlayerId, request *Request, response *Response) string {
	status := c.users.GetUserStatus(playerId)
	if status == users.UserStatusNotFound {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	switch request.Type {
	case RequestGameAction:
		return c.handleGameActionRequest(playerId, request, response)
	case RequestNextGamePhase:
		return c.handleGameNextPhaseRequest(playerId, request, response)
	case RequestGameState:
		return c.handleGameStateRequest(playerId, request, response)
	case RequestPlayerInfo:
		return c.handlePlayerInfoRequest(playerId, request, response)
	case RequestLeaveGame:
		return c.handleGameLeaveRequest(playerId, request, response)
	}
	if status.Test(users.UserStatusInGame) {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	switch request.Type {
	case RequestEnterLobby:
		return c.handleEnterLobbyRequest(playerId, request, response)
	case RequestExitLobby:
		return c.handleExitLobbyRequest(playerId, request, response)
	case RequestShopStatus:
		return c.handleShopStatusRequest(playerId, request, response)
	case RequestCreateGameRoom:
		return c.handleCreateGameRoomRequest(playerId, request, response)
	case RequestDestroyGameRoom:
		return c.handleDestroyGameRoomRequest(playerId, request, response)
	case RequestLobbyStatus:
		return c.handleLobbyStatusRequest(playerId, request, response)
	case RequestUserStatus:
		return c.handleUserStatusRequest(playerId, request, response)
	case RequestJoinGameRoom:
		return c.handleJoinGameRoomRequest(playerId, request, response)
	case RequestLeaveGameRoom:
		return c.handleLeaveGameRoomRequest(playerId, request, response)
	case RequestStartGame:
		return c.handleStartGameRequest(playerId, request, response)
	}
	if status.Test(users.UserStatusInRoom) {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	switch request.Type {
	case RequestConfiguratorAction:
		return c.handleConfiguratorActionRequest(playerId, request, response)
	case RequestShopAction:
		return c.handleShopActionRequest(playerId, request, response)
	}

	return response.WithStatus(ResponseStatusUnsupported)
}
