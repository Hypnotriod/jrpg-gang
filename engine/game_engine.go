package engine

import (
	"fmt"
)

type GameEngine struct {
	Battlefield *Battlefield `json:"battlefield"`
}

func (e GameEngine) String() string {
	return fmt.Sprintf(
		"battlefield: {%v}",
		e.Battlefield,
	)
}

func NewGameEngine(battlefield *Battlefield) *GameEngine {
	engine := &GameEngine{}
	engine.Battlefield = battlefield
	return engine
}
