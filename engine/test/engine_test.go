package test

import (
	"fmt"
	"jrpg-gang/domain"
	"jrpg-gang/engine"
	"testing"
)

func TestCreateGameEngine(t *testing.T) {
	scenario := newBasicScenario(t)
	actors := []*engine.GameUnit{newGameUnitTank(t)}
	eng := engine.NewGameEngine(scenario, actors)
	fmt.Printf("engine: {%v}\n", eng)
	fmt.Println()
	gameLoop(t, eng)
}

func gameLoop(t *testing.T, eng *engine.GameEngine) {
	event := eng.NewGameEvent()
	actionResult := eng.ExecuteAction(placeUnitAction(0, 0), "userId1")
	fmt.Printf("action result: {%v}\n", actionResult)
	for eng.NextPhaseRequired() {
		eng.NextPhase(event)
		fmt.Println()
		fmt.Printf("event phase result: {%v}\n", event)
	}
}

func placeUnitAction(x, y int) engine.GameAction {
	return engine.GameAction{
		Uid:      1,
		Action:   engine.GameAtionPlace,
		Position: domain.Position{X: x, Y: y},
	}
}
