package controller

import "jrpg-gang/engine"

type GameController struct {
	users   *Users
	rooms   *GameRooms
	engines *GameEngines
}

func NewController() *GameController {
	c := &GameController{}
	c.users = NewUsers()
	c.rooms = NewGameRooms()
	c.engines = NewGameEngines()
	return c
}

func (c *GameController) HandleRequest(requestRaw string, userId engine.UserId) string {
	response := NewResponse()
	request := parseRequest(&Request{}, requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMailformed)
	}
	response.Type = request.Type
	response.Id = request.Id
	return c.serveRequest(userId, request, requestRaw, response)
}

func (c *GameController) serveRequest(userId engine.UserId, request *Request, requestRaw string, response *Response) string {
	if request.Type == RequestJoin {
		return c.handleJoinRequest(requestRaw, response)
	}
	if !c.users.Has(userId) {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	switch request.Type {
	case RequestGameAction:
		return c.handleGameActionRequest(userId, requestRaw, response)
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
