package test

import (
	"fmt"
	"jrpg-gang/util"
	"testing"
)

func TestUnitThreatment(t *testing.T) {
	util.ApplyRandomSeed()
	hero := newAgileHero(t)
	hero.Inventory.Add(newSmallHealthPotion(t))
	fmt.Println(hero.Inventory.FindDisposable(2000))
	fmt.Println()
	fmt.Println(hero)
	fmt.Println()
	actionResult := hero.UseInventoryItemOnTarget(hero, 2000)
	hero.Inventory.FilterDisposable()
	fmt.Printf("instantRecovery: %v, temporalModification: %v\n", actionResult.InstantRecovery, actionResult.TemporalModification)
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
