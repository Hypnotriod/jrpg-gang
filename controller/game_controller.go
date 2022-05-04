package controller

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

func (c *GameController) HandleRequest(requestRaw string) string {
	response := NewResponse()
	request := parseRequest(requestRaw)
	if request != nil {
		response.Type = request.Type
		response.Id = request.Id
		return c.serveRequest(request, requestRaw, response)
	}
	return response.WithStatus(ResponseStatusMailformed)
}

func (c *GameController) serveRequest(request *Request, requestRaw string, response *Response) string {
	if request.Type == RequestJoin {
		return c.handleJoinRequest(requestRaw, response)
	}
	if !c.users.Has(request.UserId) {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	switch request.Type {
	case RequestCreateGameRoom:
		return c.handleCreateGameRoomRequest(requestRaw, response)
	case RequestDestroyGameRoom:
		return c.handleDestroyGameRoomRequest(requestRaw, response)
	case RequestLobbyStatus:
		return c.handleLobbyStatusRequest(requestRaw, response)
	case RequestJoinGameRoom:
		return c.handleJoinGameRoomRequest(requestRaw, response)
	case RequestLeaveGameRoom:
		return c.handleLeaveGameRoomRequest(requestRaw, response)
	case RequestStartGame:
		return c.handleStartGameRequest(requestRaw, response)
	}
	return response.WithStatus(ResponseStatusUnsupported)
}
