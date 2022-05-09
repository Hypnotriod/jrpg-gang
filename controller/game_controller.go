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

func NewController(broadcaster GameControllerBroadcaster) *GameController {
	c := &GameController{}
	c.users = NewUsers()
	c.rooms = NewGameRooms()
	c.engines = NewGameEngines()
	c.broadcaster = broadcaster
	return c
}

func (c *GameController) HandleRequest(requestRaw string) string {
	response := NewResponse()
	request := parseRequest(&Request{}, requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMailformed)
	}
	response.Type = request.Type
	response.Id = request.Id
	return c.serveRequest(request, requestRaw, response)
}

func (c *GameController) serveRequest(request *Request, requestRaw string, response *Response) string {
	if request.Type == RequestJoin {
		return c.handleJoinRequest(requestRaw, response)
	}
	if !c.users.Has(request.UserId) {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	switch request.Type {
	case RequestGameAction:
		return c.handleGameActionRequest(requestRaw, response)
	case RequestGameState:
		return c.handleGameStateRequest(requestRaw, response)
	}
	if c.engines.IsUserInGame(request.UserId) {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	switch request.Type {
	case RequestCreateGameRoom:
		return c.handleCreateGameRoomRequest(requestRaw, response)
	case RequestDestroyGameRoom:
		return c.handleDestroyGameRoomRequest(requestRaw, response)
	case RequestLobbyStatus:
		return c.handleLobbyStatusRequest(requestRaw, response)
	case RequestUserStatus:
		return c.handleUserStatusRequest(requestRaw, response)
	case RequestJoinGameRoom:
		return c.handleJoinGameRoomRequest(requestRaw, response)
	case RequestLeaveGameRoom:
		return c.handleLeaveGameRoomRequest(requestRaw, response)
	case RequestStartGame:
		return c.handleStartGameRequest(requestRaw, response)
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
