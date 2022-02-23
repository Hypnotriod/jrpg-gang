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
	gameLoop(t, eng)
}

func gameLoop(t *testing.T, eng *engine.GameEngine) {
	action := engine.GameAction{
		Uid:      1,
		Action:   engine.GameAtionPlace,
		Position: domain.Position{X: 0, Y: 0},
	}
	fmt.Println()
	actionResult := eng.ExecuteAction(action, "abcd1234")
	fmt.Printf("actionResult: {%v}\n", actionResult)
	// for eng.NextPhaseRequired() {

	// }
}
