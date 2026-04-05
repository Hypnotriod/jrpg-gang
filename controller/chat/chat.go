package chat

import (
	"errors"
	"jrpg-gang/engine"
	"sync"
	"time"
	"unicode/utf8"
)

var (
	ErrMessageIsTooLong    error = errors.New("chat message is too long")
	ErrParticipantNotFound error = errors.New("chat participant not found")
	ErrMessagerateLimit    error = errors.New("chat message rate limit reached")
)

type BroadcastChatMessageFunc func(playerIds []engine.PlayerId, message *ChatMessage)
type BroadcastChatParticipantFunc func(playerIds []engine.PlayerId, playerId engine.PlayerId, participant *ChatParticipant)

type ChatConfig struct {
	MaxMessages         uint          `json:"maxMessages"`
	MaxMessageLength    uint          `json:"maxMessageLength"`
	MessageRate         uint          `json:"messagerate"`
	MessageRateDuration time.Duration `json:"messagerateDuration"`
}

type ChatMessage struct {
	From      engine.PlayerId `json:"from"`
	To        engine.PlayerId `json:"to,omitempty"`
	Message   string          `json:"message"`
	Timestamp time.Time       `json:"timestamp"`
}

type ChatParticipant struct {
	Nickname             string `json:"nickname"`
	Unavailable          bool   `json:"unavailable,omitempty"`
	lastMessageTimestamp time.Time
	messageRate          float64
}

type ChatState struct {
	Participants map[engine.PlayerId]ChatParticipant `json:"participants"`
	Messages     []*ChatMessage                      `json:"messages"`
}

type Chat struct {
	participants             map[engine.PlayerId]*ChatParticipant
	messages                 []*ChatMessage
	config                   ChatConfig
	broadcastChatMessage     BroadcastChatMessageFunc
	broadcastChatParticipant BroadcastChatParticipantFunc
	mu                       sync.RWMutex
}

func NewChat(config ChatConfig, broadcastChatMessage BroadcastChatMessageFunc, broadcastChatParticipant BroadcastChatParticipantFunc) *Chat {
	c := &Chat{
		participants:             map[engine.PlayerId]*ChatParticipant{},
		messages:                 []*ChatMessage{},
		config:                   config,
		broadcastChatMessage:     broadcastChatMessage,
		broadcastChatParticipant: broadcastChatParticipant,
	}
	return c
}

func (c *Chat) AddParticipant(playerId engine.PlayerId, participant *ChatParticipant) {
	c.mu.Lock()
	defer c.mu.Unlock()
	participant.lastMessageTimestamp = time.Now()
	c.participants[playerId] = participant
	if c.broadcastChatParticipant != nil {
		to := make([]engine.PlayerId, 0, len(c.participants)-1)
		for participantId := range c.participants {
			if playerId != participantId {
				to = append(to, participantId)
			}
		}
		c.broadcastChatParticipant(to, playerId, participant)
	}
}

func (c *Chat) RemoveParticipant(playerId engine.PlayerId) {
	c.mu.Lock()
	defer c.mu.Unlock()
	participant, ok := c.participants[playerId]
	if !ok {
		return
	}
	participant.Unavailable = true
	if c.broadcastChatParticipant != nil {
		to := make([]engine.PlayerId, 0, len(c.participants)-1)
		for participantId := range c.participants {
			if playerId != participantId {
				to = append(to, participantId)
			}
		}
		c.broadcastChatParticipant(to, playerId, participant)
	}
}

func (c *Chat) SendMessage(from engine.PlayerId, message string) (*ChatMessage, error) {
	if utf8.RuneCountInString(message) > int(c.config.MaxMessageLength) {
		return nil, ErrMessageIsTooLong
	}
	c.mu.Lock()
	sender, ok := c.participants[from]
	if !ok || sender.Unavailable {
		c.mu.Unlock()
		return nil, ErrParticipantNotFound
	}
	if !c.manageMessageRate(sender) {
		c.mu.Unlock()
		return nil, ErrMessagerateLimit
	}
	msg := &ChatMessage{
		From:      from,
		Message:   string(message),
		Timestamp: time.Now(),
	}
	c.messages = append(c.messages, msg)
	if len(c.messages) > int(c.config.MaxMessages) {
		c.messages = c.messages[1:]
	}
	to := make([]engine.PlayerId, 0, len(c.participants)-1)
	for playerId := range c.participants {
		if playerId != from {
			to = append(to, playerId)
		}
	}
	c.mu.Unlock()
	c.broadcastChatMessage(to, msg)
	return msg, nil
}

func (c *Chat) State() *ChatState {
	c.mu.Lock()
	defer c.mu.Unlock()
	r := &ChatState{
		Participants: map[engine.PlayerId]ChatParticipant{},
		Messages:     c.messages[:],
	}
	for id, p := range c.participants {
		r.Participants[id] = *p
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

func (c *Chat) manageMessageRate(participant *ChatParticipant) bool {
	now := time.Now()
	elapsed := now.Sub(participant.lastMessageTimestamp)
	participant.messageRate += 1.0 - elapsed.Seconds()*(float64(c.config.MessageRate)/c.config.MessageRateDuration.Seconds())
	participant.messageRate = max(participant.messageRate, 1.0)
	participant.lastMessageTimestamp = now
	return participant.messageRate < float64(c.config.MessageRate)
}
