package test

import (
	"fmt"
	"jrpg-gang/controller"
	"jrpg-gang/engine"
	"testing"
)

type Broadcaster struct{}

func (b *Broadcaster) BroadcastGameMessage(userIds []engine.UserId, message string) {
	fmt.Println(message)
	fmt.Println()
}

func TestCreateGameRoom(t *testing.T) {
	broadcaster := &Broadcaster{}
	cntrl := controller.NewGameController(broadcaster)
	var result string
	user1Id, result := doJoinRequest(cntrl, "Megazilla999", engine.UnitClassTank)
	fmt.Println(result)
	fmt.Println()
	result = doRequest(cntrl, controller.RequestCreateGameRoom, user1Id, `
		"capacity": 2
	`)
	roomUid := parseRoomUid(result)
	fmt.Println(result)
	fmt.Println()
	user2Id, result := doJoinRequest(cntrl, "Megazilla777", engine.UnitClassMage)
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
	result = doRequest(cntrl, controller.RequestStartGame, user1Id, ``)
	fmt.Println(result)
	fmt.Println()
	result = doRequest(cntrl, controller.RequestNextGamePhase, user1Id, ``)
	fmt.Println(result)
	fmt.Println()
}

func TestCreateGameRoomAsync(t *testing.T) {
	const n int = 1000
	var result string
	broadcaster := &Broadcaster{}
	cntrl := controller.NewGameController(broadcaster)
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
	userId, _ := doJoinRequest(cntrl, "Host", engine.UnitClassMage)
	result = doRequest(cntrl, controller.RequestLobbyStatus, userId, ``)
	fmt.Println(result)
}

func doCreateRoom(ch chan<- string, cntrl *controller.GameController, i int) {
	userId1, _ := doJoinRequest(cntrl, fmt.Sprintf("Megazilla%d", i), engine.UnitClassRogue)
	result1 := doRequest(cntrl, controller.RequestCreateGameRoom, userId1, `
		"capacity": 4
	`)
	roomUid := parseUid(result1)
	userId2, _ := doJoinRequest(cntrl, fmt.Sprintf("Megazilla%d_buddy1", i), engine.UnitClassTank)
	result2 := doRequest(cntrl, controller.RequestJoinGameRoom, userId2, fmt.Sprintf(`
		"roomUid": %s
	`, roomUid))
	// doRequest(cntrl, controller.RequestDestroyGameRoom, userId1, ``)
	ch <- result2
}
