package test

import (
	"fmt"
	"jrpg-gang/util"
	"testing"
)

func TestUnitThreatment(t *testing.T) {
	util.ApplyRandomSeed()
	hero := newAgileHero(t)
	healthPotion := newSmallHealthPotion(t)
	fmt.Println(healthPotion)
	fmt.Println()
	fmt.Println(hero)
	fmt.Println()
	instantRecovery, temporalModification := hero.Modify(hero, healthPotion.Modification)
	fmt.Printf("instantRecovery: %v, temporalModification: %v\n", instantRecovery, temporalModification)
	fmt.Println()
	fmt.Println(hero)
	for len(hero.Modification) > 0 {
		hero.ApplyRecoverylOnNextTurn()
		hero.ReduceModificationOnNextTurn()
		fmt.Println()
		fmt.Println("Turn Passed")
		fmt.Println()
		fmt.Println(hero)
	}
}
