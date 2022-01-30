package test

import (
	"fmt"
	"jrpg-gang/controller"
	"testing"
)

func TestJoin(t *testing.T) {
	cntrl := controller.NewController()
	var result string
	result, _ = doJoinRequest(cntrl, "999Megazilla")
	fmt.Println(result)
	result, _ = doJoinRequest(cntrl, "Megazilla999")
	fmt.Println(result)
	result, _ = doJoinRequest(cntrl, "Megazilla999")
	fmt.Println(result)
}

func TestJoinAsync(t *testing.T) {
	const n int = 10000
	var result string
	cntrl := controller.NewController()
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
	_, userId := doJoinRequest(cntrl, "Host")
	result = doRequest(cntrl, controller.RequestLobbyStatus, userId, ``)
	fmt.Println(result)
}

func doJoin(ch chan<- string, controller *controller.GameController, i int) {
	result, _ := doJoinRequest(controller, fmt.Sprintf("Megazilla%d", i))
	ch <- result
}
