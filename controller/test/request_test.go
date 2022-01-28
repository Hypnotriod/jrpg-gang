package test

import (
	"fmt"
	"jrpg-gang/controller"
	"jrpg-gang/util"
	"regexp"
	"strings"
)

func doJoinRequest(controller *controller.GameController, nickname string) (string, string) {
	result := controller.HandleRequest(fmt.Sprintf(`{
		"id": "%s",
		"type": "join",
		"data": {
			"nickName": "%s"
		}
	}`,
		util.RandomId(),
		nickname))
	return result, parseUserId(result)
}

func parseUserId(str string) string {
	re := regexp.MustCompile(`"userId":"[a-z0-9]+`)
	found := re.FindAllString(str, 1)
	if len(found) == 0 {
		return ""
	}
	return strings.Split(found[0], `"userId":"`)[1]
}

func doRequest(controller *controller.GameController, requestType controller.RequestType, userId string, data string) string {
	return controller.HandleRequest(fmt.Sprintf(`{
		"id": "%s",
		"type": "%s",
		"data": {"userId": "%s", %s}
	}`,
		util.RandomId(),
		requestType,
		userId,
		data))
}
