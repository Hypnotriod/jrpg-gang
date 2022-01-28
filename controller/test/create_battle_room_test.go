package test

import (
	"fmt"
	"jrpg-gang/controller"
	"testing"
)

func TestCreateBattleRoom(t *testing.T) {
	cntrl := controller.NewController()
	var result string
	result, userId := doJoinRequest(cntrl, "Megazilla999")
	fmt.Println(result)
	result = doRequest(cntrl, controller.RequestCreateBattleRoom, userId, `
		"allowedUsers": ["user-id-1", "user-id-2"],
		"matrix": [
			[{"fractionIds": [0], "type": "space"}, {"fractionIds": [0], "type": "space"}, {"fractionIds": [0], "type": "space"}, {"fractionIds": [0], "type": "space"}],
			[{"fractionIds": [0], "type": "space"}, {"fractionIds": [0], "type": "space"}, {"fractionIds": [0], "type": "space"}, {"fractionIds": [0], "type": "space"}],
			[{"fractionIds": [1], "type": "space"}, {"fractionIds": [1], "type": "space"}, {"fractionIds": [1], "type": "space"}, {"fractionIds": [1], "type": "space"}],
			[{"fractionIds": [1], "type": "space"}, {"fractionIds": [1], "type": "space"}, {"fractionIds": [1], "type": "space"}, {"fractionIds": [1], "type": "space"}]
		]
	`)
	fmt.Println(result)
}
