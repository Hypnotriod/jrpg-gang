package test

import (
	"fmt"
	"jrpg-gang/controller"
	"jrpg-gang/engine"
	"jrpg-gang/util"
	"regexp"
	"strings"
)

var rndGen *util.RndGen = util.NewRndGen()

func doJoinRequest(controller *controller.GameController, nickname string, class engine.GameUnitClass) (string, string) {
	result := controller.HandleRequest(fmt.Sprintf(`{
		"id": "%s",
		"type": "join",
		"data": {
			"nickName": "%s",
			"class": "%s"
		}
	}`,
		rndGen.Hash(),
		nickname,
		class), engine.UserId(""))
	return result, parseUserId(result)
}

func doRequest(controller *controller.GameController, requestType controller.RequestType, userId string, data string) string {
	return controller.HandleRequest(fmt.Sprintf(`{
		"id": "%s",
		"type": "%s",
		"data": {%s}
	}`,
		rndGen.Hash(),
		requestType,
		data), engine.UserId(userId))
}

func parseUserId(str string) string {
	re := regexp.MustCompile(`"userId":"[a-z0-9]+`)
	found := re.FindAllString(str, 1)
	if len(found) == 0 {
		return ""
	}
	return strings.Split(found[0], `"userId":"`)[1]
}

func parseRoomUid(str string) string {
	re := regexp.MustCompile(`"room":{"uid":[a-f0-9]+`)
	found := re.FindAllString(str, 1)
	if len(found) == 0 {
		return ""
	}
	return strings.Split(found[0], `"room":{"uid":`)[1]
}

func parseStatus(str string) string {
	re := regexp.MustCompile(`"status":"[a-zA-Z]+`)
	found := re.FindAllString(str, 1)
	if len(found) == 0 {
		return ""
	}
	return strings.Split(found[0], `"status":"`)[1]
}

func parseUid(str string) string {
	re := regexp.MustCompile(`"uid":[0-9]+`)
	found := re.FindAllString(str, 1)
	if len(found) == 0 {
		return ""
	}
	return strings.Split(found[0], `"uid":`)[1]
}
