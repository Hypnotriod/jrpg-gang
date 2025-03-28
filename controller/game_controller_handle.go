package controller

import (
	"jrpg-gang/auth"
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
	"jrpg-gang/persistance/model"
)

func (c *GameController) HandleUserAuthenticated(credentials auth.UserCredentials) auth.AuthenticationStatus {
	status := auth.AuthenticationStatus{}
	if user, ok := c.users.GetByEmail(model.UserEmail(credentials.Email)); ok {
		c.Leave(user.Id)
	}
	userModel := c.persistance.GetOrCreateUser(credentials)
	if userModel == nil {
		return status
	}
	status.Token = c.persistance.AddUserToAuthCache(userModel)
	status.IsNewPlayer = len(userModel.Units) == 0
	status.IsAuthenticated = true
	return status
}

func (c *GameController) ConnectionStatusChanged(playerId engine.PlayerId, isOffline bool) {
	c.users.ConnectionStatusChanged(playerId, isOffline)
	if roomUid, ok := c.rooms.ConnectionStatusChanged(playerId, isOffline); ok {
		c.broadcastRoomStatus(roomUid)
	}
	if wrapper, ok := c.engines.Find(playerId); ok {
		wrapper.Lock()
		defer wrapper.Unlock()
		state, broadcastPlayerIds, ok := wrapper.ConnectionStatusChanged(playerId, isOffline)
		if ok {
			c.broadcastGameAction(broadcastPlayerIds, state)
		}
	}
}

func (c *GameController) Leave(playerId engine.PlayerId) {
	if room, ok := c.rooms.PopByHostId(playerId); ok {
		c.broadcastRoomStatus(room.Uid)
	} else if roomUid, ok := c.rooms.RemoveUser(playerId); ok {
		c.broadcastRoomStatus(roomUid)
	}
	if wrapper, ok := c.engines.Unregister(playerId); ok {
		wrapper.Lock()
		result, broadcastPlayerIds, unit, ok := wrapper.LeaveGame(playerId)
		wrapper.Unlock()
		if ok {
			c.users.ResetUser(playerId)
			c.users.UpdateWithUnitOnGameComplete(playerId, &unit)
			if len(broadcastPlayerIds) > 0 {
				c.broadcastGameAction(broadcastPlayerIds, result)
			}
		}
	}
	user := c.users.RemoveUser(playerId)
	if user != nil {
		c.persistUser(user)
		c.employment.ClearStatus(user)
	}
}

func (c *GameController) HandleConfigurationRequest(requestRaw []byte) []byte {
	response := NewResponse()
	request := parseRequest(requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMalformed)
	}
	response.Type = request.Type
	response.Id = request.Id
	switch request.Type {
	case RequestSetPlayerInfo:
		return c.handleSetPlayerInfoRequest(request, response)
	}
	return response.WithStatus(ResponseStatusUnsupported)
}

func (c *GameController) HandleJoinRequest(data *JoinRequestData) (engine.PlayerId, []byte) {
	response := NewResponse()
	response.Type = RequestJoin
	return c.handleJoinRequest(data, response)
}

func (c *GameController) HandleRequest(playerId engine.PlayerId, requestRaw []byte) []byte {
	response := NewResponse()
	request := parseRequest(requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMalformed)
	}
	response.Type = request.Type
	response.Id = request.Id
	status := c.users.GetUserStatus(playerId)
	if status == users.UserStatusNotFound {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	switch request.Type {
	case RequestLeave:
		c.Leave(playerId)
		return response.WithStatus(ResponseStatusOk)
	case RequestJobsStatus:
		return c.handleJobStatusRequest(playerId, request, response)
	case RequestShopStatus:
		return c.handleShopStatusRequest(playerId, request, response)
	case RequestLobbyStatus:
		return c.handleLobbyStatusRequest(playerId, request, response)
	case RequestUserStatus:
		return c.handleUserStatusRequest(playerId, request, response)
	}
	if status.Test(users.UserStatusInGame) {
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
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	if status.Test(users.UserStatusAtJob) {
		switch request.Type {
		case RequestQuitJob:
			return c.handleQuitJobRequest(playerId, request, response)
		case RequestCompleteJob:
			return c.handleCompleteJobRequest(playerId, request, response)
		}
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	switch request.Type {
	case RequestEnterLobby:
		return c.handleEnterLobbyRequest(playerId, request, response)
	case RequestExitLobby:
		return c.handleExitLobbyRequest(playerId, request, response)
	case RequestCreateGameRoom:
		return c.handleCreateGameRoomRequest(playerId, request, response)
	case RequestDestroyGameRoom:
		return c.handleDestroyGameRoomRequest(playerId, request, response)
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
	case RequestApplyForAJob:
		return c.handleApplyForAJobRequest(playerId, request, response)
	case RequestConfiguratorAction:
		return c.handleConfiguratorActionRequest(playerId, request, response)
	case RequestShopAction:
		return c.handleShopActionRequest(playerId, request, response)
	case RequestSwitchUnit:
		return c.handleSwitchUnitRequest(playerId, request, response)
	}

	return response.WithStatus(ResponseStatusUnsupported)
}
