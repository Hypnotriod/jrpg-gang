package test

import (
	"fmt"
	"jrpg-gang/controller"
	"testing"
)

func TestCreateBattleRoom(t *testing.T) {
	controller := controller.NewController()
	result := controller.HandleRequest(`{
		"userId": "user-id-1",
		"type": "createBattleRoom",
		"data": {
			"allowedUsers": ["user-id-1", "user-id-2"]
		}
	}`)
	fmt.Println(result)
}

func TestJoin(t *testing.T) {
	controller := controller.NewController()
	var result string
	result = controller.HandleRequest(`{
		"id": "1234",
		"type": "join",
		"nickName": "999Megazilla"
	}`)
	fmt.Println(result)
	result = controller.HandleRequest(`{
		"id": "1234",
		"type": "join",
		"nickName": "Megazilla999"
	}`)
	fmt.Println(result)
	result = controller.HandleRequest(`{
		"id": "1234",
		"type": "join",
		"nickName": "Megazilla999"
	}`)
	fmt.Println(result)
}
