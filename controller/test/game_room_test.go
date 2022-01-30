package test

import (
	"fmt"
	"jrpg-gang/controller"
	"testing"
)

func TestCreateGameRoom(t *testing.T) {
	cntrl := controller.NewController()
	var result string
	result, userId := doJoinRequest(cntrl, "Megazilla999")
	fmt.Println(result)
	result = doRequest(cntrl, controller.RequestCreateGameRoom, userId, `
		"capacity": 2
	`)
	fmt.Println(result)
}
