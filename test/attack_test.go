package test

import (
	"fmt"
	"jrpg-gang/util"
	"testing"
)

func TestAttackUnit(t *testing.T) {
	util.ApplyRandomSeed()
	hero := newAgileHeroWithWeapon(t)
	northTroll := newNorthTroll(t)
	fmt.Println(hero)
	fmt.Println()
	fmt.Println(northTroll)
	fmt.Println()
	actionResult := hero.UseInventoryItemOnTarget(northTroll, 4000)
	fmt.Printf("%s attacks %s with: {%v}\n", hero.Name, northTroll.Name, actionResult)
	fmt.Println()
	fmt.Println(northTroll)
	for len(northTroll.Damage) > 0 {
		fmt.Println()
		fmt.Println("Turn Passed")
		fmt.Println()
		northTroll.ApplyDamageOnNextTurn()
		northTroll.ReduceModificationOnNextTurn()
		fmt.Println(northTroll)
	}
}
