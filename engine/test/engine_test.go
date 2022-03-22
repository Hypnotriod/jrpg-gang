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
	gameLoop(t, eng)
}

func gameLoop(t *testing.T, eng *engine.GameEngine) {
	var event *engine.GameEvent
	fmt.Printf("state: {%v}\n\n", eng.NewGameEvent())
	event = eng.ExecuteUserAction(placeUnitAction(0, 0), "userId1")
	fmt.Printf("place action result: {%v}\n\n", event)
	for eng.NextPhaseRequired() {
		event = eng.NextPhase()
		fmt.Printf("event phase result: {%v}\n\n", event)
	}
	event = eng.ExecuteUserAction(attackUnitAction(3, 2), "userId1")
	fmt.Printf("attack action result: {%v}\n\n", event)
	for eng.NextPhaseRequired() {
		event = eng.NextPhase()
		fmt.Printf("event phase result: {%v}\n\n", event)
	}
}

func placeUnitAction(x, y int) domain.Action {
	return domain.Action{
		Uid:      1,
		Action:   domain.ActionPlace,
		Position: domain.Position{X: x, Y: y},
	}
}

func attackUnitAction(targetUid, itemUid uint) domain.Action {
	return domain.Action{
		Uid:       1,
		TargetUid: targetUid,
		ItemUid:   itemUid,
		Action:    domain.ActionUse,
	}
}
