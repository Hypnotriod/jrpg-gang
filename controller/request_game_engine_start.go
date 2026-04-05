package controller

import (
	"jrpg-gang/controller/chat"
	"jrpg-gang/controller/gameengines"
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
)

func (c *GameController) handleStartGameRequest(playerId engine.PlayerId, request *Request, response *Response) []byte {
	room, ok := c.rooms.PopByHostId(playerId)
	if !ok {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	scenario := c.scenarioConfig.Get(room.ScenarioId)
	if scenario == nil {
		return response.WithStatus(ResponseStatusNotFound)
	}
	actors := room.GetActors()
	ch := chat.NewChat(chat.ChatConfig{
		MaxMessages:         CHAT_MAX_MESSAGES,
		MaxMessageLength:    CHAT_MAX_MESSAGE_LENGTH,
		MessageRate:         CHAT_MESSAGE_RATE,
		MessageRateDuration: CHAT_MESSAGE_RATE_DURATION,
	}, c.broadcastGameChatMessage, nil)
	for _, actor := range actors {
		ch.AddParticipant(actor.GetPlayerId(), &chat.ChatParticipant{
			Nickname: actor.PlayerInfo.Nickname,
		})
	}
	wrapper := gameengines.NewGameEngineWrapper(engine.NewGameEngine(scenario, actors), c.broadcastGameAction, ch)
	wrapper.Lock()
	defer wrapper.Unlock()
	state, playerIds := c.engines.Register(wrapper)
	for _, playerId := range playerIds {
		c.users.ChangeUserStatus(playerId, users.UserStatusInGame)
		c.lobbyChat.RemoveParticipant(playerId)
	}
	c.broadcastGameAction(playerIds, state)
	c.broadcastRoomStatus(room.Uid)
	c.broadcastUsersStatus(playerIds)
	return response.WithStatus(ResponseStatusOk)
}
