package test

import (
	"fmt"
	"jrpg-gang/controller"
	"jrpg-gang/engine"
	"testing"
)

func TestJoin(t *testing.T) {
	cntrl := controller.NewGameController()
	var result string
	_, result = doJoinRequest(cntrl, "999Megazilla", engine.UnitClassMage)
	fmt.Println(result)
	_, result = doJoinRequest(cntrl, "Megazilla999", engine.UnitClassMage)
	fmt.Println(result)
	_, result = doJoinRequest(cntrl, "Megazilla999", engine.UnitClassMage)
	fmt.Println(result)
}

func TestJoinAsync(t *testing.T) {
	const n int = 1000
	var result string
	cntrl := controller.NewGameController()
	ch := make(chan string)
	for i := 0; i < n; i++ {
		go doJoin(ch, cntrl, i)
	}
	for i := 0; i < n; i++ {
		result = <-ch
		status := parseStatus(result)
		if status != "ok" {
			fmt.Printf("join failed: %s\n", result)
		}
	}
	userId, _ := doJoinRequest(cntrl, "Host", engine.UnitClassMage)
	result = doRequest(cntrl, controller.RequestLobbyStatus, userId, ``)
	fmt.Println(result)
}

func doJoin(ch chan<- string, controller *controller.GameController, i int) {
	_, result := doJoinRequest(controller, fmt.Sprintf("Megazilla%d", i), engine.UnitClassRogue)
	ch <- result
}
