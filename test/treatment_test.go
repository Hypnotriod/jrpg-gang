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
	instantRecovery, temporalEnhancement := hero.Enhance(hero, healthPotion.Enhancement)
	fmt.Printf("instantRecovery: %v, temporalEnhancement: %v\n", instantRecovery, temporalEnhancement)
	fmt.Println()
	fmt.Println(hero)
	for len(hero.Enhancement) > 0 {
		hero.ApplyRecoverylEnhancementOnNextTurn()
		hero.ReduceEnhancementOnNextTurn()
		fmt.Println()
		fmt.Println("Turn Passed")
		fmt.Println()
		fmt.Println(hero)
	}
}
