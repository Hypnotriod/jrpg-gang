package test

import (
	"fmt"
	"jrpg-gang/controller"
	"testing"
)

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
