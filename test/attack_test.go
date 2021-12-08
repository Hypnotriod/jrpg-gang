package test

import (
	"fmt"
	"jrpg-gang/util"
	"testing"
)

func TestAttackUnit(t *testing.T) {
	util.ApplyRandomSeed()
	northTroll := newNorthTroll(t)
	hero := newAgileHeroWithWeapon(t)
	fmt.Println(hero)
	fmt.Println()
	fmt.Println(northTroll)
	fmt.Println()
	actionResult := hero.UseInventoryItemOnTarget(northTroll, 4000)
	fmt.Printf("Attack with instant damage: %v, temporal impact: %v\n", actionResult.InstantDamage, actionResult.TemporalImpact)
	fmt.Println()
	fmt.Println(northTroll)
	for len(northTroll.Impact) > 0 {
		fmt.Println()
		fmt.Println("Turn Passed")
		fmt.Println()
		northTroll.ApplyDamageImpactOnNextTurn()
		northTroll.ReduceEnhancementOnNextTurn()
		fmt.Println(northTroll)
	}
}
