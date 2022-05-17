package controller

import "jrpg-gang/engine"

type GameControllerBroadcaster interface {
	BroadcastGameMessage(userIds []engine.UserId, message string)
}

type GameController struct {
	users       *Users
	rooms       *GameRooms
	engines     *GameEngines
	broadcaster GameControllerBroadcaster
}

func NewGameController() *GameController {
	c := &GameController{}
	c.users = NewUsers()
	c.rooms = NewGameRooms()
	c.engines = NewGameEngines()
	c.broadcaster = c
	return c
}

func (c *GameController) RegisterBroadcaster(broadcaster GameControllerBroadcaster) {
	c.broadcaster = broadcaster
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
		return c.handleJoinRequest(requestRaw, response)
	}
	return engine.UserIdEmpty, c.serveRequest(userId, request, requestRaw, response)
}

func (c *GameController) serveRequest(userId engine.UserId, request *Request, requestRaw string, response *Response) string {
	if !c.users.Has(userId) {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	switch request.Type {
	case RequestGameAction:
		return c.handleGameActionRequest(userId, requestRaw, response)
	case RequestNextGamePhase:
		return c.handleGameNextPhaseRequest(userId, requestRaw, response)
	case RequestGameState:
		return c.handleGameStateRequest(userId, requestRaw, response)
	}
	if c.engines.IsUserInGame(userId) {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	switch request.Type {
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
	return response.WithStatus(ResponseStatusUnsupported)
}

func (c *GameController) broadcastGameAction(userIds []engine.UserId, result *engine.GameEvent) {
	response := NewResponse()
	response.Type = RequestGameAction
	response.Data[DataKeyActionResult] = result
	c.broadcaster.BroadcastGameMessage(userIds, response.WithStatus(ResponseStatusOk))
}

func (c *GameController) broadcastGameState(userIds []engine.UserId, state *engine.GameEvent) {
	response := NewResponse()
	response.Type = RequestGameState
	response.Data[DataKeyGameState] = state
	c.broadcaster.BroadcastGameMessage(userIds, response.WithStatus(ResponseStatusOk))
}

func (c *GameController) BroadcastGameMessage(userIds []engine.UserId, message string) {
}
