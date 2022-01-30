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
