package session

import "jrpg-gang/engine"

type broadcast struct {
	userId  engine.UserId
	message string
}

func (h *Hub) broadcastGameMessageRoutine(broadcastPool <-chan broadcast) {
	for br := range broadcastPool {
		if client := h.getClient(br.userId); client != nil {
			client.WriteMessage(br.message)
		}
	}
}

func (h *Hub) BroadcastGameMessageAsync(userIds []engine.UserId, message string) {
	for _, userId := range userIds {
		h.broadcastPool <- broadcast{userId, message}
	}
}

func (h *Hub) BroadcastGameMessageSync(userIds []engine.UserId, message string) {
	for _, userId := range userIds {
		if client := h.getClient(userId); client != nil {
			client.WriteMessage(message)
		}
	}
}
