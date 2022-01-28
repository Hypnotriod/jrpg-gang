package test

import (
	"fmt"
	"jrpg-gang/controller"
	"jrpg-gang/util"
	"regexp"
	"strings"
	"testing"
)

func TestCreateBattleRoom(t *testing.T) {
	controller := controller.NewController()
	var result string
	result, userId := doJoinRequest(controller, "Megazilla999")
	fmt.Println(result)
	result = doRequest(controller, "createBattleRoom", userId, `
		"allowedUsers": ["user-id-1", "user-id-2"]
	`)
	fmt.Println(result)
}

func TestJoin(t *testing.T) {
	controller := controller.NewController()
	var result string
	result, _ = doJoinRequest(controller, "999Megazilla")
	fmt.Println(result)
	result, _ = doJoinRequest(controller, "Megazilla999")
	fmt.Println(result)
	result, _ = doJoinRequest(controller, "Megazilla999")
	fmt.Println(result)
}

func doJoinRequest(controller *controller.GameController, nickname string) (string, string) {
	result := controller.HandleRequest(fmt.Sprintf(`{
		"id": "%s",
		"type": "join",
		"nickName": "%s"
	}`,
		util.RandomId(),
		nickname))
	return result, parseUserId(result)
}

func parseUserId(str string) string {
	re := regexp.MustCompile(`"userId":"[a-z0-9]+`)
	found := re.FindAllString(str, 1)
	return strings.Split(found[0], `"userId":"`)[1]
}

func doRequest(controller *controller.GameController, requestType string, userId string, data string) string {
	return controller.HandleRequest(fmt.Sprintf(`{
		"id": "%s",
		"type": "%s",
		"userId": "%s",
		"data": {%s}
	}`,
		util.RandomId(),
		requestType,
		userId,
		data))
}
