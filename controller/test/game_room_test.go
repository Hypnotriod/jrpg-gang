package test

import (
	"fmt"
	"jrpg-gang/controller"
	"testing"
)

func TestCreateGameRoom(t *testing.T) {
	cntrl := controller.NewController()
	var result string
	result, user1Id := doJoinRequest(cntrl, "Megazilla999")
	fmt.Println(result)
	result = doRequest(cntrl, controller.RequestCreateGameRoom, user1Id, `
		"capacity": 2
	`)
	roomUid := parseRoomUid(result)
	fmt.Println(result)
	result, user2Id := doJoinRequest(cntrl, "Megazilla777")
	fmt.Println(result)
	result = doRequest(cntrl, controller.RequestJoinGameRoom, user2Id, fmt.Sprintf(`
		"roomUid": %s
	`, roomUid))
	fmt.Println(result)
	result = doRequest(cntrl, controller.RequestLobbyStatus, user1Id, ``)
	fmt.Println(result)
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
	_, userId := doJoinRequest(cntrl, "Host")
	result = doRequest(cntrl, controller.RequestLobbyStatus, userId, ``)
	fmt.Println(result)
}

func doCreateRoom(ch chan<- string, cntrl *controller.GameController, i int) {
	_, userId := doJoinRequest(cntrl, fmt.Sprintf("Megazilla%d", i))
	result := doRequest(cntrl, controller.RequestCreateGameRoom, userId, `
		"capacity": 4
	`)
	ch <- result
}
