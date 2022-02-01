package test

import (
	"fmt"
	"testing"
)

func TestUnitThreatmentWithDisposable(t *testing.T) {
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

func TestUnitThreatmentWithMagic(t *testing.T) {
	hero := newMagicianHeroWithMagic(t)
	fmt.Println(hero)
	fmt.Println()
	fmt.Printf("using %v on hero\n", hero.Inventory.Find(3001))
	fmt.Println()
	actionResult := hero.UseInventoryItemOnTarget(hero, 3001)
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
