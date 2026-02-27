package chat

import (
	"jrpg-gang/engine"
	"sync"
	"time"
	"unicode/utf8"
)

type BroadcastChatMessageFunc func(playerIds []engine.PlayerId, message *ChatMessage)

type ChatConfig struct {
	MaxMessages      uint `json:"maxMessages"`
	MaxMessageLength uint `json:"maxMessageLength"`
}

type ChatMessage struct {
	From      engine.PlayerId `json:"from"`
	To        engine.PlayerId `json:"to,omitempty"`
	Message   string          `json:"message"`
	Timestamp time.Time       `json:"timestamp"`
}

type ChatParticipant struct {
	Nickname string `json:"nickname"`
}

type ChatState struct {
	Participants map[engine.PlayerId]ChatParticipant `json:"participants"`
	Messages     []*ChatMessage                      `json:"messages"`
}

type Chat struct {
	participants         map[engine.PlayerId]ChatParticipant
	messages             []*ChatMessage
	config               ChatConfig
	broadcastChatMessage BroadcastChatMessageFunc
	mu                   sync.RWMutex
}

func NewChat(config ChatConfig, broadcastChatMessage BroadcastChatMessageFunc) *Chat {
	c := &Chat{
		participants:         make(map[engine.PlayerId]ChatParticipant),
		messages:             []*ChatMessage{},
		config:               config,
		broadcastChatMessage: broadcastChatMessage,
	}
	return c
}

func (c *Chat) AddParticipant(playerId engine.PlayerId, participant ChatParticipant) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.participants[playerId] = participant
}

func (c *Chat) RemoveParticipant(playerId engine.PlayerId) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.participants, playerId)
}

func (c *Chat) SendMessage(from engine.PlayerId, message string) *ChatMessage {
	if utf8.RuneCountInString(message) > int(c.config.MaxMessageLength) {
		return nil
	}
	c.mu.Lock()
	msg := &ChatMessage{
		From:      from,
		Message:   string(message),
		Timestamp: time.Now(),
	}
	c.messages = append(c.messages, msg)
	if len(c.messages) > int(c.config.MaxMessages) {
		c.messages = c.messages[1:]
	}
	to := []engine.PlayerId{}
	for playerId := range c.participants {
		if playerId != from {
			to = append(to, playerId)
		}
	}
	c.mu.Unlock()
	c.broadcastChatMessage(to, msg)
	return msg
}

func (c *Chat) State() *ChatState {
	c.mu.Lock()
	defer c.mu.Unlock()
	r := &ChatState{
		Participants: map[engine.PlayerId]ChatParticipant{},
		Messages:     c.messages[:],
	}
	for id, p := range c.participants {
		r.Participants[id] = p
	}
	return r
}

func (c *Chat) Dispose() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.messages = nil
	c.participants = nil
	c.broadcastChatMessage = nil
}
