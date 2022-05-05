package test

import (
	"fmt"
	"jrpg-gang/controller"
	"jrpg-gang/engine"
	"testing"
)

func TestCreateGameRoom(t *testing.T) {
	cntrl := controller.NewController()
	var result string
	result, user1Id := doJoinRequest(cntrl, "Megazilla999", engine.UnitClassTank)
	fmt.Println(result)
	fmt.Println()
	result = doRequest(cntrl, controller.RequestCreateGameRoom, user1Id, `
		"capacity": 2
	`)
	roomUid := parseRoomUid(result)
	fmt.Println(result)
	fmt.Println()
	result, user2Id := doJoinRequest(cntrl, "Megazilla777", engine.UnitClassMage)
	fmt.Println(result)
	fmt.Println()
	result = doRequest(cntrl, controller.RequestJoinGameRoom, user2Id, fmt.Sprintf(`
		"roomUid": %s
	`, roomUid))
	fmt.Println(result)
	fmt.Println()
	result = doRequest(cntrl, controller.RequestLobbyStatus, user1Id, ``)
	fmt.Println(result)
	fmt.Println()
}

func TestCreateGameRoomAsync(t *testing.T) {
	const n int = 1000
	var result string
	cntrl := controller.NewController()
	ch := make(chan string)
	for i := 0; i < n; i++ {
		go doCreateRoom(ch, cntrl, i)
	}
	for i := 0; i < n; i++ {
		result = <-ch
		status := parseStatus(result)
		if status != "ok" {
			fmt.Printf("create room failed: %s\n", result)
		}
	}
	_, userId := doJoinRequest(cntrl, "Host", engine.UnitClassMage)
	result = doRequest(cntrl, controller.RequestLobbyStatus, userId, ``)
	fmt.Println(result)
}

func doCreateRoom(ch chan<- string, cntrl *controller.GameController, i int) {
	_, userId1 := doJoinRequest(cntrl, fmt.Sprintf("Megazilla%d", i), engine.UnitClassRogue)
	result1 := doRequest(cntrl, controller.RequestCreateGameRoom, userId1, `
		"capacity": 4
	`)
	roomUid := parseUid(result1)
	_, userId2 := doJoinRequest(cntrl, fmt.Sprintf("Megazilla%d_buddy1", i), engine.UnitClassTank)
	result2 := doRequest(cntrl, controller.RequestJoinGameRoom, userId2, fmt.Sprintf(`
		"roomUid": %s
	`, roomUid))
	// doRequest(cntrl, controller.RequestDestroyGameRoom, userId1, ``)
	ch <- result2
}
