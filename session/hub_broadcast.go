package session

import "jrpg-gang/engine"

type broadcast struct {
	playerId engine.PlayerId
	message  []byte
}

func (h *Hub) broadcastGameMessageRoutine(broadcastPool <-chan broadcast) {
	for br := range broadcastPool {
		if client := h.getClient(br.playerId); client != nil {
			client.WriteMessage(br.message)
		}
	}
}

func (h *Hub) BroadcastGameMessageAsync(playerIds []engine.PlayerId, message []byte) {
	for _, playerId := range playerIds {
		h.broadcastPool <- broadcast{playerId, message}
	}
}

func (h *Hub) BroadcastGameMessageSync(playerIds []engine.PlayerId, message []byte) {
	for _, playerId := range playerIds {
		if client := h.getClient(playerId); client != nil {
			client.WriteMessage(message)
		}
	}
}
